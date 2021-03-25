package counterservice

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/utr1903/counter-service-app/services/counterservice/models"
)

// CounterService : Implementation of CounterService
type CounterService struct {
	Req *http.Request
	Db  *sql.DB
}

// GetCounter : Returns the current value of counter
func (s *CounterService) GetCounter() *models.CounterResponse {

	q := "select counter from counterdb.counter where id = 1"

	var counter *int = nil

	row := s.Db.QueryRow(q)
	err := row.Scan(&counter)

	if err != nil {
		return createResponse(nil, http.StatusInternalServerError, err)
	}

	result := createResponse(counter, http.StatusOK, nil)
	return result
}

// IncreaseCounter : Increases the counter by given number
func (s *CounterService) IncreaseCounter(dto *string) *models.CounterResponse {

	increment, err := strconv.Atoi(*dto)
	if err != nil {
		return createResponse(nil, http.StatusBadRequest, err)
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	q := "update counter set counter = counter + ? where id = 1"
	stmt, err := s.Db.PrepareContext(ctx, q)
	if err != nil {
		return createResponse(nil, http.StatusInternalServerError, err)
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, increment)
	if err != nil {
		return createResponse(nil, http.StatusInternalServerError, err)
	}

	numRows, err := res.RowsAffected()
	if numRows != 1 || err != nil {
		return createResponse(nil, http.StatusInternalServerError, err)
	}

	return s.GetCounter()
}

// DecreaseCounter : Decreases the counter by given number
func (s *CounterService) DecreaseCounter(dto *string) *models.CounterResponse {

	decrement, err := strconv.Atoi(*dto)
	if err != nil {
		return createResponse(nil, http.StatusBadRequest, err)
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	q := "update counter set counter = counter - ? where id = 1"
	stmt, err := s.Db.PrepareContext(ctx, q)
	if err != nil {
		return createResponse(nil, http.StatusInternalServerError, err)
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, decrement)
	if err != nil {
		return createResponse(nil, http.StatusInternalServerError, err)
	}

	numRows, err := res.RowsAffected()
	if numRows != 1 || err != nil {
		return createResponse(nil, http.StatusInternalServerError, err)
	}

	return s.GetCounter()
}

// ResetCounter : Resets the counter to zero
func (s *CounterService) ResetCounter() *models.CounterResponse {

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	q := "update counter set counter = 0 where id = 1"
	stmt, err := s.Db.PrepareContext(ctx, q)
	if err != nil {
		return createResponse(nil, http.StatusInternalServerError, err)
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx)
	if err != nil {
		return createResponse(nil, http.StatusInternalServerError, err)
	}

	numRows, err := res.RowsAffected()
	if numRows != 1 || err != nil {
		return createResponse(nil, http.StatusInternalServerError, err)
	}

	ctr := 0
	return createResponse(&ctr, http.StatusOK, nil)
}

func createResponse(counter *int, code int, err error) *models.CounterResponse {
	result := &models.CounterResponse{
		Counter: counter,
		Code:    code,
		Error:   err,
	}

	return result
}
