package platform

import "context"

type Platformer interface {
	CreateAccount(ctx context.Context, data CreateAccountRequest) (*CreateAccountResponse, error)
	Deposit(ctx context.Context, login string, data DepositRequest) (*DepositResponse, error)
	Withdrawal(ctx context.Context, login string, data WithdrawalRequest) (*WithdrawalResponse, error)
}

type CreateAccountResponse struct {
	Account Account `json:"account"`
	Message string  `json:"message"`
}

type DepositResponse struct {
	Transaction Transaction `json:"transaction"`
	Account     Account     `json:"account"`
	Message     string      `json:"message"`
}

type WithdrawalResponse struct {
	Transaction Transaction `json:"transaction"`
	Account     Account     `json:"account"`
	Message     string      `json:"message"`
}

type CreateAccountRequest struct {
	GroupName             string `json:"group_name"` // is required
	Name                  string `json:"name"`
	Email                 string `json:"email"`
	Login                 string `json:"login"`
	Password              string `json:"password"`
	City                  string `json:"city"`
	Address               string `json:"address"`
	Phone                 string `json:"phone"`
	IDNumber              string `json:"id_number"`
	Comment               string `json:"comment"`
	TaxRate               int    `json:"tax_rate"`
	LeadSource            string `json:"lead_source"`
	ReadOnly              bool   `json:"read_only"`
	AllowToChangePassword bool   `json:"allow_to_change_password"`
	SendReports           bool   `json:"send_reports"`
	EnableOneTimePassword bool   `json:"enable_one_time_password"`
	Demo                  bool   `json:"demo"`
	Leverage              int    `json:"leverage"`
}

type DepositRequest struct {
	Amount          float64 `json:"amount"` // is required
	Comment         string  `json:"comment"`
	ShowToUser      bool    `json:"show_to_user"`
	CheckFreeMargin bool    `json:"check_free_margin"`
}

type WithdrawalRequest struct {
	Amount          float64 `json:"amount"` // is required
	Comment         string  `json:"comment"`
	ShowToUser      bool    `json:"show_to_user"`
	CheckFreeMargin bool    `json:"check_free_margin"`
}
