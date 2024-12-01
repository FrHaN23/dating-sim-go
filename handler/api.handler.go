package handler

import (
	"dating-sim/res"
	"dating-sim/types"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	data := types.Res{
		Message: "im alive",
	}
	res.ResOkJSON(w, &data)
}
