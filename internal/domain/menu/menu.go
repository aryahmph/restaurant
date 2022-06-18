package menu

import (
	"database/sql"
	"github.com/aryahmph/restaurant/internal/domain"
)

type (
	Menu struct {
		ID            string
		Name          string
		Price         int
		CategoryName  sql.NullString
		ImageFilename sql.NullString

		Timestamp
		DeletedAt sql.NullTime
	}

	Timestamp = domain.Timestamp
)
