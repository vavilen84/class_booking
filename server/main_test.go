package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/class_booking/constants"
	"github.com/vavilen84/class_booking/containers"
	"github.com/vavilen84/class_booking/handlers"
	"github.com/vavilen84/class_booking/helpers"
	"github.com/vavilen84/class_booking/models"
	"github.com/vavilen84/class_booking/store"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

/**
 * ! IMPORTANT - dont use for production DB !
 */
func prepareTestDB(ctx context.Context, conn *sql.Conn) {
	dropAllTablesFromTestDB(ctx, conn)
	err := models.CreateMigrationsTableIfNotExists(ctx, conn)
	if err != nil {
		log.Println(err)
	}

	err = models.MigrateUp(ctx, conn)
	if err != nil {
		log.Println(err)
	}

	models.LoadFixtures(ctx, conn)
	return
}

/**
 * ! IMPORTANT - dont use for production DB !
 */
func dropAllTablesFromTestDB(ctx context.Context, conn *sql.Conn) {
	tables := []string{
		constants.MigrationsTableName,
		constants.VisitorTimetableItemTableName,
		constants.VisitorTableName,
		constants.TimetableItemTableName,
		constants.ClassTableName,
	}
	for i := 0; i < len(tables); i++ {
		_, err := conn.ExecContext(ctx, "DROP TABLE IF EXISTS "+tables[i])
		if err != nil {
			log.Println(err)
		}
	}
}

func setTestAppEnv() {
	err := os.Setenv("APP_ENV", "test")
	if err != nil {
		helpers.LogError(err)
	}
}

func createHTTPRequest(handler http.Handler, url string, b *strings.Reader) (resp *http.Response, bodyString string) {
	req := httptest.NewRequest(http.MethodPost, url, b)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	resp = w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	bodyString = string(body)
	return
}

func TestClasses(t *testing.T) {
	setTestAppEnv()
	store.InitTestDB()
	conn, ctx := store.GetNewTestDBConn()
	prepareTestDB(ctx, conn)
	handler := handlers.MakeHandler()

	resp, body := createHTTPRequest(handler, "/classes", strings.NewReader(``))
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Contains(t, body, "Bad Request")

	a := containers.APIClasses{}
	reqBody, err := json.Marshal(a)
	assert.Nil(t, err)
	resp, body = createHTTPRequest(handler, "/classes", strings.NewReader(string(reqBody)))
	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
	assert.Contains(t, body, fmt.Sprintf(constants.RequiredErrorMsg, constants.APIClassesStructName, "StartDate"))
	assert.Contains(t, body, fmt.Sprintf(constants.RequiredErrorMsg, constants.APIClassesStructName, "Name"))
	assert.Contains(t, body, fmt.Sprintf(constants.RequiredErrorMsg, constants.APIClassesStructName, "Capacity"))
	assert.Contains(t, body, fmt.Sprintf(constants.RequiredErrorMsg, constants.APIClassesStructName, "EndDate"))

	startDate := time.Now().AddDate(0, 0, 2)
	endDate := time.Now().AddDate(0, 0, 3)
	capacity := 10
	name := "New Class"
	a = containers.APIClasses{
		Name:      name,
		StartDate: &startDate,
		EndDate:   &endDate,
		Capacity:  &capacity,
	}
	reqBody, err = json.Marshal(a)
	assert.Nil(t, err)
	resp, body = createHTTPRequest(handler, "/classes", strings.NewReader(string(reqBody)))
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	c := models.Class{}
	err = c.FindByName(ctx, conn, name)
	assert.Nil(t, err)
	assert.NotEmpty(t, c.Id)
	assert.Equal(t, name, c.Name)
	assert.Equal(t, capacity, *c.Capacity)

	ti := models.TimetableItem{}
	err = ti.FindByDate(ctx, conn, &startDate)
	assert.Nil(t, err)
	assert.NotEmpty(t, ti.Id)
	assert.Equal(t, c.Id, ti.ClassId)

	ti = models.TimetableItem{}
	err = ti.FindByDate(ctx, conn, &endDate)
	assert.Nil(t, err)
	assert.NotEmpty(t, ti.Id)
	assert.Equal(t, c.Id, ti.ClassId)
}
