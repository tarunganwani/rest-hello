package router

import(
	"net/http"
	"encoding/json"
)

func SendJsonResponse(w http.ResponseWriter, code int, payload interface{}){
	payloadJsonBytes, _ := json.Marshal(payload)
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	w.Write(payloadJsonBytes)
}

func SendErrorResponse(w http.ResponseWriter, code int, errormsg string){
	SendJsonResponse(w, code, map[string]string{"error":errormsg})
}

func SendOkResponse(w http.ResponseWriter){
	SendJsonResponse(w, http.StatusOK, map[string]string{"result":"success"})
}