package migrator

import (
	"context"
	"fmt"
	"reflect"
	"riki/internal/db"
	"riki/internal/model"
	"riki/internal/riki_client"
	"sync"
)

const (
	THREADS = 20

	LOCATIONS  = "LOCATIONS"
	EPISODES   = "EPISODES"
	CHARACTERS = "CHARACTERS"
)

type Migrator struct {
	rikiClient *riki_client.RikiClient
	db         *db.Database
}

func New(
	rikiClient *riki_client.RikiClient,
	db *db.Database,
) *Migrator {
	return &Migrator{rikiClient, db}
}
func (m *Migrator) Run(ctx context.Context) error {
	for _, dataSet := range []string{LOCATIONS, CHARACTERS, EPISODES} {
		if err := m.loadAll(dataSet); err != nil {
			return err
		}
	}

	if err := m.db.ApplySequences(ctx); err != nil {
		return err
	}

	return nil
}

func (m *Migrator) loadAll(methodName string) error {
	var wg sync.WaitGroup
	errChan := make(chan error, 1)

	for page := 0; page < THREADS; page++ {
		wg.Add(1)
		go m.loadItems(page, methodName, &wg, errChan)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}

type MigrationFuncs struct {
	FetchFunc  func(int, *riki_client.RikiClient) (any, error)
	InsertFunc func(context.Context, any, *db.Database) error
}

var migrationMap = map[string]*MigrationFuncs{
	LOCATIONS: {
		FetchFunc: func(page int, rikiClient *riki_client.RikiClient) (any, error) {
			return rikiClient.GetLocationsByPage(page)
		},
		InsertFunc: func(ctx context.Context, items any, db *db.Database) error {
			return db.InsertLocations(ctx, items.([]model.Location))
		},
	},
	CHARACTERS: {
		FetchFunc: func(page int, rikiClient *riki_client.RikiClient) (any, error) {
			return rikiClient.GetCharactersByPage(page)
		},
		InsertFunc: func(ctx context.Context, items any, db *db.Database) error {
			return db.InsertCharacters(ctx, items.([]model.Character))
		},
	},

	EPISODES: {
		FetchFunc: func(page int, rikiClient *riki_client.RikiClient) (any, error) {
			return rikiClient.GetEpisodesByPage(page)
		},
		InsertFunc: func(ctx context.Context, items any, db *db.Database) error {
			return db.InsertEpisodeCharacters(ctx, items.([]model.Episode))
		},
	},
}

func getMigrationFuncs(methodName string) (*MigrationFuncs, error) {

	if fs, ok := migrationMap[methodName]; ok {
		return fs, nil

	}

	return nil, fmt.Errorf("unknown migration type")

}
func (m *Migrator) loadItems(
	page int,
	methodName string,
	wg *sync.WaitGroup,
	errChan chan error,
) {
	defer wg.Done()

	for {

		fs, err := getMigrationFuncs(methodName)
		if err != nil {
			errChan <- err
			return
		}

		fetch, insert := fs.FetchFunc, fs.InsertFunc

		items, err := fetch(page, m.rikiClient)
		if err != nil {
			errChan <- err
			return
		}

		if v := reflect.ValueOf(items); v.Kind() == reflect.Slice && v.Len() == 0 {
			break
		}

		if err := insert(context.Background(), items, m.db); err != nil {
			errChan <- err
			return
		}

		page += THREADS
	}
}
