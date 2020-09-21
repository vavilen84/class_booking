package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/vavilen84/class_booking/containers"
	"github.com/vavilen84/class_booking/helpers"
	"github.com/vavilen84/class_booking/models"
	"github.com/vavilen84/class_booking/store"
	"net/http"
)

func Bookings(w http.ResponseWriter, r *http.Request) {
	conn, ctx := store.GetNewDBConn()
	defer conn.Close()

	dec := json.NewDecoder(r.Body)
	apiBookings := containers.APIBookings{}
	err := dec.Decode(&apiBookings)
	if err != nil {
		helpers.LogError(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	vti := models.VisitorTimetableItem{}
	err = vti.ValidateAPIBookings(ctx, conn, apiBookings)
	if err != nil {
		helpers.LogError(err)
		http.Error(w, "Unprocessable Entity", http.StatusUnprocessableEntity)
		return
	}

	v := models.Visitor{}
	err = v.FindByEmail(ctx, conn, apiBookings.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			helpers.LogError(err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		helpers.LogError(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	ti := models.TimetableItem{}
	err = ti.FindByDate(ctx, conn, apiBookings.Date)
	if err != nil {
		if err == sql.ErrNoRows {
			helpers.LogError(err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		helpers.LogError(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	vti = models.VisitorTimetableItem{
		VisitorId:       v.Id,
		TimetableItemId: ti.Id,
	}
	err = vti.Insert(ctx, conn)
	if err != nil {
		helpers.LogError(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
