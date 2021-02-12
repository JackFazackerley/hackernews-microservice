package store

import (
	"context"
	"fmt"

	"github.com/JackFazackerley/hackernews-microservice/internal/config"

	"github.com/JackFazackerley/hackernews-microservice/internal/client"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

const (
	queryAll = `SELECT id, type, by, url, score, title, time FROM hackernews;`

	queryType = `SELECT id, type, by, url, score, title, time FROM hackernews WHERE "type" = $1;`

	insertType = `INSERT INTO hackernews (id, type, by, url, score, title, time)  VALUES ($1, $2, $3, $4, $5, $6,
	$7) ON CONFLICT (id) DO NOTHING;`
)

type Postgres struct {
	conn *pgxpool.Pool
}

func NewPostgres(c config.Postgres) (*Postgres, error) {
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable&pool_max_conns=20",
		c.PostgresUsername(), c.PostgresPassword(), c.PostgresAddress(), c.PostgresDatabase(),
	)

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, errors.Wrap(err, "Call to pgx.NewConnPool failed")
	}

	return &Postgres{
		conn: pool,
	}, nil
}

func (d Postgres) Put(item *client.Item) error {
	if _, err := d.conn.Exec(context.Background(), insertType, item.ID, item.Type, item.By, item.URL, item.Score, item.Title, item.Time); err != nil {
		return err
	}

	return nil
}

func (d Postgres) All() ([]*Item, error) {
	return d.query(queryAll)
}

func (d Postgres) Jobs() ([]*Item, error) {
	return d.query(queryType, "job")
}

func (d Postgres) Stories() ([]*Item, error) {
	return d.query(queryType, "story")
}

func (d Postgres) query(query string, args ...interface{}) ([]*Item, error) {
	rows, err := d.conn.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*Item

	for rows.Next() {
		item := &Item{}
		if err := rows.Scan(&item.ID, &item.Type, &item.By, &item.URL, &item.Score, &item.Title, &item.Time); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func (d Postgres) Close() error {
	d.conn.Close()
	return nil
}
