package infrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JsonResponse(res interface{}, w http.ResponseWriter) {
	response, err := json.Marshal(res)
	CheckError(err)

	fmt.Fprint(w, string(response))
}
