package jsonapi

import (
	"encoding/json"
	"net/http"
)

type apiError struct {
	Err    string
	Status int
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			if e, ok := err.(apiError); ok {
				writeJSON(w, e.Status, e)
				return
			}
			writeJSON(w, http.StatusInternalServerError, apiError{Err: "Internal Server Error"})
		}
	}
}

func (e apiError) Error() string {
	return e.Err
}

func Setup() {
	http.HandleFunc("/user", makeHTTPHandler(handleGetUserByID))
	http.ListenAndServe(":3000", nil)
}

type User struct {
	ID    int
	Valid bool
}

func handleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return writeJSON(w, http.StatusMethodNotAllowed, apiError{Err: "invalid method", Status: http.StatusMethodNotAllowed})
	}

	user := User{}

	if !user.Valid {
		return apiError{Err: "not authorised", Status: http.StatusUnauthorized}
	}

	return writeJSON(w, http.StatusOK, User{})
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
