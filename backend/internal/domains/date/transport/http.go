package transport

import (
	"log/slog"
	"net/http"

	"github.com/Corray333/dating/pkg/server/auth"
	"github.com/gorilla/websocket"
)

type DateStorage interface {
	StartSearching(id int, conn *websocket.Conn) error
}

func NewDate(store DateStorage) http.HandlerFunc {
	upgrader := websocket.Upgrader{}
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()

		creds, err := auth.ExtractCredentials(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Failed to extract credentials", http.StatusBadRequest)
			slog.Error("Failed to extract credentials: " + err.Error())
			return
		}

		if err := store.StartSearching(creds.ID, conn); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			slog.Error("Failed to add user to searching map: " + err.Error())
			return
		}

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				return
			}
			WebSocketHandler(creds.ID, string(message), store)
		}
	}
}
