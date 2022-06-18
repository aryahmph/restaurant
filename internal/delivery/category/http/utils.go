package http

import (
	httpCommon "github.com/aryahmph/restaurant/common/http"
	domainCategory "github.com/aryahmph/restaurant/internal/domain/category"
)

func (h HTTPCategoryDelivery) mapCategoryBodyToDomain(payload httpCommon.Category) domainCategory.Category {
	category := domainCategory.Category{
		ID:   payload.ID,
		Name: payload.Name,
	}
	return category
}

func (h HTTPCategoryDelivery) mapCategoryDomainToResponse(c domainCategory.Category) httpCommon.Category {
	return httpCommon.Category{
		ID:        c.ID,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}
