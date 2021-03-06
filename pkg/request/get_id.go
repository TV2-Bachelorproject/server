package request

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetID(r *http.Request) (uint, error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		return 0, err
	}

	return uint(id), nil
}
