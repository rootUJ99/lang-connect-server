package jsonutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func JsonDecode[T any](r *http.Request) (T, error){
	var body T
	err:= json.NewDecoder(r.Body).Decode(&body)
	if (err == nil) {
		return body, nil
	}
	if errors.Is(err, io.EOF) {
		return body, fmt.Errorf("request body is empty to parse") 	
	} else {
		return body, err
	}
}

func SendJson(w http.ResponseWriter, status int, requestBody any) (error){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err:=json.NewEncoder(w).Encode(requestBody)	
	if (err!= nil) {
		return err
	}
	return nil
}


