package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/hieuphq/tontine/src/model"
)

// GroupHandler handler for group
type GroupHandler interface {
	CreateGroup(c *gin.Context)
	UpdateGroup(c *gin.Context)
	GetGroupList(c *gin.Context)
	GetGroupDetail(c *gin.Context)
	AddInvestorIntoGroup(c *gin.Context)
	InvestorTopup(c *gin.Context)
	RemoveInvestorFromGroup(c *gin.Context)
	// WithdrawProfit(c *gin.Context)
	UpdateBalance(c *gin.Context)
	GetGroupLogs(c *gin.Context)
	// Close(c *gin.Context)
}

func (h *impl) CreateGroup(c *gin.Context) {
	var dt model.Group

	if err := c.ShouldBindJSON(&dt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	new, err := h.repo.Group.Create(ctx, h.store, dt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, new)
}

func (h *impl) UpdateGroup(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dt model.Group

	if err := c.ShouldBindJSON(&dt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dt.ID = id
	ctx := c.Request.Context()
	new, err := h.repo.Group.Update(ctx, h.store, dt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, new)
}

func (h *impl) GetGroupList(c *gin.Context) {
	ctx := c.Request.Context()
	dt, err := h.repo.Group.GetList(ctx, h.store)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dt)
}

func (h *impl) GetGroupDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	dt, err := h.repo.Group.GetDetailByID(ctx, h.store, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dt)
}

type addInvestorRequest struct {
	GroupID    int64
	InvestorID int64   `json:"investor_id" binding:"required"`
	Amount     float64 `json:"amount" binding:"required"`
	Currency   string  `json:"currency" binding:"oneof=VND USD"`
}

func (h *impl) AddInvestorIntoGroup(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var air addInvestorRequest

	if err := c.BindJSON(&air); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	air.GroupID = id
	gi := model.GroupInvestor{
		GroupID:    air.GroupID,
		InvestorID: air.InvestorID,
		Amount:     air.Amount,
		Currency:   air.Currency,
	}

	ctx := c.Request.Context()

	// Flow
	// - check investor existed in system
	// - check investor have joined group
	// - accept join group
	// - log data into group_log
	_, err = h.repo.Investor.GetByID(ctx, h.store, gi.InvestorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "investor is not existed"})
		return
	}
	existed, err := h.repo.Group.ExistedInvestor(ctx, h.store, gi.GroupID, gi.InvestorID)
	if err == nil && existed != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "investor is joined"})
		return
	}

	dt, err := h.repo.Group.AddInvestor(ctx, h.store, gi)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	existedG, err := h.repo.Group.GetByID(ctx, h.store, id)
	if err != nil && existedG == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "group is not existed"})
		return
	}

	if existedG.Currency != "" && existedG.Currency != air.Currency {
		c.JSON(http.StatusBadRequest, gin.H{"error": "group currency is NOT matched " + existedG.Currency})
		return
	}
	existedG.Amount = existedG.Amount + air.Amount
	_, err = h.repo.Group.Update(ctx, h.store, *existedG)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable update group"})
		return
	}

	msg := fmt.Sprintf("Investor %v joined Group", gi.InvestorID)
	h.repo.ActivityLog.LogGroup(ctx, h.store, model.GroupLog{
		Name:     msg,
		GroupID:  gi.GroupID,
		Amount:   gi.Amount,
		Currency: gi.Currency,
	})

	c.JSON(http.StatusOK, dt)
}

type topupRequest struct {
	GroupID    int64
	InvestorID int64
	Amount     float64 `json:"amount" binding:"required"`
	Currency   string  `json:"currency" binding:"oneof=VND USD"`
}

