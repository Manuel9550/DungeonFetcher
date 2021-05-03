package dal

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Manuel9550/DungeonFetch/pkg/entities"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// Implimentation of the interface
type DBManagerSQL struct {
	db     *sql.DB
	logger log.Logger
}

func NewDBManagerSQL(connectionString string, logger log.Logger) (DBManager, error) {
	newDBManager := DBManagerSQL{}

	db, err := sql.Open("sqlserver", connectionString)

	if err != nil {
		return &newDBManager, err
	}

	// Test connection to database
	err = db.Ping()
	if err != nil {
		return &newDBManager, err
	}

	newDBManager = DBManagerSQL{
		db:     db,
		logger: log.With(logger, "DBManager", "mssql"),
	}

	return &newDBManager, nil
}

func (dm *DBManagerSQL) CreateItem(ctx context.Context, item entities.Item) error {

	// Want boolen value to be a 1 or 0 in SQL
	consumableValue := 0
	if item.Consumable == true {
		consumableValue = 1
	}

	// Create the insert query
	insertQuery := fmt.Sprintf("INSERT INTO Item VALUES ('%s','%s',%d,%d,%d)", item.ID, item.ItemName, item.MinLevel, item.MaxLevel, consumableValue)

	// Attempt to insert it into the database
	_, err := dm.db.ExecContext(ctx, insertQuery)

	if err != nil {
		level.Error(dm.logger).Log("InsertError", err, "id", item.ID, "name", item.ItemName)
		return err
	}

	return nil

}

func (dm *DBManagerSQL) GetItem(ctx context.Context, id string) (entities.Item, error) {
	returnedItem := entities.Item{}

	queryString := fmt.Sprintf("SELECT ID, ITEM_NAME, MIN_LEVEL, MAX_LEVEL, CONSUMABLE FROM ITEM WHERE ID = '%s'", id)

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
