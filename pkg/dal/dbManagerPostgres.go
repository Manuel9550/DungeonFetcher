package dal

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Manuel9550/DungeonFetch/pkg/entities"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	_ "github.com/lib/pq"
)

// Implimentation of the interface
type DBManagerPostgres struct {
	db     *sql.DB
	logger log.Logger
}

func NewDBManagerPostgres(connectionString string, logger log.Logger) (DBManager, error) {
	newDBManager := DBManagerPostgres{}

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		return &newDBManager, err
	}

	// Test connection to database
	err = db.Ping()
	if err != nil {
		return &newDBManager, err
	}

	newDBManager = DBManagerPostgres{
		db:     db,
		logger: log.With(logger, "DBManager", "pq"),
	}

	return &newDBManager, nil
}

func (dm *DBManagerPostgres) CreateItem(ctx context.Context, item entities.Item) error {

	// Want boolen value to be a 1 or 0 in SQL
	consumableValue := "TRUE"
	if item.Consumable == true {
		consumableValue = "FALSE"
	}

	// Create the insert query
	insertQuery := fmt.Sprintf("INSERT INTO Item (ID, ITEM_NAME, MIN_LEVEL, MAX_LEVEL,CONSUMABLE) VALUES ('%s','%s',%d,%d,'%s');", item.ID, item.ItemName, item.MinLevel, item.MaxLevel, consumableValue)

	// Attempt to insert it into the database
	_, err := dm.db.ExecContext(ctx, insertQuery)

	if err != nil {
		level.Error(dm.logger).Log("InsertError", err, "id", item.ID, "name", item.ItemName)
		return err
	}

	return nil

}

func (dm *DBManagerPostgres) GetItem(ctx context.Context, id string) (entities.Item, error) {
	returnedItem := entities.Item{}

	queryString := fmt.Sprintf("SELECT ID, ITEM_NAME, MIN_LEVEL, MAX_LEVEL, CONSUMABLE FROM ITEM WHERE ID = '%s';", id)

	err := dm.db.QueryRowContext(ctx, queryString).Scan(
		&returnedItem.ID,
		&returnedItem.ItemName,
		&returnedItem.MinLevel,
		&returnedItem.MaxLevel,
		&returnedItem.Consumable)

	if err != nil {
		level.Error(dm.logger).Log("QueryError", err, "id", id)
		return returnedItem, err
	}

	return returnedItem, nil
}
