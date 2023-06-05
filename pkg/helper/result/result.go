package result

import (
	"encoding/json"
	"hitss/pkg/helper/logger"
	"net/http"
)

type D map[string]interface{}
type CtxKey struct{}

func JSON(w http.ResponseWriter, data D) {
	result, err := json.Marshal(data)
	if err != nil {
		logger.Write(err)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(data["code"].(int))
	w.Write(result)
}
