package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/savvy-bit/gin-react-postgres/config"
)

// IndexController is the default controller
type IndexController struct{}

// GetIndex home page
func (ctrl *IndexController) GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "fine. this is homepage",
	})
}

// GetVersion version json
func (ctrl *IndexController) GetVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": config.Global.Server.Version,
	})
}
