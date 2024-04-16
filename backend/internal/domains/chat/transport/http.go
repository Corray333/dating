package transport

import "net/http"

type ChatStorage interface{}

func GetAllChats(store ChatStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
