package main

import (
	"net/http"

	hdfshandler "github.com/psj2867/simpleHdfsWebServer/hdfsHandler"
)

func main() {
	http.ListenAndServe("0.0.0.0:8080", hdfshandler.NewHandler("n1.psj2867.com:8020"))

}
