package countercontroller

import (
	"net/http"

	"github.com/utr1903/counter-service-app/commons"
	"github.com/utr1903/counter-service-app/controllers"
	"github.com/utr1903/counter-service-app/services/counterservice"
)

// CounterController : Controller for counter services
type CounterController struct {
	Base *controllers.ControllerBase
}

// GetCounter : Handler for getting current counter
func (c *CounterController) GetCounter(w http.ResponseWriter, r *http.Request) {

	s := &counterservice.CounterService{
		Db: c.Base.Db,
	}

	response := s.GetCounter()

	if response.Error != nil {
		if response.Code == http.StatusBadRequest {
			c.Base.CreateResponse(&w, http.StatusBadRequest, commons.BadRequest(&response.Error))
			return
		} else if response.Code == http.StatusInternalServerError {
			c.Base.CreateResponse(&w, http.StatusInternalServerError, commons.InternalServerError(&response.Error))
			return
		}
	}

	result := commons.Success(response.Counter, nil)
	c.Base.CreateResponse(&w, http.StatusOK, result)
}

// IncreaseCounter : Handler for increasing the counter
func (c *CounterController) IncreaseCounter(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToString(&w, r)

	s := &counterservice.CounterService{
		Db: c.Base.Db,
	}

	response := s.IncreaseCounter(dto)

	if response.Error != nil {
		if response.Code == http.StatusBadRequest {
			c.Base.CreateResponse(&w, http.StatusBadRequest, commons.BadRequest(&response.Error))
			return
		} else if response.Code == http.StatusInternalServerError {
			c.Base.CreateResponse(&w, http.StatusInternalServerError, commons.InternalServerError(&response.Error))
			return
		}
	}

	result := commons.Success(response.Counter, response.Error)
	c.Base.CreateResponse(&w, http.StatusOK, result)
}

// DecreaseCounter : Handler for decreasing the counter
func (c *CounterController) DecreaseCounter(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToString(&w, r)

	s := &counterservice.CounterService{
		Db: c.Base.Db,
	}

	response := s.DecreaseCounter(dto)

	if response.Error != nil {
		if response.Code == http.StatusBadRequest {
			c.Base.CreateResponse(&w, http.StatusBadRequest, commons.BadRequest(&response.Error))
			return
		} else if response.Code == http.StatusInternalServerError {
			c.Base.CreateResponse(&w, http.StatusInternalServerError, commons.InternalServerError(&response.Error))
			return
		}
	}

	result := commons.Success(response.Counter, response.Error)
	c.Base.CreateResponse(&w, http.StatusOK, result)
}

// ResetCounter : Handler for resetting the counter
func (c *CounterController) ResetCounter(w http.ResponseWriter, r *http.Request) {

	s := &counterservice.CounterService{
		Db: c.Base.Db,
	}

	response := s.ResetCounter()

	if response.Error != nil {
		if response.Code == http.StatusBadRequest {
			c.Base.CreateResponse(&w, http.StatusBadRequest, commons.BadRequest(&response.Error))
			return
		} else if response.Code == http.StatusInternalServerError {
			c.Base.CreateResponse(&w, http.StatusInternalServerError, commons.InternalServerError(&response.Error))
			return
		}
	}

	result := commons.Success(response.Counter, response.Error)
	c.Base.CreateResponse(&w, http.StatusOK, result)
}
