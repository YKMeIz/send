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
	mux.HandleFunc("/file/", internal.GetFileWithoutTkn)

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
	} else {
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

			if _, err := os.Stat("./dist" + r.URL.Path); os.IsNotExist(err) || r.URL.Path == "/" {
				http.ServeFile(w, r, "./dist/index.html")
				return
			}
			http.ServeFile(w, r, "./dist"+r.URL.Path)
		})
	}

	http.ListenAndServe(":9090", mux)
}
