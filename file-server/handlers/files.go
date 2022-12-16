package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/kokoichi206/go-expert/file-server/files"
)

// Files is a handler for reading and writing files
type Files struct {
	log   hclog.Logger
	store files.Storage
}

// NewFiles creates a new File handler
func NewFiles(s files.Storage, l hclog.Logger) *Files {
	return &Files{store: s, log: l}
}

func (f *Files) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fn := vars["filename"]

	f.log.Info("Handle POST", "id", id, "filename", fn)

	f.saveFile(id, fn, rw, r)
}

func (f *Files) invalidURI(uri string, rw http.ResponseWriter) {
	f.log.Error("Invalid path", "path", uri)
	http.Error(rw, "Invalid file path should be in the format: /[id]/[filepath]", http.StatusBadRequest)
}

// saveFile saves the contents of the request to a file
func (f *Files) saveFile(id, path string, rw http.ResponseWriter, r *http.Request) {
	f.log.Info("Save file for product", "id", id, "path", path)

	fp := filepath.Join(id, path)
	err := f.store.Save(fp, r.Body)
	if err != nil {
		f.log.Error("Unable to save file", "error", err)
		http.Error(rw, "Unable to save file", http.StatusInternalServerError)
	}
}

func (f *Files) AllFiles(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	f.log.Info("AllFiles", "id", id)

	// f.saveFile(id, fn, rw, r)
	files, err := ioutil.ReadDir(fmt.Sprintf("imagestore/%s", id))
	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf("%v", err),
			http.StatusInternalServerError)
		return
	}

	var names []string
	for _, file := range files {
		names = append(names, file.Name())
	}
	response := AllFilesResponse{
		Names: names,
	}
	e := json.NewEncoder(rw)
	e.Encode(response)
}

type AllFilesResponse struct {
	Names []string `json:"names"`
}
