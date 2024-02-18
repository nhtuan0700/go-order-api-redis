package util

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, code int, obj any) {
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(obj); err != nil {
		panic(err)
	}
}
