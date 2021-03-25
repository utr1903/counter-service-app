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
		Req: r,
		Db:  c.Base.Db,
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
		Req: r,
		Db:  c.Base.Db,
	}

	code, err := s.IncreaseCounter(dto)

	if err != nil {
		if code == http.StatusBadRequest {
			c.Base.CreateResponse(&w, http.StatusBadRequest, commons.BadRequest(&err))
			return
		} else if code == http.StatusInternalServerError {
			c.Base.CreateResponse(&w, http.StatusInternalServerError, commons.InternalServerError(&err))
			return
		}
	}

	result := commons.Success(nil, nil)
	c.Base.CreateResponse(&w, http.StatusOK, result)
}

// DecreaseCounter : Handler for decreasing the counter
func (c *CounterController) DecreaseCounter(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToString(&w, r)

	s := &counterservice.CounterService{
		Req: r,
		Db:  c.Base.Db,
	}

	code, err := s.DecreaseCounter(dto)

	if err != nil {
		if code == http.StatusBadRequest {
			c.Base.CreateResponse(&w, http.StatusBadRequest, commons.BadRequest(&err))
			return
		} else if code == http.StatusInternalServerError {
			c.Base.CreateResponse(&w, http.StatusInternalServerError, commons.InternalServerError(&err))
			return
		}
	}

	result := commons.Success(nil, nil)
	c.Base.CreateResponse(&w, http.StatusOK, result)
}

// ResetCounter : Handler for resetting the counter
func (c *CounterController) ResetCounter(w http.ResponseWriter, r *http.Request) {

	s := &counterservice.CounterService{
		Req: r,
		Db:  c.Base.Db,
	}

	code, err := s.ResetCounter()

	if err != nil {
		if code == http.StatusBadRequest {
			c.Base.CreateResponse(&w, http.StatusBadRequest, commons.BadRequest(&err))
			return
		} else if code == http.StatusInternalServerError {
			c.Base.CreateResponse(&w, http.StatusInternalServerError, commons.InternalServerError(&err))
			return
		}
	}

	result := commons.Success(nil, nil)
	c.Base.CreateResponse(&w, http.StatusOK, result)
}
