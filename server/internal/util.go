package internal

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func handleOK(w http.ResponseWriter, v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		println(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func randomName(size int) string {
	b := make([]byte, size)
	rand.Seed(time.Now().UnixNano())
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
