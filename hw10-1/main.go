package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/mux"
	zlog "github.com/rs/zerolog/log"
)

var resp = make(map[string]string)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/weather", weather).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	zlog.Err(srv.ListenAndServe())
}

func weather(w http.ResponseWriter, r *http.Request) {
	zlog.Print("Handled " + r.URL.Path)

	city := r.URL.Query().Get("city")
	if city == "" {
		resp["error"] = fmt.Sprintf("Missing required parameter 'city'")
		jsonResp, _ := json.Marshal(resp)
		http.Error(w, string(jsonResp), http.StatusInternalServerError)
		return
	}

	weather, err := GetWeather(r.Context(), city)
	if err != nil {
		resp["error"] = fmt.Sprintf(err.Error())
		jsonResp, _ := json.Marshal(resp)
		http.Error(w, string(jsonResp), http.StatusInternalServerError)
		return
	}

	jsonResp, _ := json.Marshal(weather)
	Response(w, string(jsonResp), http.StatusOK)
}

func GetWeather(ctx context.Context, city string) (Weather, error) {
	var result Weather
	var error Error

	baseUrl, err := url.Parse("http://api.weatherapi.com/v1/forecast.json")
	if err != nil {
		return result, err
	}
	urlParams := url.Values{}
	urlParams.Add("q", city)
	urlParams.Add("key", "4488f4394c734960bd9163339232510")
	baseUrl.RawQuery = urlParams.Encode()

	request, err := http.NewRequestWithContext(ctx, "GET", baseUrl.String(), nil)
	if err != nil {
		return result, err
	}

	client := &http.Client{
		Timeout: time.Second * 15,
	}

	response, err := client.Do(request)
	if err != nil {
		return result, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}

	if result == (Weather{}) {
		err = json.Unmarshal(body, &error)
		if error != (Error{}) {
			return result, errors.New(error.Error.Message)
		}
	}

	return result, nil
}

func Response(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprintln(w, msg)
}

type Error struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type Weather struct {
	Location struct {
		Name      string  `json:"name"`
		Region    string  `json:"region"`
		Country   string  `json:"country"`
		Lat       float64 `json:"lat"`
		Lon       float64 `json:"lon"`
		TzID      string  `json:"tz_id"`
		Localtime string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdated string  `json:"last_updated"`
		TempC       float64 `json:"temp_c"`
		IsDay       int     `json:"is_day"`
		Condition   struct {
			Text string `json:"text"`
		} `json:"condition"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree int     `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		PrecipMm   float64 `json:"precip_mm"`
		PrecipIn   float64 `json:"precip_in"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		VisKm      float64 `json:"vis_km"`
		Uv         float64 `json:"uv"`
		GustKph    float64 `json:"gust_kph"`
	} `json:"current"`
}
