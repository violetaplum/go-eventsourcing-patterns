package http

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go-eventsourcing-patterns/domain"
	"go-eventsourcing-patterns/domain/mock"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

//func setupTestRouter(t *testing.T) (*gin.Engine, *mock.MockAccountCommandService, *mock.MockAccountQueryService) {
//	gin.SetMode(gin.TestMode)
//	ctrl := gomock.NewController(t)
//	mockCommandService := mock.NewMockAccountCommandService(ctrl)
//	mockQueryService := mock.NewMockAccountQueryService(ctrl)
//
//	handler := NewAccountHandler(
//		mockCommandService,
//		mockQueryService,
//	)
//	router := gin.New()
//	handler.SetupRoutes(router)
//	return router, mockCommandService, mockQueryService
//}

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
			CreateAccount(gomock.Any(), gomock.Any()).Return(nil)

		mockQueryService.EXPECT().
			GetAccountByID(gomock.Any(), gomock.Any()).Return(&domain.Account{
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

		var response domain.Account
		err := json.Unmarshal(resp.Body.Bytes(), &response)
		assert.NoError(t, err)
		numberStr := strconv.Itoa(int(rand.Float64()))
		assert.Equal(t, "test_user_"+numberStr, response.UserName)
		assert.Equal(t, int64(1000), response.Balance)
	})
}
