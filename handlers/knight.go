package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.com/upaphong/go-assignment/domain"
	"gitlab.com/upaphong/go-assignment/engine"
	"gopkg.in/go-playground/validator.v9"
)

func HandleListKnights(e engine.Engine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		knights, err := e.ListKnights()
		if err != nil {
			Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		Json(w, http.StatusOK, knights)
	}
}

func HandleGetKnight(e engine.Engine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		knightID := params["id"]
		knight, err := e.GetKnight(knightID)
		if err != nil {
			Error(w, http.StatusNotFound, err.Error())
			return
		}
		Json(w, http.StatusOK, knight)
	}
}

func HandleSaveKnight(e engine.Engine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var knight domain.Knight
		var err error
		var validate *validator.Validate
		validate = validator.New()
		err = json.NewDecoder(r.Body).Decode(&knight)
		if err != nil {
			Error(w, http.StatusBadRequest, err.Error())
			return
		}
		err = validate.Struct(&knight)
		if err != nil {
			Error(w, http.StatusBadRequest, err.Error())
			return
		}

		result, err := e.Save(&knight)
		if err != nil {
			Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		Json(w, http.StatusCreated, result)
	}
}
