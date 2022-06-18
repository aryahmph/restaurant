package category

import (
	"github.com/aryahmph/restaurant/internal/domain"
)

type (
	Category struct {
		ID   string
		Name string

		Timestamp
	}

	Timestamp = domain.Timestamp
)
