package counterservice

import (
	"net/http"
	"testing"
)

func (s *CounterService) TestGetCounter(t *testing.T) {
	r := s.GetCounter()

	if r.Counter != nil && r.Code == http.StatusOK && r.Error == nil {
		t.Log("True usage")
	} else if r.Counter == nil && r.Code == http.StatusInternalServerError && r.Error != nil {
		t.Log("Exception caugth")
	} else {
		t.Error("Failed")
	}
}

func (s *CounterService) TestIncreaseCounter(t *testing.T) {

	inputs := []string{"1", "5", "x"}

	for _, input := range inputs {
		r := s.IncreaseCounter(&input)

		if r.Counter != nil && r.Code == http.StatusOK && r.Error == nil {
			t.Log("True usage")
		} else if r.Counter == nil && r.Code == http.StatusBadRequest && r.Error != nil {
			t.Log("Exception caugth")
		} else if r.Counter == nil && r.Code == http.StatusInternalServerError && r.Error != nil {
			t.Log("Exception caugth")
		} else {
			t.Error("Failed")
		}
	}
}

func (s *CounterService) TestDecreaseCounter(t *testing.T) {

	inputs := []string{"1", "5", "x"}

	for _, input := range inputs {
		r := s.IncreaseCounter(&input)

		if r.Counter != nil && r.Code == http.StatusOK && r.Error == nil {
			t.Log("True usage")
		} else if r.Counter == nil && r.Code == http.StatusBadRequest && r.Error != nil {
			t.Log("Exception caugth")
		} else if r.Counter == nil && r.Code == http.StatusInternalServerError && r.Error != nil {
			t.Log("Exception caugth")
		} else {
			t.Error("Failed")
		}
	}
}

func (s *CounterService) TestResetCounter(t *testing.T) {

	r := s.ResetCounter()

	if r.Counter != nil && r.Code == http.StatusOK && r.Error == nil {
		t.Log("True usage")
	} else if r.Counter == nil && r.Code == http.StatusInternalServerError && r.Error != nil {
		t.Log("Exception caugth")
	} else {
		t.Error("Failed")
	}
}
