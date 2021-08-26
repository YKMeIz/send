package main

import (
	"github.com/YKMeIz/send/server/internal"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/uploadFile", internal.UploadFile)
	mux.HandleFunc("/api/createLink", internal.CreateLink)
	mux.HandleFunc("/api/getFile", internal.GetFile)
	mux.HandleFunc("/ln/", internal.Link)

	if os.Getenv("DEV_MODE") == "1" {
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			director := func(req *http.Request) {
				req.Header = r.Header
				req.URL = r.URL
				req.URL.Scheme = "http"
				req.URL.Host = "localhost:8080"
			}
			proxy := &httputil.ReverseProxy{Director: director}
			proxy.ServeHTTP(w, r)
		})
	}

	http.ListenAndServe(":9090", mux)
}
