package internal

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func Link(w http.ResponseWriter, r *http.Request) {
	// trim prefix: /ln/
	id := r.URL.Path[4:]

	tkn := r.Header.Get("Authorization")
	if tkn != "" {
		id += tkn
	}

	b, err := kvs.Get([]byte(id))
	if b == nil || err != nil {
		if tkn == "" {
			w.WriteHeader(http.StatusNonAuthoritativeInfo)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		return
	}

	f, err := ioutil.ReadFile(uploadPath + string(b))
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(f)
}

func CreateLink(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var (
		data struct {
			ID       string `json:"id"`
			Tkn      string `json:"tkn"`
			BurnDate string `json:"burn_date"`
		}
		duration time.Duration
	)

	err = json.Unmarshal(b, &data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := randomName(8)
	for {
		if b, err := kvs.Get([]byte(id)); err != nil || b == nil {
			break
		}
		id = randomName(8)
	}

	if data.BurnDate != "" {
		t, err := time.Parse(time.RFC3339, data.BurnDate)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		duration = t.Sub(time.Now())
		if err := kvs.Put([]byte(id+data.Tkn), []byte(data.ID), duration); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	} else {
		if err := kvs.Put([]byte(id+data.Tkn), []byte(data.ID), 0); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	b, err = json.Marshal(struct {
		ID string `json:"id"`
	}{id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
