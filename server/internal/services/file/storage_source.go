package fileservice

import (
	"io"
	"net/http"
)

const (
	fileStorageSourceHeader = "X-1Backend-Storage-Source"
	fileStorageSourceLocal  = "local"
	fileStorageSourceGCS    = "gcs"
)

type storageSourceCarrier interface {
	StorageSource() string
}

type storageSourceReadCloser struct {
	io.ReadCloser
	source string
}

func (r *storageSourceReadCloser) StorageSource() string {
	return r.source
}

func withStorageSource(reader io.ReadCloser, source string) io.ReadCloser {
	if source == "" {
		source = fileStorageSourceLocal
	}
	return &storageSourceReadCloser{
		ReadCloser: reader,
		source:     source,
	}
}

func storageSourceFromReader(reader io.Reader) string {
	if reader == nil {
		return ""
	}
	if carrier, ok := reader.(storageSourceCarrier); ok {
		return carrier.StorageSource()
	}
	return fileStorageSourceLocal
}

func setFileStorageSourceHeader(w http.ResponseWriter, source string) {
	if source == "" {
		source = fileStorageSourceLocal
	}
	w.Header().Set(fileStorageSourceHeader, source)
}

func setDefaultFileStorageSourceHeader(w http.ResponseWriter, source string) {
	if w.Header().Get(fileStorageSourceHeader) != "" {
		return
	}
	setFileStorageSourceHeader(w, source)
}
