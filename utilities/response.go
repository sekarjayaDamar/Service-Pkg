package utilities

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(writer http.ResponseWriter, code int, data interface{}) {
	writer.WriteHeader(code)
	writer.Header().Add("Content-Type", "application/json")
	parsingError := json.NewEncoder(writer).Encode(data)
	if parsingError != nil {
		panic(parsingError)
	}
}

func WriteSuccessResponse(writer http.ResponseWriter, data interface{}) {
	writer.WriteHeader(http.StatusOK)
	writer.Header().Add("Content-Type", "application/json")
	response := DefaultResponse{
		Code:    http.StatusOK,
		Message: "Success Request Data",
		Data:    data,
	}
	parsingError := json.NewEncoder(writer).Encode(response)
	if parsingError != nil {
		panic(parsingError)
	}
}

type DefaultResponse struct {
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
}
