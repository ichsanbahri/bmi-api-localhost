package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Url for test:
	// http://127.0.0.1:8181/?height=167&weight=70
	//
	r.Get("/api/", func(w http.ResponseWriter, r *http.Request) {
		height, err := strconv.ParseFloat(r.URL.Query().Get("height"), 32)
		if err != nil {
			returnError(w, "Height must in number")
			return
		}

		weight, err := strconv.ParseFloat(r.URL.Query().Get("weight"), 32)
		if err != nil {
			returnError(w, "Height must in number")
			return
		}

		bmi := weight / math.Pow(height/100, 2)
		label := "Normal"
		if bmi > 25.0 {
			label = "Overweight"
		}

		res := fmt.Sprintf("%.1f", bmi)
		out, _ := json.Marshal(map[string]string{
			"bmi":   res,
			"label": label,
		})

		w.Write(out)

	})
	http.ListenAndServe(":8181", r)
}

func returnError(w http.ResponseWriter, str string) {
	out, _ := json.Marshal(map[string]string{
		"error": str,
	})

	w.Write(out)
}
