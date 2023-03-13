package delivery

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aziz-wahyudin/registration-api/features/participant"
	"github.com/aziz-wahyudin/registration-api/middlewares"
	"github.com/aziz-wahyudin/registration-api/utils/helper"
)

type ParticipantDelivery struct {
	participantService participant.ServiceInterface
}

func New(service participant.ServiceInterface) {
	handler := &ParticipantDelivery{
		participantService: service,
	}
	http.HandleFunc("/participants", middlewares.JWTMiddleware(handler.Create))
	http.HandleFunc("/participants", func(w http.ResponseWriter, r *http.Request) {
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

func (delivery *ParticipantDelivery) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	roleToken := middlewares.ExtractTokenUserRole(r)
	if roleToken != "participant" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(helper.FailedResponse("can only register if you are a participant"))
		return
	}

	photoUrl, errUrl := helper.UploadFile(r, "photo")
	if errUrl != nil {
		http.Error(w, errUrl.Error(), http.StatusBadRequest)
		return
	}

	var input ParticipantReq
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err.Error(), "error decode")
		return
	}

	input.Photo = photoUrl

	dataCore := toCore(input)
	errCreate := delivery.participantService.Create(dataCore)
	if errCreate != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("Failed to register: %s", err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(helper.SuccessResponse("success register"))
}
