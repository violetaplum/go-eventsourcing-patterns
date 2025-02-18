package http

import (
	"github.com/gin-gonic/gin"
	"go-eventsourcing-patterns/domain"
	"go-eventsourcing-patterns/domain/mock"
)

func setupTestRouter() (*gin.Engine, *mock.MockAccountCommandService, *mock.MockAccountQueryService) {
	gin.SetMode(gin.TestMode)
	mockCommandService := new(mock.MockAccountCommandService)
	mockQueryService := new(mock.MockAccountQueryService)

	handler := NewAccountHandler(
		(*domain.AccountCommandService)(mockCommandService),
		(*domain.AccountQueryService)(mockQueryService),
	)

}
