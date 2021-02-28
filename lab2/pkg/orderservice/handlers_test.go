package orderservice

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOrders(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/orders", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	Router().ServeHTTP(w, req)
	response := w.Result()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Status code is wrong. Have: %d, want: %d", response.StatusCode, http.StatusOK)
	}

	jsonString, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	items := make([]order, 1)
	if err = json.Unmarshal(jsonString, &items); err != nil {
		t.Errorf("Can't parse json response with error %v", err)
	}
}

func TestOrder(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/order/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	Router().ServeHTTP(w, req)
	response := w.Result()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Status code is wrong. Have: %d, want: %d", response.StatusCode, http.StatusOK)
	}

	jsonString, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	
	orderItem := order2{}
	if err = json.Unmarshal(jsonString, &orderItem); err != nil {
		t.Errorf("Can't parse json response with error %v", err)
	}
}
