package response

import "github.com/gofiber/fiber/v2"

// Meta provides standard metadata for API responses.
type Meta struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	// Additional flexible fields for pagination, etc.
	TotalPages *int `json:"total_pages,omitempty"`
	TotalItems *int `json:"total_items,omitempty"`
}

// Data defines the standard response wrapper.
type Data struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

// Success returns a standardized success JSON response.
func Success(c *fiber.Ctx, statusCode int, message string, data interface{}, metaExtras ...func(*Meta)) error {
	meta := Meta{
		Status:  statusCode,
		Message: message,
	}
	for _, extra := range metaExtras {
		extra(&meta)
	}

	return c.Status(statusCode).JSON(Data{
		Meta: meta,
		Data: data,
	})
}

// Error returns a standardized error JSON response.
func Error(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(Data{
		Meta: Meta{
			Status:  statusCode,
			Message: message,
		},
		Data: nil,
	})
}

// WithPagination is a helper modifier to add pagination info to meta.
func WithPagination(totalPages, totalItems int) func(*Meta) {
	return func(m *Meta) {
		m.TotalPages = &totalPages
		m.TotalItems = &totalItems
	}
}
