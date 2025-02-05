package http

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-eventsourcing-patterns/application/command"
	"go-eventsourcing-patterns/application/query"
	"go-eventsourcing-patterns/domain"
	"net/http"
)

type AccountHandler struct {
	commandService *command.AccountCommandService
	queryService   *query.AccountQueryService
}

func NewAccountHandler(commandService *command.AccountCommandService,
	queryService *query.AccountQueryService) *AccountHandler {
	return &AccountHandler{
		commandService: commandService,
		queryService:   queryService,
	}
}

func (h *AccountHandler) CreateAccount(c *gin.Context) {
	var req domain.CreateAccountRequest
	if err := c.ShouldBindQuery()
}