func (h *impl) InvestorTopup(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	invtIDStr := c.Param("investor_id")
	invtID, err := strconv.ParseInt(invtIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var tr topupRequest

	if err := c.BindJSON(&tr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tr.GroupID = id
	tr.InvestorID = invtID

	ctx := c.Request.Context()
	store, finally := h.store.BeginTx(ctx)

	// Flow
	// - check investor existed in system
	// - check investor have joined group
	// - accept join group
	// - log data into group_log
	existed, err := h.repo.Group.ExistedInvestor(ctx, store, tr.GroupID, tr.InvestorID)
	if err != nil && existed == nil {
		finally(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "investor is not existed"})
		return
	}

	if existed.Currency != "" && existed.Currency != tr.Currency {
		finally(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "currency is NOT matched " + existed.Currency})
		return
	}
	existed.Amount = existed.Amount + tr.Amount

	dt, err := h.repo.Group.UpdateInvestor(ctx, store, *existed)
	if err != nil {
		finally(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	existedG, err := h.repo.Group.GetByID(ctx, store, tr.GroupID)
	if err != nil && existedG == nil {
		finally(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "group is not existed"})
		return
	}

	if existedG.Currency != "" && existedG.Currency != tr.Currency {
		finally(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "group currency is NOT matched " + existedG.Currency})
		return
	}
	existedG.Amount = existedG.Amount + tr.Amount
	_, err = h.repo.Group.Update(ctx, store, *existedG)
	if err != nil {
		finally(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable update group"})
		return
	}

	msg := fmt.Sprintf("Investor %v top up %v (%v)", tr.InvestorID, tr.Amount, tr.Currency)
	_, err = h.repo.ActivityLog.LogGroup(ctx, store, model.GroupLog{
		Name:     msg,
		GroupID:  tr.GroupID,
		Amount:   tr.Amount,
		Currency: tr.Currency,
	})

	if err != nil {
		finally(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable log group"})
		return
	}

	finally(err)
	c.JSON(http.StatusOK, dt)
}

func (h *impl) RemoveInvestorFromGroup(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	invtIDStr := c.Param("investor_id")
	invtID, err := strconv.ParseInt(invtIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	store, finally := h.store.BeginTx(ctx)

	// Flow
	// - check investor existed in system
	// - check investor have joined group
	// - accept join group
	// - log data into group_log
	existed, err := h.repo.Group.ExistedInvestor(ctx, store, id, invtID)
	if err != nil && existed == nil {
		finally(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "investor is not existed"})
		return
	}

	amount := existed.Amount
	currency := existed.Currency

	err = h.repo.Group.FarawellInvestor(ctx, store, invtID, id)
	if err != nil {
		finally(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	existedG, err := h.repo.Group.GetByID(ctx, store, id)
	if err != nil && existedG == nil {
		finally(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "group is not existed"})
		return
	}

	if existedG.Currency != "" && existedG.Currency != currency {
		finally(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "group currency is NOT matched " + existedG.Currency})
		return
	}
	existedG.Amount = existedG.Amount - amount
	_, err = h.repo.Group.Update(ctx, store, *existedG)
	if err != nil {
		finally(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable update group"})
		return
	}

	msg := fmt.Sprintf("Investor %v is out from group %v with %v (%v)", invtID, id, amount, currency)
	log, err := h.repo.ActivityLog.LogGroup(ctx, store, model.GroupLog{
		Name:     msg,
		GroupID:  id,
		Amount:   -amount,
		Currency: currency,
	})

	if err != nil {
		finally(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable log group"})
		return
	}

	finally(err)
	c.JSON(http.StatusOK, log)
}

type updateBalanceRequest struct {
	GroupID  int64
	Amount   float64 `json:"amount" binding:"ne=0"`
	Currency string  `json:"currency" binding:"oneof=VND USD"`
}

func (h *impl) UpdateBalance(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var tr updateBalanceRequest

	if err := c.BindJSON(&tr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tr.GroupID = id

	ctx := c.Request.Context()
	store, finally := h.store.BeginTx(ctx)

	// Flow
	// - check investor existed in system

	existedG, err := h.repo.Group.GetByID(ctx, store, id)
	if err != nil && existedG == nil {
		finally(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "group is not existed"})
		return
	}

	if existedG.Currency != "" && existedG.Currency != tr.Currency {
		finally(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "group currency is NOT matched " + existedG.Currency})
		return
	}

	total := existedG.Amount

	existedG.Amount = existedG.Amount + tr.Amount
	_, err = h.repo.Group.Update(ctx, store, *existedG)
	if err != nil {
		finally(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable update group"})
		return
	}
	msg := fmt.Sprintf("Group %v Update balance from %v to %v (%v)", tr.GroupID, total, total+tr.Amount, tr.Currency)
	log, err := h.repo.ActivityLog.LogGroup(ctx, store, model.GroupLog{
		Name:     msg,
		GroupID:  id,
		Amount:   tr.Amount,
		Currency: tr.Currency,
	})
	if err != nil {
		finally(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable log group"})
		return
	}

	invts, err := h.repo.Group.InvestorList(ctx, store, tr.GroupID)
	if err != nil {
		finally(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable get investors"})
		return
	}

	if len(invts) < 0 {
		finally(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "investors is empty"})
		return
	}

	currTotal := calculateTotalInGroup(invts)
	totalPercent := float64(0)
	for idx := range invts {
		itm := invts[idx]
		p := itm.Amount / currTotal
		if idx == len(invts)-1 {
			p = 1 - totalPercent
		}
		totalPercent = totalPercent + p
		itm.Amount = itm.Amount + tr.Amount*p

		_, err := h.repo.Group.UpdateInvestor(ctx, store, itm)
		if err != nil {
			finally(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "unable update investor"})
			return
		}

	}

	finally(err)
	c.JSON(http.StatusOK, log)
}

func calculateTotalInGroup(invts []model.GroupInvestor) float64 {
	total := float64(0)
	for idx := range invts {
		itm := invts[idx]
		total = total + itm.Amount
	}

	return total
}

func (h *impl) GetGroupLogs(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	rs, err := h.repo.ActivityLog.GetGroupLogs(ctx, h.store, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	c.JSON(http.StatusOK, rs)

}
