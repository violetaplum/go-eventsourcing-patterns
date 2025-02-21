package http

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go-eventsourcing-patterns/domain"
	"go-eventsourcing-patterns/domain/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAccountService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("CreateAccount", func(t *testing.T) {
		mockCommandService := mock.NewMockAccountCommandService(ctrl)
		mockQueryService := mock.NewMockAccountQueryService(ctrl)
		gin.SetMode(gin.TestMode)

		handler := NewAccountHandler(mockCommandService, mockQueryService)
		router := gin.New()
		handler.SetupRoutes(router)

		mockCommandService.EXPECT().
			CreateAccount(gomock.Any(), gomock.Any()).
			Return(nil)

		mockQueryService.EXPECT().
			GetAccountByID(gomock.Any(), gomock.Any()).
			Return(&domain.AccountResponse{
				ID:       "test-id",
				Balance:  1000,
				UserName: "Test User",
			}, nil)

		reqBody := domain.CreateAccountRequest{
			InitialBalance: 1000,
			UserName:       "Test User",
		}
		jsonData, _ := json.Marshal(reqBody)

		req := httptest.NewRequest("POST", "/v1/account.create", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Code)

		var response domain.AccountResponse

		err := json.Unmarshal(resp.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "Test User", response.UserName)
		assert.Equal(t, int64(1000), response.Balance)
	})

	t.Run("Deposit", func(t *testing.T) {
		mockCommandService := mock.NewMockAccountCommandService(ctrl)
		mockQueryService := mock.NewMockAccountQueryService(ctrl)

		handler := NewAccountHandler(mockCommandService, mockQueryService)
		router := gin.New()
		handler.SetupRoutes(router)

		mockCommandService.EXPECT().
			Deposit(gomock.Any(), gomock.Any()).
			Return(nil)

		mockQueryService.EXPECT().
			GetAccountByID(gomock.Any(), gomock.Any()).
			Return(&domain.AccountResponse{
				ID:       "test-id",
				Balance:  2000,
				UserName: "Test User",
			}, nil)

		reqBody := domain.DepositRequest{
			Amount:    1000,
			AccountID: "account_id",
		}
		jsonData, _ := json.Marshal(reqBody)
		req := httptest.NewRequest("POST", "/v1/account.deposit", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Code)

		var response domain.AccountResponse
		err := json.Unmarshal(resp.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.NotEmpty(t, response.ID)
		assert.NotEmpty(t, response.UserName)
		assert.NotEmpty(t, response.Balance)
	})

	t.Run("Withdraw", func(t *testing.T) {
		mockCommandService := mock.NewMockAccountCommandService(ctrl)
		mockQueryService := mock.NewMockAccountQueryService(ctrl)

		handler := NewAccountHandler(mockCommandService, mockQueryService)
		router := gin.New()
		handler.SetupRoutes(router)

		mockCommandService.EXPECT().
			Withdraw(gomock.Any(), gomock.Any()).
			Return(nil)

		mockQueryService.EXPECT().
			GetAccountByID(gomock.Any(), gomock.Any()).
			Return(&domain.AccountResponse{
				ID:       "test-id",
				Balance:  1000,
				UserName: "Test User",
			}, nil)

		reqBody := domain.WithdrawRequest{
			Amount:    1000,
			AccountId: "account_id",
		}
		jsonData, _ := json.Marshal(reqBody)
		req := httptest.NewRequest("POST", "/v1/account.withdraw", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Code)

		var response domain.AccountResponse
		err := json.Unmarshal(resp.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.NotEmpty(t, response.ID)
		assert.NotEmpty(t, response.UserName)
		assert.Equal(t, int64(1000), response.Balance)
	})
}
