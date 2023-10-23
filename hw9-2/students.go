package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	zlog "github.com/rs/zerolog/log"
)

var Class map[string]*Participant = make(map[string]*Participant)

type Participant struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Teacher bool   `json:"is_teacher"`
}

func students(w http.ResponseWriter, r *http.Request) {
	zlog.Print("Handled " + r.URL.Path)
	resp := make(map[string]string)
	body, err := io.ReadAll(r.Body)

	if err != nil {
		zlog.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var credentials Account

	err = json.Unmarshal(body, &credentials)
	if err != nil {
		zlog.Print(err)
		resp["error"] = "Failed to unmarshal json!"
		jsonResp, _ := json.Marshal(resp)
		http.Error(w, string(jsonResp), http.StatusInternalServerError)
		return
	}

	account := Accounts[credentials.Login]
	partipicant := Class[account.UID]

	if !partipicant.IsTeacher() {
		zlog.Print(err)
		resp["error"] = "Only teachers can get info about students!"
		jsonResp, _ := json.Marshal(resp)
		http.Error(w, string(jsonResp), http.StatusInternalServerError)
		return

	}

	jsonResp, _ := json.Marshal(&Class)
	http.Error(w, string(jsonResp), http.StatusOK)
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	zlog.Print("Handled " + r.URL.Path)
	vars := mux.Vars(r)
	resp := make(map[string]string)
	body, err := io.ReadAll(r.Body)

	if err != nil {
		zlog.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var credentials Account

	err = json.Unmarshal(body, &credentials)
	if err != nil {
		zlog.Print(err)
		resp["error"] = "Failed to unmarshal json!"
		jsonResp, _ := json.Marshal(resp)
		http.Error(w, string(jsonResp), http.StatusInternalServerError)
		return
	}

	account := Accounts[credentials.Login]
	partipicant := Class[account.UID]

	if !partipicant.IsTeacher() {
		zlog.Print(err)
		resp["error"] = "Only teachers can get info about students!"
		jsonResp, _ := json.Marshal(resp)
		http.Error(w, string(jsonResp), http.StatusInternalServerError)
		return

	}

	jsonResp, _ := json.Marshal(Class[vars["uid"]])

	http.Error(w, string(jsonResp), http.StatusOK)
}
