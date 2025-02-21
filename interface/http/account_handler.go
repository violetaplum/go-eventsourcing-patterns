package http

import (
	"github.com/google/uuid"
	"net/http"

	"github.com/gin-gonic/gin"
	"go-eventsourcing-patterns/domain"
)

type AccountHandler struct {
	commandService domain.AccountCommandService
	queryService   domain.AccountQueryService
}

func NewAccountHandler(
	commandService domain.AccountCommandService,
	queryService domain.AccountQueryService,
) *AccountHandler {
	return &AccountHandler{
		commandService: commandService,
		queryService:   queryService,
	}
}

func (h *AccountHandler) CreateAccount(c *gin.Context) {
	var req domain.CreateAccountRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accountId := uuid.New().String()

	cmd := domain.CreateAccountCommand{
		InitialBalance: req.InitialBalance,
		UserName:       req.UserName,
		AccountId:      accountId,
	}
	if err := h.commandService.CreateAccount(c, cmd); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	account, err := h.queryService.GetAccountByID(c, accountId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}

func (h *AccountHandler) Deposit(c *gin.Context) {
	var req domain.DepositRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmd := domain.DepositCommand{
		AccountID: req.AccountID,
		Amount:    req.Amount,
	}

	if err := h.commandService.Deposit(c, cmd); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	account, err := h.queryService.GetAccountByID(c, req.AccountID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}

func (h *AccountHandler) Withdraw(c *gin.Context) {
	var req domain.WithdrawRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmd := domain.WithdrawCommand{
		AccountID: req.AccountId,
		Amount:    req.Amount,
	}

	if err := h.commandService.Withdraw(c, cmd); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	account, err := h.queryService.GetAccountByID(c, req.AccountId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}

func (h *AccountHandler) GetAccount(c *gin.Context) {
	req := domain.GetAccountRequest{}
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account, err := h.queryService.GetAccountByID(c, req.AccountId)
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

	res := map[string]interface{}{}

	if len(accounts) == 0 {
		res["list"] = []string{}
	} else {
		res["list"] = accounts
	}

	c.JSON(http.StatusOK, res)
}

func (h *AccountHandler) GetHealthCheck(c *gin.Context) {
	if h == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "un healthy"})
	}

	res := map[string]interface{}{
		"status": "ok",
	}
	c.JSON(http.StatusOK, res)
}

// SetupRoutes Gin 라우터 설정
func (h *AccountHandler) SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.POST("/account.create", h.CreateAccount)
		v1.GET("/account.list", h.ListAccounts)
		v1.GET("/account.info", h.GetAccount)
		v1.POST("/account.deposit", h.Deposit)
		v1.POST("/account.withdraw", h.Withdraw)
		v1.GET("/_healthz", h.GetHealthCheck)
	}
}
