package main

import (
	"io"
	"net/http"
	// "encoding/json"
)

func sendErrorResponse(w http.ResponseWriter, sc int, errMsg string) {
	w.WriteHeader(sc)
	io.WriteString(w, errMsg)
}