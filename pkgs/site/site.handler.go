package site

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SiteHandler struct {
	db *gorm.DB
}

func NewSiteHandler(db *gorm.DB) *SiteHandler {
	return &SiteHandler{db: db}

}

func (handler *SiteHandler) index(context *gin.Context) {
	
	context.HTML(http.StatusOK, "site.index.gohtml", gin.H{})
}

func (handler *SiteHandler) about(context *gin.Context) {
	context.HTML(http.StatusOK, "site.about.gohtml", gin.H{})
}
func (handler *SiteHandler) category(context *gin.Context) {
	context.HTML(http.StatusOK, "site.category.gohtml", gin.H{})
}
func (handler *SiteHandler) contact(context *gin.Context) {
	context.HTML(http.StatusOK, "site.contact.gohtml", gin.H{})
}
func (handler *SiteHandler) gallery(context *gin.Context) {
	context.HTML(http.StatusOK, "site.gallery.gohtml", gin.H{})
}
func (handler *SiteHandler) services(context *gin.Context) {
	context.HTML(http.StatusOK, "site.services.gohtml", gin.H{})
}
func (handler *SiteHandler) pill(context *gin.Context) {
	context.HTML(http.StatusOK, "site.pill.gohtml", gin.H{})
}
