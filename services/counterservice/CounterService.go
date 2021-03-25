package counterservice

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"
	"time"
)

// CounterService : Implementation of CounterService
type CounterService struct {
	Req *http.Request
	Db  *sql.DB
}

// GetCounter : Returns the current value of counter
func (s *CounterService) GetCounter() (*int, int, error) {

	q := "select counter from counterdb.counter where id = 1"

	var counter *int = nil

	row := s.Db.QueryRow(q)
	err := row.Scan(&counter)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return counter, http.StatusOK, nil
}

// IncreaseCounter : Increases the counter by given number
func (s *CounterService) IncreaseCounter(dto *string) (int, error) {

	increment, err := strconv.Atoi(*dto)
	if err != nil {
		return http.StatusBadRequest, err
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	q := "update counter set counter = counter + ? where id = 1"
	stmt, err := s.Db.PrepareContext(ctx, q)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, increment)
	if err != nil {
		return 500, err
	}

	numRows, err := res.RowsAffected()
	if numRows != 1 || err != nil {
		return 500, err
	}

	return http.StatusOK, nil
}

// DecreaseCounter : Decreases the counter by given number
func (s *CounterService) DecreaseCounter(dto *string) (int, error) {

	decrement, err := strconv.Atoi(*dto)
	if err != nil {
		return http.StatusBadRequest, err
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	q := "update counter set counter = counter - ? where id = 1"
	stmt, err := s.Db.PrepareContext(ctx, q)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, decrement)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	numRows, err := res.RowsAffected()
	if numRows != 1 || err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

// ResetCounter : Resets the counter to zero
func (s *CounterService) ResetCounter() (int, error) {

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	q := "update counter set counter = 0 where id = 1"
	stmt, err := s.Db.PrepareContext(ctx, q)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	numRows, err := res.RowsAffected()
	if numRows != 1 || err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, err
}
