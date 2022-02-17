package storage

import (
	"github.com/Georgiiagon/otus_golang_hw/hw12_13_14_15_calendar/internal/app"
	"github.com/Georgiiagon/otus_golang_hw/hw12_13_14_15_calendar/internal/config"
	memorystorage "github.com/Georgiiagon/otus_golang_hw/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/Georgiiagon/otus_golang_hw/hw12_13_14_15_calendar/internal/storage/sql"
)

func New(config config.DatabaseConf) app.Storage {
	switch config.Connection {
	case "postgres":
		return sqlstorage.New(config)
	case "memory":
		return memorystorage.New()
	}

	return memorystorage.New()
}
