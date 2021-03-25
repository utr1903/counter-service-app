package counterservice

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
)

// CounterService : Implementation of CounterService
type CounterService struct {
	Req *http.Request
	Db  *sql.DB
}

// GetCounter : Returns the current value of counter
func (s *CounterService) GetCounter() (*int, error) {

	q := "select counter from counter where id = 1"

	var counter *int = nil

	row := s.Db.QueryRow(q, "counter")
	err := row.Scan(&counter)

	if err != nil {
		return nil, err
	}

	return counter, nil
}

// IncreaseCounter : Increases the counter by given number
func (s *CounterService) IncreaseCounter(dto *string) error {

	var increment *int = nil
	json.Unmarshal([]byte(*dto), &increment)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	q := "update counter set counter = counter + ? where id = 1"
	stmt, err := s.Db.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, increment)
	if err != nil {
		return err
	}

	numRows, err := res.RowsAffected()
	if numRows != 1 || err != nil {
		return err
	}

	return nil
}

// DecreaseCounter : Decreases the counter by given number
func (s *CounterService) DecreaseCounter(dto *string) error {

	var decrement *int = nil
	json.Unmarshal([]byte(*dto), &decrement)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	q := "update counter set counter = counter - ? where id = 1"
	stmt, err := s.Db.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, decrement)
	if err != nil {
		return err
	}

	numRows, err := res.RowsAffected()
	if numRows != 1 || err != nil {
		return err
	}

	return nil
}

// ResetCounter : Resets the counter to zero
func (s *CounterService) ResetCounter() error {

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	q := "update counter set counter = 0 where id = 1"
	stmt, err := s.Db.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx)
	if err != nil {
		return err
	}

	numRows, err := res.RowsAffected()
	if numRows != 1 || err != nil {
		return err
	}

	return nil
}
