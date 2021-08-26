package internal

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	uploadPath = "./"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	f, handler, err := r.FormFile("file")
	if err != nil {
		println(w, http.StatusBadRequest, err)
		return
	}
	defer f.Close()

	normalizedName := randomName(64)
	for {
		if _, err := os.Stat(uploadPath + "/" + normalizedName); os.IsNotExist(err) {
			break
		}
		normalizedName = randomName(64)
	}

	nf, err := os.OpenFile(uploadPath+"/"+normalizedName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		println(w, http.StatusInternalServerError, err)
		return
	}

	defer nf.Close()

	h := md5.New()

	writers := io.MultiWriter(h, nf)

	if _, err := io.Copy(writers, f); err != nil {
		println(w, http.StatusInternalServerError, err)
		return
	}

	id := fmt.Sprintf("%x.%s", h.Sum(nil), handler.Filename)

	if err := os.Rename(uploadPath+"/"+normalizedName, uploadPath+"/"+id); err != nil {
		println(w, http.StatusInternalServerError, err)
		return
	}

	handleOK(w, struct {
		ID string `json:"id"`
	}{id})
}

func GetFile(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var data struct {
		ID  string `json:"id"`
		Tkn string `json:"tkn"`
	}

	err = json.Unmarshal(b, &data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if data.Tkn != "" {
		data.ID += data.Tkn
	}

	b, err = kvs.Get([]byte(data.ID))
	if b == nil || err != nil {
		if data.Tkn == "" {
			w.WriteHeader(http.StatusNonAuthoritativeInfo)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if _, err := os.Stat(uploadPath + string(b)); os.IsNotExist(err) {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	name := strings.SplitAfterN(string(b), ".", 2)

	b, err = json.Marshal(struct {
		ID string `json:"id"`
	}{name[1]})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
