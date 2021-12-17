package infrastructure

import (
	"encoding/json"
	"fmt"
	"gospy/models"
	"net/http"
)

func JsonResponse(res interface{}, w http.ResponseWriter) {
	response, err := json.Marshal(res)
	CheckError(err)

	fmt.Fprint(w, string(response))
}

func ParseBody(model models.Model, w http.ResponseWriter, request *http.Request) models.Model {
	err := json.NewDecoder(request.Body).Decode(&model)

	ControllerErrorResponder(err, w, http.StatusBadGateway)

	return model
}
