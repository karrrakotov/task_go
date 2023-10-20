package rest

import (
	"encoding/json"
	"net/http"
)

type ResponseError struct {
	Status  int    `json:"status" bson:"status"`
	IsOk    bool   `json:"isOk" bson:"isOk"`
	Message string `json:"message" bson:"message"`
}

type ResponseOk struct {
	Data    interface{} `json:"data" bson:"data"`
	Status  int         `json:"status" bson:"status"`
	IsOk    bool        `json:"isOk" bson:"isOk"`
	Message string      `json:"message" bson:"message"`
}

func ResponseJson(w http.ResponseWriter, status int, structure interface{}) {
	response, err := json.Marshal(structure)
	if err != nil {
		responseError := ResponseError{
			Status:  500,
			IsOk:    false,
			Message: "Ошибка json.Marshal: " + err.Error(),
		}
		ResponseJson(w, 500, responseError)
		return
	}

	w.WriteHeader(status)
	w.Write(response)
}
