package infrastructure

import "net/http"

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func ControllerErrorResponder(err error, w http.ResponseWriter, status int) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
	}
}
