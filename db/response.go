package db

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"gopkg.in/mgo.v2/bson"
)

//Close close http
func Close(resp *http.Request) {

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			LogError(err)
		}
	}()

}

func LogError(err interface{}) {
	_, file, line, _ := runtime.Caller(1)

	log.Printf("[cgl] debug %s:%d", file, line)
	log.Print("ErrorLog==>", err)
}

// RespondJSON makes the response with payload as json format
func RespondJSON(r *http.Request, w http.ResponseWriter, message string, status int, payload interface{}) {
	go Close(r)
	response, err := json.Marshal(bson.M{"status": status, "message": message, "data": payload})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

// RespondError makes the error response with payload as json format
func RespondError(r *http.Request, w http.ResponseWriter, code int, message string) {
	RespondJSON(r, w, message, code, map[string]string{"error": message})
}

//RespondJSONSocket socket response
func RespondJSONSocket(message string, status int, payload interface{}) string {
	response, err := json.Marshal(bson.M{"status": status, "message": message, "data": payload})
	if err != nil {
		fmt.Print("jsn")
	}
	return string(response)
}

//RespondJSONSocketMessage socket response
func RespondJSONSocketMessage(payload interface{}) string {
	response, err := json.Marshal(bson.M{"data": payload})
	if err != nil {
		fmt.Print("jsn")
	}
	return string(response)
}
