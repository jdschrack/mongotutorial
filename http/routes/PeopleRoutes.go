package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	Data "github.com/jdschrack/mongotutorial/data"
	"github.com/jdschrack/mongotutorial/repository"
	Repository "github.com/jdschrack/mongotutorial/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	var person Data.Person

	json.NewDecoder(request.Body).Decode(&person)
	result, err := repository.AddPerson(person)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		log.Fatal(err)
		return
	}

	json.NewEncoder(response).Encode(result)
}

func GetPeopleEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	people, err := Repository.GetAllPeople()

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		log.Fatal(err)
		return
	}

	json.NewEncoder(response).Encode(people)
}

func GetPersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		log.Fatal(err)
		return
	}

	person, err := Repository.GetPerson(id)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		log.Fatal(err)
		return
	}

	json.NewEncoder(response).Encode(person)
}