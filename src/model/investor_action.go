package model

// InvestorAction Investor's action
type InvestorAction string

const (
	// InvestorActionJoin join to a group
	InvestorActionJoin InvestorAction = "join"

	// InvestorActionLeave leave from a group
	InvestorActionLeave InvestorAction = "leave"

	// InvestorActionWithdraw withdraw profit
	InvestorActionWithdraw InvestorAction = "withdraw"

	// InvestorActionDeposit deposit profit
	InvestorActionDeposit InvestorAction = "deposit"
)
