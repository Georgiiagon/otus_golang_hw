package sqlstorage

import (
	"context"
	"strings"

	"github.com/Georgiiagon/otus_golang_hw/hw12_13_14_15_calendar/internal/config"
)

type Storage struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func New(config config.DatabaseConf) *Storage {
	return &Storage{}
}

func (s *Storage) Connect(ctx context.Context) error {
	// TODO
	return nil
}

func (s *Storage) Close(ctx context.Context) error {
	// TODO
	return nil
}

func (s *Storage) makdDSN() string {
	str := strings.Builder{}
	// host=localhost port=5432 user=root password=password dbname=docker sslmode=disable
	str.WriteString("host=")
	str.WriteString(s.Host)
	str.WriteString(" ")

	str.WriteString("port=")
	str.WriteString(s.Port)
	str.WriteString(" ")

	str.WriteString("user=")
	str.WriteString(s.User)
	str.WriteString(" ")

	str.WriteString("password=")
	str.WriteString(s.Password)
	str.WriteString(" ")

	str.WriteString("dbname=")
	str.WriteString(s.Database)
	str.WriteString(" ")

	return str.String()
}
