package delivery

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aziz-wahyudin/registration-api/features/user"
	"github.com/aziz-wahyudin/registration-api/utils/helper"
	// "github.com/aziz-wahyudin/registration-api/middlewares"
)

type UserDelivery struct {
	userService user.ServiceInterface
}

func New(service user.ServiceInterface) {
	handler := &UserDelivery{
		userService: service,
	}

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			// handler.Get(w, r)
		case "POST":
			handler.Create(w, r)
		default:
			http.Error(w, fmt.Sprintf("Unsupported method %s", r.Method), http.StatusMethodNotAllowed)
		}
	})
}

func (delivery *UserDelivery) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input UserReq
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err.Error(), "error decode")
		return
	}

	dataCore := toCore(input)
	errCreate := delivery.userService.Create(dataCore)
	if errCreate != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("Failed to create user: %s", err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(helper.SuccessResponse("success created an account"))
}
