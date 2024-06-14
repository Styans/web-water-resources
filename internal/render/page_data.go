package render

import (
	"test/internal/models"
)

// "test/pkg/forms"

type PageData struct {
	Topic      string
	Users      []*models.User
	User       *models.User
	Paymants   []models.Payments
	Accruals   []models.AccrualsDTO
	Tariffs    []models.Tariffs
	Tariff     models.Tariffs
	TariffsSum int
	Cost       int
	// Form              *forms.Form
	// AuthenticatedUser *models.User
	// Post              *models.Post
	// Posts      []*models.Post
	// Categories []*models.Category
	// Comments   []*models.Comment
}
