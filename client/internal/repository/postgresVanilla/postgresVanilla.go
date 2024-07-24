package postgresVanilla

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"log"
	team00v1 "teamclient/api/protos/gen/go/gRPCServer"
)

type Storage struct {
	db *sql.DB
}

func NewDB() (*Storage, error) {
	const op = "repo.postgres,NewDB"

	db, err := sql.Open("postgres", "postgres://postgres:1234@localhost:5432/postgres?sslmode=disable")
	//defer db.Close()

	if err != nil {
		return nil, fmt.Errorf("%s: %w 1", op, err)
	}
	if _, err = db.Exec(`CREATE TABLE IF NOT EXISTS anomalies (
	  id serial not null unique,
    session_id varchar(255) not null,
    frequency  float,
    tmstp      timestamp with time zone
	);`); err != nil {
		return nil, fmt.Errorf("%s: %w 2", op, err)
	}
	return &Storage{db: db}, nil
}

func (s *Storage) AddAnomaly(stream team00v1.ConnectResponse) error {
	const op = "repo.postgres.AddAnomaly"
	stmt, err := s.db.Prepare("INSERT INTO anomalies (session_id, frequency, tmstp) VALUES ($1, $2, $3)")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.Exec(stream.SessionId, stream.Frequency, stream.Time.AsTime().UTC())
	if err != nil {
		pgErr, ok := err.(*pq.Error)
		if ok && pgErr.Code == "23505" {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: failed to get rows affected: %w", op, err)
	}
	log.Printf("Rows affected: %d", rowsAffected)

	return nil
}
