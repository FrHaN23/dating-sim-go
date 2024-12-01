package res

import (
	"dating-sim/types"
	"encoding/json"
	"net/http"
)

func ResOkJSON(w http.ResponseWriter, data *types.Res) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&data)
}

func ResErrJson(w http.ResponseWriter, status int, err error) {
	data := types.Res{
		Message: err.Error(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&data)
}

func PreFlightJson(w http.ResponseWriter, status int, allowedMethod string) {
	headers := w.Header()
	headers.Add("Access-Control-Allow-Methods", allowedMethod)
	headers.Add("Access-Control-Allow-Origin", "*")
	headers.Add("Access-Control-Max-Age", "86400") // 24 Hours
	headers.Add("Access-Control-Allow-Headers", "Content-Type, x-requested-with, Origin, Accept	")
	headers.Add("Vary", "Origin")
	headers.Add("Vary", "Access-Control-Request-Method")
	headers.Add("Vary", "Access-Control-Request-Headers")
	w.WriteHeader(status)
}
