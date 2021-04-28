package dal

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/Manuel9550/DungeonFetch/pkg/entities"
	"github.com/go-kit/kit/log"

	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var db *sql.DB
var mock sqlmock.Sqlmock
var ctx context.Context

var MockDBManager DBManagerSQL

func TestMain(m *testing.M) {

	ctx = context.Background()

	var err error

	db, mock, err = sqlmock.New()

	logger := log.NewLogfmtLogger(os.Stderr)

	MockDBManager.db = db
	MockDBManager.logger = log.With(logger, "testing", "dal")
	if err != nil {
		MockDBManager.logger.Log("setup_error", err.Error())
	}

	code := m.Run()
	db.Close()

	os.Exit(code)
}

func TestFindItem(t *testing.T) {

	testItem := entities.Item{
		ID:         "TestID",
		ItemName:   "TestName",
		MinLevel:   12,
		MaxLevel:   20,
		Consumable: false,
	}

	query := fmt.Sprintf("SELECT ID, ITEM_NAME, MIN_LEVEL, MAX_LEVEL, CONSUMABLE FROM ITEM WHERE ID = '%s'", testItem.ID)

	rows := sqlmock.NewRows([]string{"ID", "ITEM_NAME", "MIN_LEVEL", "MAX_LEVEL", "CONSUMABLE"}).
		AddRow(testItem.ID, testItem.ItemName, testItem.MinLevel, testItem.MaxLevel, testItem.Consumable)

	mock.ExpectQuery(query).WillReturnRows(rows)

	foundItem, err := MockDBManager.GetItem(ctx, testItem.ID)
	if err != nil {
		t.Errorf("error was not expected when fetching a new item: %s", err)
	}

	assert.Equal(t, foundItem.ID, testItem.ID)
	assert.Equal(t, foundItem.ItemName, testItem.ItemName)
	assert.Equal(t, foundItem.MinLevel, testItem.MinLevel)
	assert.Equal(t, foundItem.MaxLevel, testItem.MaxLevel)
	assert.Equal(t, foundItem.Consumable, testItem.Consumable)

}

func TestCreateItem(t *testing.T) {

	testItem := entities.Item{
		ID:         "TestID",
		ItemName:   "TestName",
		MinLevel:   12,
		MaxLevel:   20,
		Consumable: false,
	}

	query := regexp.QuoteMeta(fmt.Sprintf("INSERT INTO Item VALUES ('%s','%s',%d,%d,%d)", testItem.ID, testItem.ItemName, testItem.MinLevel, testItem.MaxLevel, 0))

	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))

	err := MockDBManager.CreateItem(ctx, testItem)
	if err != nil {
		t.Errorf("error was not expected when inserting a new item: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}
