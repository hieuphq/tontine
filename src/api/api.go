package api

import (
	"github.com/gin-gonic/gin"

	"github.com/hieuphq/tontine/src/config"
	"github.com/hieuphq/tontine/src/interfaces/repo/sqlite"
)

// InitRouter api handlers
func InitRouter(cfg config.Config) {
	s := sqlite.NewStore("./bin/db.db")
	rep := sqlite.NewRepo()

	h := NewHandlers(s, rep)

	r := gin.Default()

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// Investor
	r.GET("/investors", h.GetInvestorList)
	r.GET("/investors/:id", h.GetInvestorByID)
	r.POST("/investors", h.CreateInvestor)
	r.PUT("/investors/:id", h.UpdateInvestor)

	// Groups
	r.GET("/groups", h.GetGroupList)
	r.GET("/groups/:id", h.GetGroupDetail)
	r.POST("/groups", h.CreateGroup)
	r.PUT("/groups/:id", h.UpdateGroup)

	r.PUT("/groups/:id/balances", h.UpdateBalance)

	r.POST("/groups/:id/investors", h.AddInvestorIntoGroup)
	r.PUT("/groups/:id/investors/:investor_id/topup", h.InvestorTopup)
	r.DELETE("/groups/:id/investors/:investor_id", h.RemoveInvestorFromGroup)

	r.GET("/groups/:id/logs", h.GetGroupLogs)

	r.Run(cfg.Port)
}
