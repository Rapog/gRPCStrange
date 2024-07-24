package postgresGorm

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	team00v1 "teamclient/api/protos/gen/go/gRPCServer"
)

type Storage struct {
	db *gorm.DB
}

func NewDB() (*Storage, error) {
	const op = "repo.postgres,NewDB"

	dns := "postgres://postgres:1234@localhost:5432/postgres?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dns))

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = db.AutoMigrate(Anomaly{})
	if err != nil {
		return nil, fmt.Errorf("%s: Gorm migrate: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) AddAnomaly(stream team00v1.ConnectResponse) error {

	err := s.db.Create(&Anomaly{
		SessionId: stream.SessionId,
		Frequency: stream.Frequency,
		Timestamp: stream.Time.AsTime().UTC(),
	})

	return err.Error
}
