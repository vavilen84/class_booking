package handlers

import (
	"encoding/json"
	"github.com/vavilen84/class_booking/containers"
	"github.com/vavilen84/class_booking/helpers"
	"github.com/vavilen84/class_booking/models"
	"github.com/vavilen84/class_booking/store"
	"net/http"
)

func Classes(w http.ResponseWriter, r *http.Request) {
	conn, ctx := store.GetNewDBConn()
	defer conn.Close()

	dec := json.NewDecoder(r.Body)
	apiClasses := containers.APIClasses{}
	err := dec.Decode(&apiClasses)
	if err != nil {
		helpers.LogError(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	class := models.Class{}
	err = class.ValidateAPIClasses(ctx, conn, apiClasses)
	if err != nil {
		helpers.LogError(err)
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	err = class.BatchInsert(ctx, conn, apiClasses)
	if err != nil {
		helpers.LogError(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
