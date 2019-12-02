package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hieuphq/tontine/src/model"
)

// InvestorHandler handler for investor
type InvestorHandler interface {
	CreateInvestor(c *gin.Context)
	UpdateInvestor(c *gin.Context)
	GetInvestorByID(c *gin.Context)
	GetInvestorList(c *gin.Context)
}

func (h *impl) CreateInvestor(c *gin.Context) {
	var dt model.Investor

	if err := c.ShouldBindJSON(&dt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	new, err := h.repo.Investor.Create(ctx, h.store, dt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, new)
}

func (h *impl) UpdateInvestor(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dt model.Investor

	if err := c.ShouldBindJSON(&dt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dt.ID = id
	ctx := c.Request.Context()
	new, err := h.repo.Investor.Update(ctx, h.store, dt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, new)
}

func (h *impl) GetInvestorByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	dt, err := h.repo.Investor.GetByID(ctx, h.store, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dt)
}

func (h *impl) GetInvestorList(c *gin.Context) {
	ctx := c.Request.Context()
	dt, err := h.repo.Investor.GetList(ctx, h.store)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dt)
}
