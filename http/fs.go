package http

import "net/http"

type Fs struct {
	http.FileSystem
}

func Serve() {
	handler := http.FileServer(http.Dir("."))
	mu := http.NewServeMux()
	mu.Handle("/", handler)
	err := http.ListenAndServe(":1313", mu)
	if err != nil {
		panic(err)
	}
}
