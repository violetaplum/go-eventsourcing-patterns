package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-eventsourcing-patterns/application/command"
	"go-eventsourcing-patterns/application/query"
	"go-eventsourcing-patterns/domain"
)

type AccountHandler struct {
	commandService *command.AccountCommandService
	queryService   *query.AccountQueryService
}

func NewAccountHandler(
	commandService *command.AccountCommandService,
	queryService *query.AccountQueryService,
) *AccountHandler {
	return &AccountHandler{
		commandService: commandService,
		queryService:   queryService,
	}
}

func (h *AccountHandler) CreateAccount(c *gin.Context) {
	var req domain.CreateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmd := domain.CreateAccountCommand{
		InitialBalance: req.InitialBalance,
	}

	if err := h.commandService.CreateAccount(c, cmd); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *AccountHandler) Deposit(c *gin.Context) {
	accountID := c.Param("id")

	var req domain.DepositRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmd := domain.DepositCommand{
		AccountID: accountID,
		Amount:    req.Amount,
	}

	if err := h.commandService.Deposit(c, cmd); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *AccountHandler) Withdraw(c *gin.Context) {
	accountID := c.Param("id")

	var req domain.WithdrawRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmd := domain.WithdrawCommand{
		AccountID: accountID,
		Amount:    req.Amount,
	}

	if err := h.commandService.Withdraw(c, cmd); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *AccountHandler) GetAccount(c *gin.Context) {
	accountID := c.Param("id")

	account, err := h.queryService.GetAccountByID(c, accountID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}

func (h *AccountHandler) ListAccounts(c *gin.Context) {
	accounts, err := h.queryService.ListAccounts(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, accounts)
}

// SetupRoutes Gin 라우터 설정
func (h *AccountHandler) SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.POST("/accounts", h.CreateAccount)
		v1.GET("/accounts", h.ListAccounts)
		v1.GET("/accounts/:id", h.GetAccount)
		v1.POST("/accounts/:id/deposit", h.Deposit)
		v1.POST("/accounts/:id/withdraw", h.Withdraw)
	}
}
