package http

import (
	"net/http"
	"encoding/json"
	"log"
	_"fmt"
)

func readJSON(req *http.Request, target interface{}) error {
	if err := json.NewDecoder(req.Body).Decode(&target); err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func writeJSON(w http.ResponseWriter, data interface{}, status int) {
	my_json, err := json.Marshal(data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	//fmt.Print("Success")
	w.WriteHeader(http.StatusCreated)
	w.Write(my_json)
}