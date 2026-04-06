package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"sync"

	"github.com/bobby-palmer/degen-now/internal/game"
)

func generateRandomID(n int) string {
	const chars = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"
	b := make([]byte, n)
	rand.Read(b)
	for i := range b {
		b[i] = chars[int(b[i])%len(chars)]
	}
	return string(b)
}

type API struct {
	mu     sync.Mutex
	tables map[string]*game.Table
}

func NewApi() *API {
	return &API{
		mu:     sync.Mutex{},
		tables: make(map[string]*game.Table),
	}
}

func (api *API) getTable(tableID string) (*game.Table, error) {
	api.mu.Lock()
	defer api.mu.Unlock()

	table, ok := api.tables[tableID]
	if !ok {
		return nil, fmt.Errorf("table not found, id: %s", tableID)
	}

	return table, nil
}

func (api *API) Create(w http.ResponseWriter, r *http.Request) {
	api.mu.Lock()
	defer api.mu.Unlock()

	var tableID string
	for {
		tableID = generateRandomID(6)

		_, ok := api.tables[tableID]
		if !ok {
			break
		}
	}

	api.tables[tableID] = game.NewTable()

	type Response struct {
		TableID string `json:"table_id"`
	}

	if err := json.NewEncoder(w).Encode(Response{tableID}); err != nil {
		slog.Error("encoding response", "err", err)
	}

	slog.Debug("table created", "TableID", tableID)
}

func main() {

	api := NewApi()

	http.HandleFunc("/api/create", api.Create)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("serving:", err)
	}

}
