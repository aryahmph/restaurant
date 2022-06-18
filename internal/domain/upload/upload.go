package upload

import (
	"github.com/aryahmph/restaurant/internal/domain"
)

type (
	Upload struct {
		ID       string
		Filename string

		Timestamp
	}

	Timestamp = domain.Timestamp
)
