package postgres

import (
	"context"
	"fmt"
	"tz-zero-agency/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

/*
var sqlDB *sql.DB

	func InitDB(log *slog.Logger, cfg *config.Config) *sql.DB {
		var err error
		databaseString := fmt.Sprintf("host=%s  port=%d user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)
		sqlDB, err = sql.Open("postgres", databaseString)
		if err != nil {
			log.Error("Ошибка подключения к базе данных: %v", logger.Err(err))
		}

		sqlDB.SetMaxOpenConns(10)
		sqlDB.SetMaxIdleConns(5)
		sqlDB.SetConnMaxLifetime(0)

		if err = sqlDB.Ping(); err != nil {
			log.Error("Ошибка проверки подключения к базе данных: %v", logger.Err(err))
		}

		return sqlDB
	}

	func CloseDB() {
		if sqlDB != nil {
			fmt.Println("Закрываем соединения c базой данных.")
			sqlDB.Close()
		}
	}
*/

var pool *pgxpool.Pool

func NewPostgresDB(cfg *config.Config) (*pgxpool.Pool, error) {
	databaseString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	pool, err := pgxpool.New(context.Background(), databaseString)
	return pool, err
}

func CloseDB() {
	if pool != nil {
		fmt.Println("Закрываем соединения c базой данных.")
		pool.Close()
	}
}
