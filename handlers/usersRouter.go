package handlers

import (
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

func UsersRouter(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path + "/")

	if path == "/users" {
		switch r.Method {
		case http.MethodGet:
			return
		case http.MethodPost:
			return
		default:
			postError(w, http.StatusMethodNotAllowed)
			return
		}
	}

	path = strings.TrimPrefix(path, "/users/")
	if !bson.IsObjectIdHex(path) {
		postError(w, http.StatusNotFound)
		return
	}

	// id := bson.ObjectIdHex(path)

	switch r.Method {
	case http.MethodGet:
		return
	case http.MethodPut:
		return
	case http.MethodPatch:
		return
	case http.MethodDelete:
		return
	default:
		postError(w, http.StatusMethodNotAllowed)
		return

	}

}
