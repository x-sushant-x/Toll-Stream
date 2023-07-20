package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/sushant102004/Traffic-Toll-Microservice/types"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "the listen address of http server")
	flag.Parse()

	store := NewMemoryStore()
	var (
		svc = NewInvoiceAggregator(store)
	)
	makeHTTP_Transport(*listenAddr, svc)
}

func makeHTTP_Transport(listenAddr string, svc *InvoiceAggregator) {
	fmt.Println("HTTP Transport Running on Port", listenAddr)
	http.HandleFunc("/aggregate", handleAggregate(svc))
	http.ListenAndServe(listenAddr, nil)
}

func handleAggregate(svc *InvoiceAggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var distance types.CalculatedDistance
		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		fmt.Println(distance)

		if err := svc.AggregateDistance(distance); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, body any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(body)
}
