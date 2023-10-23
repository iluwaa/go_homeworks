package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	zlog "github.com/rs/zerolog/log"
)

var Accounts map[string]Account = make(map[string]Account)

type Account struct {
	UID      string `json:"-"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func register(w http.ResponseWriter, r *http.Request) {
	zlog.Print("Handled " + r.URL.Path)
	resp := make(map[string]string)
	body, err := io.ReadAll(r.Body)

	if err != nil {
		zlog.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var newParticipant *Participant

	err = json.Unmarshal(body, &newParticipant)
	if err != nil {
		zlog.Print(err)
		resp["error"] = "Failed to unmarshal json"
		jsonResp, _ := json.Marshal(resp)
		http.Error(w, string(jsonResp), http.StatusInternalServerError)
		return
	}

	var newAccount Account
	err = json.Unmarshal(body, &newAccount)
	if err != nil {
		zlog.Print(err)
		resp["error"] = "Failed to unmarshal json"
		jsonResp, _ := json.Marshal(resp)
		http.Error(w, string(jsonResp), http.StatusInternalServerError)
		return
	}

	if newAccount.Login == "" || newAccount.Password == "" || newParticipant.Name == "" || newParticipant.Surname == "" {
		resp["error"] = "Login, passowrd, name and surname should be defined and can't be empty"
		jsonResp, _ := json.Marshal(resp)
		http.Error(w, string(jsonResp), http.StatusBadRequest)
		return
	}

	if _, exists := Accounts[newAccount.Login]; exists {
		zlog.Print(err)
		resp["error"] = "Login already taken."
		jsonResp, _ := json.Marshal(resp)
		http.Error(w, string(jsonResp), http.StatusBadRequest)
		return
	}

	passwordHash := sha1.New()
	passwordHash.Write([]byte(newAccount.Password))
	newAccount.Password = hex.EncodeToString(passwordHash.Sum(nil))
	newAccount.UID = uuid.New().String()

	Accounts[newAccount.Login] = newAccount
	Class[newAccount.UID] = newParticipant

	zlog.Print("Registered!")
	resp["success"] = fmt.Sprintf("%s %s with login %s registered!", newParticipant.Name, newParticipant.Surname, newAccount.Login)
	jsonResp, _ := json.Marshal(resp)
	http.Error(w, string(jsonResp), http.StatusOK)
	return

}

func Authorized(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := make(map[string]string)
		body, err := io.ReadAll(r.Body)

		if err != nil {
			zlog.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var account Account
		err = json.Unmarshal(body, &account)
		if err != nil {
			zlog.Print(err)
			resp["error"] = "Failed to unmarshal json!"

			jsonResp, _ := json.Marshal(resp)
			http.Error(w, string(jsonResp), http.StatusInternalServerError)
			return
		}

		if account.Login == "" || account.Password == "" {
			resp["error"] = "Login, passowrd should be defined and can't be empty"
			jsonResp, _ := json.Marshal(resp)
			http.Error(w, string(jsonResp), http.StatusBadRequest)
			return
		}

		if _, exists := Accounts[account.Login]; !exists {
			zlog.Print(err)
			resp["error"] = "Not registered!"
			jsonResp, _ := json.Marshal(resp)
			http.Error(w, string(jsonResp), http.StatusUnauthorized)
			return
		}

		h := sha1.New()
		h.Write([]byte(account.Password))
		password := hex.EncodeToString(h.Sum(nil))

		if password != Accounts[account.Login].Password {
			zlog.Print(err)
			resp["error"] = "Unautorized!"
			jsonResp, _ := json.Marshal(resp)
			http.Error(w, string(jsonResp), http.StatusUnauthorized)
			return
		}

		// make body accesible again
		r.Body = io.NopCloser(bytes.NewBuffer(body))

		next(w, r)
	}
}

func (participant *Participant) IsTeacher() bool {
	if participant.Teacher {
		return true
	}
	return false
}
