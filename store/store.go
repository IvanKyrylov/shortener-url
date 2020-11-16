package store

import (
	"database/sql"
	"fmt"

	"github.com/IvanKyrylov/shortener-url/config"
	"github.com/IvanKyrylov/shortener-url/internal/generator"

	_ "github.com/lib/pq"
)

type postgres struct{ db *sql.DB }

func New(cfg *config.Config) (Service, error) {

	connect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Db)

	db, err := sql.Open("postgres", connect)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	strQuery := "CREATE TABLE IF NOT EXISTS shortener (uid serial NOT NULL, url VARCHAR not NULL, key VARCHAR not NULL, visited boolean DEFAULT FALSE, count INTEGER DEFAULT 0);"

	_, err = db.Exec(strQuery)
	if err != nil {
		return nil, err
	}
	return &postgres{db}, nil
}

func (p *postgres) Save(url string) (string, error) {
	key := generator.GenerateKey()
	err := p.db.QueryRow("INSERT INTO shortener(url,key,visited,count) VALUES($1,$2,$3,$4);", url, key, false, 0).Err()
	if err != nil {
		return "", err
	}
	return key, nil
}

func (p *postgres) Load(key string) (string, error) {
	var url string
	err := p.db.QueryRow("update shortener set visited=true, count = count + 1 where key=$1 RETURNING url", key).Scan(&url)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (p *postgres) LoadInfo(key string) (*ItemStore, error) {
	var item ItemStore
	err := p.db.QueryRow("SELECT url, visited, count FROM shortener where key=$1 limit 1", key).
		Scan(&item.URL, &item.Visited, &item.Count)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (p *postgres) Close() error { return p.db.Close() }
