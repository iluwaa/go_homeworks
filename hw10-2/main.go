package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/translate"
	"github.com/gorilla/mux"
	zlog "github.com/rs/zerolog/log"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

var resp = make(map[string]string)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/translate", handleTranslate).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	zlog.Err(srv.ListenAndServe())
}

func handleTranslate(w http.ResponseWriter, r *http.Request) {
	zlog.Print("Handled " + r.URL.Path)

	body, err := io.ReadAll(r.Body)

	if err != nil {
		zlog.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var textForTranslate TextForTranslate
	err = json.Unmarshal(body, &textForTranslate)
	if err != nil {
		zlog.Print(err)
		http.Error(w, "Failed to unmarshal json", http.StatusInternalServerError)
		return
	}

	translatedText, err := translateText(r.Context(), textForTranslate.Text)
	if err != nil {
		resp["error"] = fmt.Sprintf(err.Error())
		jsonResp, _ := json.Marshal(resp)
		http.Error(w, string(jsonResp), http.StatusInternalServerError)
		return
	}

	jsonResp, _ := json.Marshal(translatedText)
	Response(w, string(jsonResp), http.StatusOK)
}

func translateText(ctx context.Context, text string) (translate.Translation, error) {
	var translatedText translate.Translation
	jsonKey, err := os.ReadFile("credentials.json")

	client, err := translate.NewClient(ctx, option.WithCredentialsJSON(jsonKey))
	if err != nil {
		return translatedText, err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, language.Ukrainian, nil)
	if err != nil {
		return translatedText, fmt.Errorf("Translate: %w", err)
	}

	if len(resp) == 0 {
		return translatedText, fmt.Errorf("Translate returned empty response to text: %s", text)
	}

	return resp[0], nil

}

func Response(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprintln(w, msg)
}

type TextForTranslate struct {
	Text string `json:"text"`
}
