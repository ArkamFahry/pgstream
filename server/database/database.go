package database

import "github.com/ArkamFahry/pgwarp/server/models"

type IDatabase interface {
	Get(database string) (models.Database, error)
	Create(database models.Database) (models.Database, error)
	Update(database models.Database) (models.Database, error)
	Delete(database string) error
}

type Database struct{}
