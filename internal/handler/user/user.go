package user

import (
	"encoding/json"
	"hitss/internal/model"
	"hitss/internal/store/user"
	"hitss/pkg/helper/result"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type (
	controller func(w http.ResponseWriter, r *http.Request)

	handler struct {
		GetAllUsers controller
		CreateUser  controller
		UpdateUser  controller
		DeleteUser  controller
	}
)

func New() *handler {
	return &handler{
		GetAllUsers: getAll(),
		CreateUser:  create(),
		UpdateUser:  update(),
		DeleteUser:  delete(),
	}
}

func getAll() controller {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		users, err := user.New().GetAll()
		if err != nil {
			result.JSON(w, result.D{
				"code":  http.StatusBadRequest,
				"error": "internal server error",
			})
			return
		}

		result.JSON(w, result.D{
			"code":  http.StatusOK,
			"users": users,
		})
	}
}

func create() controller {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		data := model.User{}
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

		err = user.New().Create(&data)
		if err != nil {
			result.JSON(w, result.D{
				"code":  http.StatusInternalServerError,
				"error": "internal server error",
			})
			return
		}

		result.JSON(w, result.D{
			"code": http.StatusOK,
			"user": data,
		})
	}
}

func update() controller {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var err error
		data := model.User{}

		data.Id, err = strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			result.JSON(w, result.D{
				"code":  http.StatusBadRequest,
				"error": map[string]string{"id": "cannot be blank"},
			})
			return
		}

		err = json.NewDecoder(r.Body).Decode(&data)
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

		err = user.New().Update(&data)
		if err != nil {
			result.JSON(w, result.D{
				"code":  http.StatusInternalServerError,
				"error": "internal server error",
			})
			return
		}

		result.JSON(w, result.D{
			"code": http.StatusOK,
			"user": data,
		})
	}
}

func delete() controller {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			result.JSON(w, result.D{
				"code":  http.StatusBadRequest,
				"error": map[string]string{"id": "cannot be blank"},
			})
			return
		}

		err = user.New().Delete(id)
		if err != nil {
			result.JSON(w, result.D{
				"code":  http.StatusInternalServerError,
				"error": "internal server error",
			})
			return
		}

		result.JSON(w, result.D{
			"code": http.StatusOK,
			"user": "deleted",
		})
	}
}
