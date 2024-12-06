package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/korg/atlas/internal/config"
	"github.com/korg/atlas/internal/infra/http/dtos"
	"github.com/korg/atlas/internal/infra/http/errors"
	"go.uber.org/zap"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New().String()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Request-ID", requestID)

	var reqBody *dtos.CreateUserInputDto

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		config.Logger().Error(
			fmt.Sprintf("..: [%s] - error decoding request body", requestID),
			zap.String("error", err.Error()),
		)

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.NewGenericRestResponseError(
			requestID,
			http.StatusBadRequest,
		))

		return
	}

	json.NewEncoder(w).Encode(reqBody)
}
