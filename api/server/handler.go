package server

import (
	"MyWallet/api/dblayer"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerInterface interface {
	GetOK(c *gin.Context)
	GetFixedCosts(c *gin.Context)
}

type Handler struct {
	db dblayer.DBLayer
}

func NewHandler() (HandlerInterface, error) {
	db, err := dblayer.NewORM("postgres", "gorm:gorm@/mywallet")
	if err != nil {
		return nil, err
	}
	return &Handler{
		db: db,
	}, nil
}

func (h *Handler) GetOK(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"hello": "world"})
}

func (h *Handler) GetFixedCosts(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database server error"})
		return
	}
}
