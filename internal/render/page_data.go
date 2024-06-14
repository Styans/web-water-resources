package render

import (
	"test/internal/models"
)

// "test/pkg/forms"

type PageData struct {
	Topic string
	Users []*models.User
	// Form              *forms.Form
	// AuthenticatedUser *models.User
	// Post              *models.Post
	// Posts      []*models.Post
	// Categories []*models.Category
	// Comments   []*models.Comment
}
