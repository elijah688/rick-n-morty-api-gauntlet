package gateway

import (
	"context"
	"errors"
	"riki_gateway/internal/model"
	"sync"
)

func (cs *GatewayService) CompileCharacter(ctx context.Context, id int) (*model.Character, error) {
	character, err := cs.getCharacterByID(ctx, id)
	if err != nil {
		return nil, err
	}

	chars, err := cs.compileCharacters(ctx, []model.Character{*character})
	if err != nil {
		return nil, err
	}

	if len(chars) > 0 {
		return &chars[0], nil
	}

	return nil, errors.New("filed compiling character")
}

func (cs *GatewayService) compileCharacters(ctx context.Context, characters []model.Character) ([]model.Character, error) {

	if len(characters) == 0 {
		return nil, nil
	}

	charsWithLocs, err := cs.fetchAndSetLocations(ctx, characters)
	if err != nil {
		return nil, err
	}

	ids := getIDs(charsWithLocs)
	wg, errChan := new(sync.WaitGroup), make(chan error, 1)
	wg.Add(2)

	episodes, eErr := map[int][]int(nil), error(nil)
	go func() {
		defer wg.Done()

		episodes, eErr = cs.getCharacterEpisodesByIDs(ctx, ids)
		if err != nil {
			errChan <- eErr
			return
		}
	}()

	debuts, dErr := map[int]*model.Episode(nil), error(nil)
	go func() {
		defer wg.Done()

		debuts, dErr = cs.getDebusByIDs(ctx, ids)
		if err != nil {
			errChan <- dErr
			return
		}

	}()

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return nil, err
		}
	}

	for i := range charsWithLocs {
		es := episodes[i]
		charsWithLocs[i].Episodes = &es
		charsWithLocs[i].Debut = debuts[i]

	}

	return charsWithLocs, nil
}

func getIDs(cs []model.Character) []int {
	res := make([]int, 0)
	for _, c := range cs {
		if c.ID != nil {
			res = append(res, *c.ID)
		}
	}
	return res
}
