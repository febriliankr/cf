package entities

type UpdateBalanceType struct {
	Withdrawal string
	Deposit    string
}

var UpdateBalanceTypes = UpdateBalanceType{
	Withdrawal: "withdrawal",
	Deposit:    "deposit",
}
