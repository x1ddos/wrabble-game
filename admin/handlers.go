package admin

import (
	"fmt"
	"html/template"
	"net/http"
	"runtime/debug"

	"appengine"
	"appengine/blobstore"
)

const (
	pathAdmin        = "/admin"
	pathImport       = "/admin/import"
	pathImportUpload = "/admin/import/upload"
	pathImportWords  = "/admin/import/words"
	pathImportDict   = "/admin/import/dict"

	uploadTemplateHTML = `
    <html><body>
    <form action="{{.}}" method="POST" enctype="multipart/form-data">
      Dict name: <input type="text" name="name"><br>
      Upload File: <input type="file" name="dict"><br>
      <input type="submit" name="submit" value="Submit">
    </form></body></html>`
)

var (
	uploadTemplate = template.Must(template.New("upload").Parse(uploadTemplateHTML))
)

func init() {
	http.Handle(pathAdmin, http.RedirectHandler(pathImport, http.StatusFound))
	http.HandleFunc(pathImport, importFormHandler)
	http.HandleFunc(pathImportUpload, importUploadHandler)
	http.HandleFunc(pathImportDict, importDictHandler)
	http.HandleFunc(pathImportWords, importWordsHandler)
}

func importFormHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	uploadURL, err := blobstore.UploadURL(c, pathImportUpload, nil)
	if err != nil {
		c.Errorf("error creating upload URL: %v", err)
		serveError(c, w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	err = uploadTemplate.Execute(w, uploadURL)
	if err != nil {
		c.Errorf("error parsing upload form: %v", err)
		serveError(c, w, err)
	}
}

func importUploadHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	blobs, params, err := blobstore.ParseUpload(r)
	if err != nil {
		serveError(c, w, err)
		return
	}
	if len(blobs["dict"]) == 0 {
		c.Errorf("no dictionary file uploaded")
		http.Redirect(w, r, pathImport, http.StatusFound)
		return
	}
	dict := blobs["dict"][0]
	name := params.Get("name")
	c.Infof("uploaded dict %q as %q (%s), %d bytes",
		dict.Filename, name, dict.ContentType, dict.Size)

	if err := asyncImportDict(c, dict.BlobKey, name); err != nil {
		c.Errorf("error async importing dict: %v", err)
		serveError(c, w, err)
	}

	fmt.Fprintf(w, "Successfully started importing %q as %q (%d bytes).",
		dict.Filename, name, dict.Size)
}

func importWordsHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	if err := r.ParseForm(); err != nil {
		c.Errorf("importWordsHandler: error parsing form: %v", err)
		return
	}
	dict := r.FormValue("dict")
	words := r.Form["words"]
	if err := importWords(c, dict, words); err != nil {
		c.Errorf("importWords error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func importDictHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	if err := r.ParseForm(); err != nil {
		c.Errorf("importDictHandler: error parsing form: %v", err)
		return
	}
	dict := appengine.BlobKey(r.FormValue("blobkey"))
	name := r.FormValue("name")
	if err := importDict(c, dict, name); err != nil {
		c.Errorf("importDict error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func serveError(c appengine.Context, w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Internal Server Error: %v\n\n%s\n", err, debug.Stack())
}
