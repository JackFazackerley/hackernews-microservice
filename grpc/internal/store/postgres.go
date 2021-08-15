package store

import (
	"context"
	"fmt"

	"grpc/internal/config"
	pb "grpc/internal/proto"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	queryAll   = `SELECT id, type, by, url, score, title, time FROM hackernews;`
	queryType  = `SELECT id, type, by, url, score, title, time FROM hackernews WHERE "type" = $1;`
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
		return nil, fmt.Errorf("parsing connection string: %w", err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("connecting to database: %w", err)
	}

	return &Postgres{
		conn: pool,
	}, nil
}

func (d Postgres) Put(item *pb.Item) error {
	if _, err := d.conn.Exec(context.Background(), insertType, item.Id, item.Type, item.By, item.Url, item.Score, item.Title, item.Time); err != nil {
		return err
	}

	return nil
}

func (d Postgres) All() ([]*pb.Item, error) {
	return d.query(queryAll)
}

func (d Postgres) Jobs() ([]*pb.Item, error) {
	return d.query(queryType, "job")
}

func (d Postgres) Stories() ([]*pb.Item, error) {
	return d.query(queryType, "story")
}

func (d Postgres) query(query string, args ...interface{}) ([]*pb.Item, error) {
	rows, err := d.conn.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*pb.Item

	for rows.Next() {
		item := &pb.Item{}
		if err := rows.Scan(&item.Id, &item.Type, &item.By, &item.Url, &item.Score, &item.Title, &item.Time); err != nil {
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
