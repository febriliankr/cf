package entities

type UpdateBalanceType struct {
	Withdrawal string
	Deposit    string
}

var UpdateBalance = UpdateBalanceType{
	Withdrawal: "withdraw",
	Deposit:    "deposit",
}
