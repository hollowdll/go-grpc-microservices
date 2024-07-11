package api

import (
	"context"

	"github.com/hollowdll/go-grpc-microservices/services/payment/internal/application/core/domain"
)

// Application contains core business logic.
type Application struct {
	// Here we can add outbound ports for e.g. database
	// that business logic can call.
}

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) Charge(ctx context.Context, payment domain.Payment) (domain.Payment, error) {
	// Some business logic
	// e.g. call 3rd party payment gateway and save payment to database

	return payment, nil
}
