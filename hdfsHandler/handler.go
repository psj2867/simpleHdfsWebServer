package hdfshandler

import (
	"io"
	"net/http"

	"github.com/colinmarc/hdfs"
)

type HdfsHandler struct {
	client *hdfs.Client
}

func NewHandler(url string) *HdfsHandler {
	client, _ := hdfs.New(url)
	return &HdfsHandler{client}
}

func (s *HdfsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.GetFile(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("not supported"))
	}
}

func (s *HdfsHandler) GetFile(w http.ResponseWriter, r *http.Request) {
	f, err := s.client.Open(r.URL.EscapedPath())
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No such file"))
		return
	}
	w.Header().Set("Content-Type", "image/jpeg")
	io.Copy(w, f)
}
