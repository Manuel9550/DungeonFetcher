package dungeonFetcher

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/Manuel9550/DungeonFetcher/pkg/dal"
	"github.com/Manuel9550/DungeonFetcher/pkg/entities"
)

type Service interface {
	CreateItem(ctx context.Context, name string, minLevel int32, maxLevel int32, consumable bool) (string, string, error)
	GetItem(ctx context.Context, id string) (entities.Item, error)
}

type DungeonFetcher struct {
	logger          log.Logger
	databaseManager dal.DBManager
}

func NewService(logger log.Logger, dm dal.DBManager) DungeonFetcher {
	newService := DungeonFetcher{
		logger:          logger,
		databaseManager: dm,
	}

	return newService
}

func (d *DungeonFetcher) CreateItem(ctx context.Context, name string, minLevel int32, maxLevel int32, consumable bool) (string, string, error) {

	logger := log.With(d.logger, "method", "CreateItem")

	// Create item
	newItem, err := entities.CreateItem(name, minLevel, maxLevel, consumable)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "", "", err
	}

	// Add item into database
	if err := d.databaseManager.CreateItem(ctx, newItem); err != nil {
		level.Error(logger).Log("err", err)
		return "", "", err
	}

	logger.Log("CreatedItem", newItem.ItemName, "ID", newItem.ID)

	return newItem.ID, "success", nil
}

func (d *DungeonFetcher) GetItem(ctx context.Context, id string) (entities.Item, error) {

	logger := log.With(d.logger, "method", "GetItem")

	retrievedItem, err := d.databaseManager.GetItem(ctx, id)

	if err != nil {
		level.Error(logger).Log("err", err)
		return retrievedItem, err
	}

	logger.Log("RequestedItem", retrievedItem.ItemName, "ID", retrievedItem.ID)

	return retrievedItem, nil
}
