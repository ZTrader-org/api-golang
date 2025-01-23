package platform

import "time"

type Account struct {
	ID                    string     `json:"id"`
	Name                  string     `json:"name"`
	Email                 string     `json:"email"`
	Login                 int        `json:"login"`
	Balance               float64    `json:"balance"`
	Credit                float64    `json:"credit"`
	Equity                float64    `json:"equity"`
	Margin                float64    `json:"margin"`
	MarginLevel           float64    `json:"margin_level"`
	Free                  float64    `json:"free"`
	City                  string     `json:"city"`
	Address               string     `json:"address"`
	Phone                 string     `json:"phone"`
	IDNumber              string     `json:"id_number"`
	Comment               string     `json:"comment"`
	TaxRate               int        `json:"tax_rate"`
	LeadSource            string     `json:"lead_source"`
	ReadOnly              bool       `json:"read_only"`
	AllowToChangePassword bool       `json:"allow_to_change_password"`
	SendReports           bool       `json:"send_reports"`
	EnableOneTimePassword bool       `json:"enable_one_time_password"`
	Demo                  bool       `json:"demo"`
	SmartStopOut          float64    `json:"smart_stop_out"`
	Leverage              int        `json:"leverage"`
	UnrealizedNetPL       float64    `json:"unrealized_net_pl"`
	RealizedNetPL         float64    `json:"realized_net_pl"`
	IsOnline              bool       `json:"is_online"`
	CallMargin            bool       `json:"call_margin"`
	CallMarginAt          *time.Time `json:"call_margin_at"`
	CreatedAt             time.Time  `json:"created_at"`
	UpdatedAt             *time.Time `json:"updated_at"`
}

type Transaction struct {
	ID              string    `json:"id"`
	TransactionType string    `json:"transaction_type"`
	Amount          float64   `json:"amount"`
	Comment         string    `json:"comment"`
	CheckFreeMargin bool      `json:"check_free_margin"`
	Time            time.Time `json:"time"`
}
