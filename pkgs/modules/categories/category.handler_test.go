package categories

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockService struct{}

func (ms *MockService) All(conds ...interface{}) []*Category {
	return []*Category{{NameAr: "test_ar", NameEn: "test_en", Cover: "test_cover"}}
}

// Implement similar mock methods for GetSingleCategoryWithImages, GetAll, Find, Update, DeleteById, Store, and DeleteImageByCategoryId

func TestIndex(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockService := &MockService{}
	handler := CategoryHandler{service: mockService}

	router.GET("/categories", handler.Index)

	req, _ := http.NewRequest(http.MethodGet, "/categories", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "test_ar")
}

// Repeat similar tests for Create, Store, Edit, Update, Delete, and DeleteCategoryImage
