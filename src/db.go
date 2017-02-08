package main

import (
	"log"

	"github.com/jackc/pgx"
)

var (
	barsLastQuery    *pgx.PreparedStatement
	barsBetweenQuery *pgx.PreparedStatement
	pool             *pgx.ConnPool
)

func initDatabase() {

	var cfg pgx.ConnPoolConfig
	cfg.Host = "localhost"
	cfg.User = "chartingdbuser"
	cfg.Password = "chartingdbpass"
	cfg.Database = "chartingdb"
	cfg.Port = 5432
	cfg.MaxConnections = 40

	cfg.AfterConnect = func(conn *pgx.Conn) error {
		barsLastQuery = prepareQueries(conn, "barsLastQuery", "SELECT * FROM PriceTick_Yearly ORDER BY TickDate DESC LIMIT $1")
		barsBetweenQuery = prepareQueries(conn, "barsBetweenQuery", "SELECT * FROM PriceTick_Yearly WHERE TickDate >= '&1' AND TickDate < '$2'")
		return nil
	}

	connPool, err := pgx.NewConnPool(cfg)
	if err != nil {
		log.Println("Database connection failed. Database: ", cfg.Database, ", user: ", cfg.User, ", error: ", err)
	} else {
		log.Println("Database connection established. Database: ", cfg.Database, ", user: ", cfg.User, ", error: ", err)
	}
	pool = connPool
}

func barsLast(count int) []PriceBar {
	rows, err := pool.Query("barsLastQuery", count)
	if err != nil {
		log.Fatalf("Database error: %v", err)
	}
	defer rows.Close()
	return readBars(rows)
}

func barsBetween(from string, to string) []PriceBar {
	rows, err := pool.Query("barsBetweenQuery", from, to)
	if err != nil {
		log.Fatalf("Database error: %v", err)
	}
	defer rows.Close()
	return readBars(rows)
}

func readBars(rows *pgx.Rows) []PriceBar {
	bars := make([]PriceBar, 0)
	for rows.Next() {
		var bar PriceBar
		if err := rows.Scan(
			&bar.TickDate,
			&bar.TickYear,
			&bar.MarketID,
			&bar.Open,
			&bar.Close,
			&bar.High,
			&bar.Low,
			&bar.OpenTickdate,
			&bar.OpenVersionNo,
			&bar.CloseTickDate,
			&bar.CloseVersionNo,
			&bar.Spike,
			&bar.Gap,
			&bar.OpenBid,
			&bar.OpenAsk,
			&bar.HighBid,
			&bar.HighAsk,
			&bar.LowBid,
			&bar.LowAsk,
			&bar.CloseBid,
			&bar.CloseAsk); err != nil {
			log.Fatalf("Error reading data from database: %s", err)
		}
		bars = append(bars, bar)
	}
	return bars
}

func prepareQueries(connection *pgx.Conn, name, query string) *pgx.PreparedStatement {
	stmt, err := connection.Prepare(name, query)
	if err != nil {
		log.Fatalf("Query preparing error %q: %s", query, err)
	}
	return stmt
}
