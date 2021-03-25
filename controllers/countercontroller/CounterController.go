package controllers

import (
	"net/http"

	"github.com/utr1903/counter-service-app/commons"
	"github.com/utr1903/counter-service-app/services/counterservice"
)

// CounterController : Controller for counter services
type CounterController struct {
	Base *ControllerBase
}

// GetCounter : Handler for getting current counter
func (c *CounterController) GetCounter(w http.ResponseWriter, r *http.Request) {

	s := &counterservice.CounterService{
		Req: r,
		Db:  c.Base.Db,
	}
	counter, err := s.GetCounter()

	if err != nil {
		c.Base.CreateResponse(&w, http.StatusBadRequest, nil)
	}

	result := commons.Success(counter, nil)
	c.Base.CreateResponse(&w, http.StatusOK, result)
}

// IncreaseCounter : Handler for increasing the counter
func (c *CounterController) IncreaseCounter(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToString(&w, r)

	s := &counterservice.CounterService{
		Req: r,
		Db:  c.Base.Db,
	}

	err := s.IncreaseCounter(dto)

	if err != nil {
		c.Base.CreateResponse(&w, http.StatusBadRequest, nil)
	}

	result := commons.Success(true, nil)
	c.Base.CreateResponse(&w, http.StatusOK, result)
}

// UpdateTodoList : Handler for updating an existing list
func (c *TodoListController) UpdateTodoList(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToString(&w, r)

	s := &counterservice.CounterService{
		Req: r,
		Cu: &commons.CommonUtils{
			Db: c.Base.Db,
		},
		Db: c.Base.Db,
	}

	err := s.UpdateTodoList(dto)
	if err != nil {
		c.Base.CreateResponse(&w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(&w, http.StatusOK, nil)
}

// DeleteTodoList : Handler for deleting an existing list
func (c *TodoListController) DeleteTodoList(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToString(&w, r)

	s := &todolistmodule.TodoListService{
		Req: r,
		Cu: &commons.CommonUtils{
			Db: c.Base.Db,
		},
		Db: c.Base.Db,
	}

	err := s.DeleteTodoList(dto)
	if err != nil {
		c.Base.CreateResponse(&w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(&w, http.StatusOK, nil)
}
