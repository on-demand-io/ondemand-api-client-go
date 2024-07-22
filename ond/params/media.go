package params

type CreateMediaParams struct{}

type FetchMediaParams struct {
	// Sort
	Sort string `url:"sort"`

	// Page number for pagination
	// Default is 1
	Page int `url:"page"`

	// Limit
	// Number of items per page
	// Default is 1
	Limit int `url:"limit"`
}
