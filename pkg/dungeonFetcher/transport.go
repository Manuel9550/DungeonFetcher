package dungeonFetcher

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	transporthttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

type CreateItemRequest struct {
	ItemName   string `json:"name"`
	MinLevel   int32  `json:"minLevel,omitempty"`
	MaxLevel   int32  `json:"maxLevel,omitempty"`
	Consumable bool   `json:"consumable,omitempty"`
}

type CreateItemResponse struct {
	Ok string `json:"ok"`
	ID string `json:"id"`
}

type GetItemRequest struct {
	ID string `json:"id"`
}

type GetItemResponse struct {
	ID         string `json:"id"`
	ItemName   string `json:"name"`
	MinLevel   int32  `json:"minLevel,omitempty"`
	MaxLevel   int32  `json:"maxLevel,omitempty"`
	Consumable bool   `json:"consumable,omitempty"`
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func DecodeCreateItemRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request CreateItemRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil

}

func DecodeGetItemRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetItemRequest
	vars := mux.Vars(r)

	req = GetItemRequest{
		ID: vars["id"],
	}

	return req, nil
}

func NewHTTPServer(ctx context.Context, endpoints ServiceEndpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter()

	// As we make more endpoints, might consider wrapping these in a for loop
	//handler := Logger(logger, transporthttp.NewServer(endpoints.GetItem, DecodeGetItemRequest, EncodeResponse))
	r.Methods("GET").Path("/item/{id}").Handler(transporthttp.NewServer(endpoints.GetItem, DecodeGetItemRequest, EncodeResponse))

	//handler = Logger(logger, transporthttp.NewServer(endpoints.CreateItem, DecodeCreateItemRequest, EncodeResponse))
	r.Methods("POST").Path("/item").Handler(transporthttp.NewServer(endpoints.CreateItem, DecodeCreateItemRequest, EncodeResponse))
	r.Use(LoggingMiddleware(logger))
	return r
}

func LoggingMiddleware(logger log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			t := time.Now()

			logger = log.With(
				logger,
				"time", t.Format(time.ANSIC),
				"method", r.Method,
				"uri", r.RequestURI,
			)

			level.Info(logger).Log("msg", "request received")

			// Call the next handler
			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
