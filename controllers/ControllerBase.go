package controllers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/utr1903/counter-service-app/commons"
)

// Controller : Base class for controllers
type ControllerBase struct {
	Db *sql.DB
}

// ParseRequestToString : Generic way to parse JSON request to string
func (c *ControllerBase) ParseRequestToString(w *http.ResponseWriter, r *http.Request) *string {

	dto, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.CreateResponse(w, http.StatusBadRequest, commons.BadRequest(&err))
		return nil
	}
	defer r.Body.Close()

	dtoString := string(dto)

	return &dtoString
}

// CreateResponse : Generic way for all controllers to create JSON response
func (c *ControllerBase) CreateResponse(w *http.ResponseWriter, code int, result *commons.CustomResult) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(code)
	json.NewEncoder(*w).Encode(result)
}
