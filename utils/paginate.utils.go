package utils

import (
	"dating-sim/res"
	"net/http"
	"strconv"
)

func Paginate(w http.ResponseWriter, r *http.Request) (size int, page int) {
	rSize := r.URL.Query().Get("size")
	if rSize == "" {
		rSize = "10"
	}
	rPage := r.URL.Query().Get("page")
	if rPage == "" {
		rPage = "0"
	}
	size, err := strconv.Atoi(rSize)
	if err != nil {
		res.ResErrJson(w, http.StatusBadRequest, err)
		return
	}
	page, err = strconv.Atoi(rPage)
	if err != nil {
		res.ResErrJson(w, http.StatusBadRequest, err)
		return
	}
	return size, page
}
