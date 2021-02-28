package orderservice

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

type menuItem struct {
	ID       string `json:"id"`
	Quantity int    `json:"quantity"`
}

type order struct {
	ID        string     `json:"id"`
	MenuItems []menuItem `json:"menuItems"`
}

type order2 struct {
	ID                 string     `json:"id"`
	MenuItems          []menuItem `json:"menuItems"`
	OrderedAtTimestamp int        `json:"orderedAtTimestamp"`
	Cost               int        `json:"cost"`
}

type kitty struct {
	Name string `json:"name"`
}

// Router ...
func Router() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/hello-world", handleHelloWorld).Methods(http.MethodGet)
	r.HandleFunc("/cat", handleKitty).Methods(http.MethodGet)

	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/order/{ID}", handleOrder).Methods(http.MethodGet)
	s.HandleFunc("/orders", handleOrders).Methods(http.MethodGet)

	return logMiddleware(r)
}

func logMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		beginTime := time.Now()
		h.ServeHTTP(w, r)

		endTime := time.Now()
		log.WithFields(log.Fields{
			"method":      r.Method,
			"url":         r.URL,
			"remoteAddr":  r.RemoteAddr,
			"userAgent":   r.UserAgent(),
			"elapsedTime": endTime.Sub(beginTime),
		}).Info("got a new request")
	})
}

func handleKitty(w http.ResponseWriter, _ *http.Request) {
	cat := kitty{"Кот"}
	b, err := json.Marshal(cat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if _, err = io.WriteString(w, string(b)); err != nil {
		log.WithField("err", err).Error("write response error")
	}
}

func handleHelloWorld(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func handleOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]
	order := order2{
		ID: id,
		MenuItems: []menuItem{
			{
				ID:       "asdasdasdasd",
				Quantity: 1,
			},
		},
		OrderedAtTimestamp: 123123123123,
		Cost:               999,
	}
	result, err := json.Marshal(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if _, err = io.WriteString(w, string(result)); err != nil {
		log.WithField("err", err).Error("write response error")
	}
}

func handleOrders(w http.ResponseWriter, _ *http.Request) {
	orders := []order{
		{
			ID: "asdasdasdasd",
			MenuItems: []menuItem{
				{
					ID:       "asdasdasdasd",
					Quantity: 1,
				},
			},
		},
	}
	result, err := json.Marshal(orders)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if _, err = io.WriteString(w, string(result)); err != nil {
		log.WithField("err", err).Error("write response error")
	}
}
