package auth

import (
	"context"
	"github.com/GearFramework/emarket/internal/entities"
	"github.com/GearFramework/emarket/internal/models"
	"github.com/GearFramework/emarket/internal/pkg/db"
)

func (a Auth) Register(ctx context.Context, r *models.RequestRegister) (*entities.Customer, error) {
	if r.Scenario == models.ScenarioEmail {

	}
	customer, err := a.Storage.(*db.Storage).Get("customers", map[string]any{"id": 1})
	if err != nil {
		return nil, err
	}
	return customer.(*entities.Customer), nil
}
