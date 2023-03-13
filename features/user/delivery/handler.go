package delivery

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aziz-wahyudin/registration-api/features/user"
	"github.com/aziz-wahyudin/registration-api/utils/helper"
)

type UserDelivery struct {
	userService user.ServiceInterface
}

func New(service user.ServiceInterface) {
	handler := &UserDelivery{
		userService: service,
	}

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, fmt.Sprintf("Unsupported method %s", r.Method), http.StatusMethodNotAllowed)
		} else {
			handler.Create(w, r)
		}
	})
	http.HandleFunc("/login/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, fmt.Sprintf("Unsupported method %s", r.Method), http.StatusMethodNotAllowed)
		} else {
			handler.Login(w, r)
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

func (delivery *UserDelivery) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input UserReq
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err.Error(), "error decode")
		return
	}

	dataUser, token, errLogin := delivery.userService.Login(input.Email, input.Password)
	if errLogin != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("Failed to login: %s", err.Error())
		return
	}

	data := map[string]interface{}{
		"name":  dataUser.Name,
		"token": token,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.SuccessWithDataResponse("success login", data))
}
