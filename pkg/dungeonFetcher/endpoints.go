package dungeonFetcher

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type ServiceEndpoints struct {
	CreateItem endpoint.Endpoint
	GetItem    endpoint.Endpoint
}

func MakeEndpoints(service Service) ServiceEndpoints {

	return ServiceEndpoints{
		CreateItem: makeCreateItemEndpoint(service),
		GetItem:    makeGetItemEndpoint(service),
	}
}

func makeCreateItemEndpoint(service Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateItemRequest)
		ok, Id, err := service.CreateItem(ctx, req.ItemName, req.MinLevel, req.MaxLevel, req.Consumable)
		return CreateItemResponse{
			Ok: ok,
			ID: Id,
		}, err
	}
}

func makeGetItemEndpoint(service Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetItemRequest)
		item, err := service.GetItem(ctx, req.ID)

		return GetItemResponse{
			ID:         item.ID,
			ItemName:   item.ItemName,
			MinLevel:   item.MinLevel,
			MaxLevel:   item.MaxLevel,
			Consumable: item.Consumable,
		}, err

	}
}
