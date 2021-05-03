package dal

import (
	"context"
	"errors"

	"github.com/Manuel9550/DungeonFetch/pkg/entities"
	"github.com/go-kit/kit/log"
)

// Interface - Could be handy to have an interface if we are also using a Postgres database
type DBManager interface {
	CreateItem(ctx context.Context, item entities.Item) error
	GetItem(ctx context.Context, id string) (entities.Item, error)
}

var DataError = errors.New("Unable to handle database request")
var DBTypeError = errors.New("DBType unknown")

func CreateDBManager(connectiongString string, logger log.Logger, dbType string) (DBManager, error) {
	if createFunc, ok := GetDBTypeMap()[dbType]; ok {
		return createFunc(connectiongString, logger)
	}

	return nil, DBTypeError
}

func GetDBTypeMap() map[string]func(connectiongString string, logger log.Logger) (DBManager, error) {
	return map[string]func(connectiongString string, logger log.Logger) (DBManager, error){
		"MSSQL": NewDBManagerSQL,
		"PQ":    NewDBManagerPostgres,
	}
}
