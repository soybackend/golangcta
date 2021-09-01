package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type rest struct {
	entries map[string] string
}

func (mr *rest) requestGet(ident string, rw http.ResponseWriter) {
	if document, ok := mr.entries[ident]; ok {
		fmt.Fprintln(rw, document)
	} else {
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(rw, "not found: ", ident)
	}
}

func (mr *rest) requestPost (ident, document string, rw http.ResponseWriter){
	if _, ok := mr.entries[ident]; ok {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(rw, "entry already exist", ident)
	} else {
		mr.entries[ident] = document
		fmt.Fprintln(rw, "OK")
	}
}

func (mr *rest) requestDelete (ident string, rw http.ResponseWriter) {
	if _, ok := mr.entries[ident]; ok {
		delete(mr.entries, ident)
		fmt.Fprintln(rw, "OK")
	} else {
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(rw, "not found: ", ident)
	}
}

func (mr *rest) requestPut (ident, document string, rw http.ResponseWriter) {
	if _, ok := mr.entries[ident]; ok {
		mr.entries[ident] = document
		fmt.Fprintln(rw, "OK")
	} else {
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(rw, "not found: ", ident)
	}
}

func (mr *rest) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	//get value from path
	ident := req.URL.Path
	var document string

	if req.Body != nil {
		b, _ := ioutil.ReadAll(req.Body)
		document = string(b)
	}

	rw.Header().Add("Content-Type", "text/plain")

	switch req.Method {
	case http.MethodGet:
		mr.requestGet(ident, rw)
	case http.MethodPost:
		mr.requestPost(ident, document, rw)
	case http.MethodDelete:
		mr.requestDelete(ident, rw)
	case http.MethodPut:
		mr.requestPut(ident, document, rw)
	default:
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(rw, "method not valid", req.Method)
	}
}

func main()  {
	rest := Rest{
		entries: map[string] string{},
	}
	http.Handle("/", &rest)
	panic(http.ListenAndServe(":8080", nil))
}