package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"riki/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

type Database struct {
	pool *pgxpool.Pool
	cfg  *config.DBConfig
}

func newPool(ctx context.Context, poolCfg *pgxpool.Config) (*pgxpool.Pool, error) {
	return pgxpool.NewWithConfig(ctx, poolCfg)
}

func createAppDatabaseIfNotExists(ctx context.Context, cfg *config.DBConfig) error {
	poolCfg, err := cfg.MainDBConfig()
	if err != nil {
		return fmt.Errorf("unable to make main db pool cfg: %w", err)
	}

	pool, err := newPool(ctx, poolCfg)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %w", err)
	}
	defer pool.Close()

	appDBName := cfg.AppDBName

	var dbExists bool
	sql := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = '%s');", appDBName)
	if err = pool.QueryRow(ctx, sql).Scan(&dbExists); err != nil {
		return fmt.Errorf("failed to check if database exists: %w", err)
	}

	if dbExists {
		fmt.Printf("Database '%s' already exists.\n", appDBName)
		return nil
	}

	if _, err := pool.Exec(ctx, fmt.Sprintf("CREATE DATABASE %s;", appDBName)); err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}

	fmt.Printf("Database '%s' created successfully.\n", appDBName)

	return nil
}

func createAppTablesIfNotExist(ctx context.Context, cfg *config.DBConfig) error {
	poolCfg, err := cfg.AppDBConfig()
	if err != nil {
		return fmt.Errorf("unable to make main db pool cfg: %w", err)
	}

	pool, err := newPool(ctx, poolCfg)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %w", err)
	}
	defer pool.Close()

	initSQL, err := os.ReadFile("internal/db/init.sql")
	if err != nil {
		return fmt.Errorf("failed loading init sql: %w", err)
	}

	if _, err = pool.Exec(ctx, string(initSQL)); err != nil {
		return fmt.Errorf("unable to execute init.sql: %w", err)
	}
	appDBName := cfg.AppDBName
	fmt.Printf("Successfully applied all '%s' database tables.\n", appDBName)

	return nil
}

func initDB(ctx context.Context, cfg *config.DBConfig) error {
	if err := createAppDatabaseIfNotExists(ctx, cfg); err != nil {
		return err
	}

	if err := createAppTablesIfNotExist(ctx, cfg); err != nil {
		return err
	}

	return nil
}
func New(ctx context.Context, cfg *config.DBConfig) (*Database, error) {

	if err := initDB(ctx, cfg); err != nil {
		return nil, err
	}

	poolCfg, err := cfg.AppDBConfig()
	if err != nil {
		return nil, fmt.Errorf("unable to make main db pool cfg: %w", err)
	}

	pool, err := newPool(ctx, poolCfg)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	return &Database{pool: pool, cfg: cfg}, nil
}

func (db *Database) Close() {
	db.pool.Close()
}

func (db *Database) Ping(ctx context.Context) error {
	fmt.Printf("Successfully pinged the %s database.\n", db.pool.Config().ConnConfig.Database)

	return db.pool.Ping(ctx)
}

func (db *Database) ApplySequences(ctx context.Context) error {
	initSQL, err := os.ReadFile("internal/db/seq.sql")
	if err != nil {
		return fmt.Errorf("failed loading init sql: %w", err)
	}

	if _, err = db.pool.Exec(ctx, string(initSQL)); err != nil {
		return fmt.Errorf("unable to execute init.sql: %w", err)
	}

	appDBName := db.cfg.AppDBName
	log.Printf("Successfully applied all sequences tables for the %s DB.\n", appDBName)

	return nil
}
