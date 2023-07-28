package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/sushant102004/Traffic-Toll-Microservice/types"
)

const basePrice = 0.04

func main() {
	listenAddr := flag.String("httpAddr", ":3001", "the listen address of http server")

	fmt.Println("Invoicer running on port: 3001")

	invoicer := NewMongoStore()
	http.HandleFunc("/get-invoice", handleGetInvoice(invoicer))
	http.ListenAndServe(*listenAddr, nil)

}

func handleGetInvoice(svc *MongoStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		values, ok := r.URL.Query()["obu"]
		if !ok {
			writeJSON(w, http.StatusNotAcceptable, map[string]string{
				"error": "please provide a valid OBU ID",
			})
			return
		}

		date, ok := r.URL.Query()["date"]
		if !ok {
			writeJSON(w, http.StatusNotAcceptable, map[string]string{
				"error": "please provide a valid date",
			})
			return
		}

		reqDate := date[0]

		obuID, err := strconv.Atoi(values[0])
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
			return
		}
		_ = obuID

		resp, err := svc.GetInvoice(context.Background(), int64(obuID), reqDate)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{
				"error": "OBU ID is invalid.",
			})
			return
		}

		res := types.Invoice{
			OBUID:         obuID,
			TotalAmount:   math.Floor((resp*basePrice)*100) / 100,
			TotalDistance: math.Floor(resp*100) / 100,
			Date:          reqDate,
		}

		writeJSON(w, http.StatusOK, res)
	}
}

func writeJSON(w http.ResponseWriter, status int, body any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(body)
}
