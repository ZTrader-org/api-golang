package platform

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type ztrader struct {
	client    *http.Client
	baseUrl   string
	ApiKey    string
	ApiSecret string
}

var _ Platformer = (*ztrader)(nil)

func NewZtrader(client *http.Client, baseUrl, apiKey, apiSecret string) ztrader {
	return ztrader{
		baseUrl:   baseUrl,
		client:    client,
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
	}
}

func (z ztrader) CreateAccount(ctx context.Context, data CreateAccountRequest) (*CreateAccountResponse, error) {

	url := fmt.Sprintf("%s/broker/api/v1/account", z.baseUrl)

	jsonStr, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	username := z.ApiKey
	password := z.ApiSecret
	req.SetBasicAuth(username, password)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := z.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("status code is not 200, " + res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resData CreateAccountResponse
	if err = json.Unmarshal(body, &resData); err != nil {
		return nil, err
	}

	return &resData, nil
}

func (z ztrader) Deposit(ctx context.Context, login string, data DepositRequest) (*DepositResponse, error) {

	url := fmt.Sprintf("%s/broker/api/v1/deposit/%s", z.baseUrl, login)

	jsonStr, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	username := z.ApiKey
	password := z.ApiSecret
	req.SetBasicAuth(username, password)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := z.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("status code is not 200, " + res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resData DepositResponse
	if err = json.Unmarshal(body, &resData); err != nil {
		return nil, err
	}

	return &resData, nil
}

func (z ztrader) Withdrawal(ctx context.Context, login string, data WithdrawalRequest) (*WithdrawalResponse, error) {

	url := fmt.Sprintf("%s/broker/api/v1/withdrawal/%s", z.baseUrl, login)

	jsonStr, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	username := z.ApiKey
	password := z.ApiSecret
	req.SetBasicAuth(username, password)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := z.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("status code is not 200, " + res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resData WithdrawalResponse
	if err = json.Unmarshal(body, &resData); err != nil {
		return nil, err
	}

	return &resData, nil
}
