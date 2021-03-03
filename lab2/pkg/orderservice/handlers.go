package orderservice

import (
	"encoding/json"
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

// Router ...
func Router() http.Handler {
	r := mux.NewRouter()
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

func handleOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]
	order := findOrder(id)
	if order == nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	result, err := json.Marshal(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err = io.WriteString(w, string(result)); err != nil {
		log.WithField("err", err).Error("write response error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func handleOrders(w http.ResponseWriter, _ *http.Request) {
	orders := findAllOrders()
	result, err := json.Marshal(orders)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err = io.WriteString(w, string(result)); err != nil {
		log.WithField("err", err).Error("write response error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func findAllOrders() []order {
	return []order{
		{
			ID: "d290f1ee-6c54-4b01-90e6-d701748f0851",
			MenuItems: []menuItem{
				{
					ID:       "f290d1ce-6c234-4b31-90e6-d701748f0851",
					Quantity: 1,
				},
			},
		},
	}
}

func findOrder(id string) *order2 {
	if id != "d290f1ee-6c54-4b01-90e6-d701748f0851" {
		return nil
	}

	return &order2{
		ID: id,
		MenuItems: []menuItem{
			{
				ID:       "f290d1ce-6c234-4b31-90e6-d701748f0851",
				Quantity: 1,
			},
		},
		OrderedAtTimestamp: 1613758423,
		Cost:               999,
	}
}
