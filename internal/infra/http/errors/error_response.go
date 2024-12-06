package errors

import "time"

type RestResponseError struct {
	CorrelationID string    `json:"correlationId"`
	Status        int       `json:"status"`
	Message       string    `json:"message"`
	Timestamp     time.Time `json:"timestamp"`
}

func NewGenericRestResponseError(id string, status int) *RestResponseError {
	return &RestResponseError{
		CorrelationID: id,
		Status:        status,
		Message:       "Internal server error, contact administrator for support",
		Timestamp:     time.Now(),
	}
}

func NewRestResponseError(id string, status int, message string) *RestResponseError {
	return &RestResponseError{
		CorrelationID: id,
		Status:        status,
		Message:       message,
		Timestamp:     time.Now(),
	}
}
