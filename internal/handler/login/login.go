package login

import (
	"encoding/json"
	"hitss/internal/model"
	"hitss/internal/store/login"
	"hitss/pkg/helper/result"
	"net/http"
)

type (
	controller func(w http.ResponseWriter, r *http.Request)

	handler struct {
		SignInUser controller
	}
)

func New() *handler {
	return &handler{
		SignInUser: signIn(),
	}
}

func signIn() controller {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		data := model.Login{}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			result.JSON(w, result.D{
				"code":  http.StatusBadRequest,
				"error": "incorrect data structure",
			})
			return
		}

		if err := data.Validate(); err != nil {
			result.JSON(w, result.D{
				"code":  http.StatusBadRequest,
				"error": err,
			})
			return
		}

		token, err := login.New().SignIn(data)
		if err != nil {
			result.JSON(w, result.D{
				"code":  http.StatusInternalServerError,
				"error": "internal server error",
			})
			return
		}

		if token == "" {
			result.JSON(w, result.D{
				"code":  http.StatusBadRequest,
				"error": "user and password invalid",
			})
			return
		}

		result.JSON(w, result.D{
			"code":  http.StatusOK,
			"token": token,
		})
	}
}
