package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	sqldblogger "github.com/simukti/sqldb-logger"

	"github.com/chiennguyen196/go-library/internal/common/logs"
)

func NewSqlDB() *sql.DB {
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USERNAME"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot connect to DB")
	}

	db = sqldblogger.OpenDriver(psqlConn, db.Driver(), logs.NewSQLLogAdapter())

	return db
}
