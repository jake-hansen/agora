package adapter

import (
	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/domain"
)

// PagreRequestDTOToDomain conversts the given PageRequest from DTO representation to domain representation.
func PageRequestDTOToDomain(pageReq *dto.PageRequest) *domain.PageRequest {
	req := &domain.PageRequest{
		PageSize:      pageReq.PageSize,
		RequestedPage: pageReq.RequestedPage,
	}
	return req
}
