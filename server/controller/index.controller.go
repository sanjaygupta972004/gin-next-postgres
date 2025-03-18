package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/savvy-bit/gin-react-postgres/config"
)

// IndexController is the default controller
type IndexController struct{}

// @Summary Api Version
// @Description Get the api version
// @Tags Api
// @Accept json
// @Produce json
// @Success 200
// @Router /api/version [get]
func (ctrl *IndexController) GetVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": config.Global.Server.Version,
	})
}
