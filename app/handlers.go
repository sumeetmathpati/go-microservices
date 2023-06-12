package app

import (
	"banking/service"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func getAllCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["id"])
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Posted...")
}
