package models

// Counter : CounterService response model
type CounterResponse struct {
	Counter *int
	Code    int
	Error   error
}
