package swagger

import (
	"hitss/internal/store/swagger"
	"net/http"
)

type (
	controller func(w http.ResponseWriter, r *http.Request)

	handler struct {
		GetJsonOpenAPI controller
		GetYamlOpenAPI controller
	}
)

func New() *handler {
	return &handler{
		GetJsonOpenAPI: getJSON(),
		GetYamlOpenAPI: getYAML(),
	}
}

func getJSON() controller {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		data, err := swagger.New().GetJSON()
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}

func getYAML() controller {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		data, err := swagger.New().GetYAML()
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/x-yaml")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}
