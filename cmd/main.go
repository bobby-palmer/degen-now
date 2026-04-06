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

func generateTableID(n int) string {
	const chars = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"
	b := make([]byte, n)
	rand.Read(b)
	for i := range b {
		b[i] = chars[int(b[i])%len(chars)]
	}
	return string(b)
}

func main() {

	var tablesMu sync.Mutex
	tables := make(map[string]*game.Table)

	http.HandleFunc("/api/create", func(w http.ResponseWriter, r *http.Request) {
		tableId := generateTableID(6)

		tablesMu.Lock()
		tables[tableId] = game.NewTable()
		tablesMu.Unlock()

		type Response struct {
			TableID string
		}

		if err := json.NewEncoder(w).Encode(Response{tableId}); err != nil {
			slog.Error("encoding table id", "err", err)
		}

		slog.Debug("create table", "tableId", tableId)
	})

	http.HandleFunc("/api/join", func(w http.ResponseWriter, r *http.Request) {

		type Request struct {
			Name    string
			TableID string
			Stack   int64
		}

		var request Request
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "decoding request: %v", err)
			return
		}

		tablesMu.Lock()
		table, ok := tables[request.TableID]
		tablesMu.Unlock()

		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "TableID not found")
			return
		}

		if err := table.Join(request.Name, request.Stack); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "joining table: %v", err)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     request.TableID,
			Value:    request.Name,
			Path:     "/",
			HttpOnly: true,
			MaxAge:   86400,
		})

		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		tableId := r.URL.Query().Get("table")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("serving:", err)
	}

}
