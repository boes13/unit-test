package http_handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/boes13/unit-test/data/mysql"
	redis_data "github.com/boes13/unit-test/data/redis"
)

func GetUserEmail(w http.ResponseWriter, r *http.Request) {
	uID := r.FormValue("user_id")
	if uID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("user_id not found"))
		return
	}

	userID, err := strconv.ParseInt(uID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("invalid user_id %s", uID)))
		return
	}

	// retrieve from cache/redis
	email, err := redis_data.GetUserEmail(userID)
	if err != nil {
		if err.Error() == "redigo: nil returned" {
			// if there's none in cache, query database
			var exist bool
			email, exist, err = mysql.GetUserEmail(userID)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			if !exist {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("no such id found on system"))
				return
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(email))
}
