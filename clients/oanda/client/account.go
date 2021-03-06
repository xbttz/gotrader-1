package oandacl

import (
	"encoding/json"
	"time"
)

type Accounts struct {
	Accounts []AccountProperties `json:"accounts"`
}

type AccountProperties struct {
	ID string `json:"id"`
}

type AccountSummaryResponse struct {
	Account AccountSummary `json:"account"`
}

type AccountSummary struct {
	NAV                         float64   `json:"NAV,string"`
	Alias                       string    `json:"alias"`
	Balance                     float64   `json:"balance,string"`
	CreatedByUserID             int       `json:"createdByUserID"`
	CreatedTime                 time.Time `json:"createdTime"`
	Currency                    string    `json:"currency"`
	HedgingEnabled              bool      `json:"hedgingEnabled"`
	ID                          string    `json:"id"`
	LastTransactionID           string    `json:"lastTransactionID"`
	MarginAvailable             float64   `json:"marginAvailable,string"`
	MarginCloseoutMarginUsed    float64   `json:"marginCloseoutMarginUsed,string"`
	MarginCloseoutNAV           float64   `json:"marginCloseoutNAV,string"`
	MarginCloseoutPercent       float64   `json:"marginCloseoutPercent,string"`
	MarginCloseoutPositionValue float64   `json:"marginCloseoutPositionValue,string"`
	MarginCloseoutUnrealizedPL  float64   `json:"marginCloseoutUnrealizedPL,string"`
	MarginRate                  float64   `json:"marginRate,string"`
	MarginUsed                  float64   `json:"marginUsed,string"`
	OpenPositionCount           int       `json:"openPositionCount"`
	OpenTradeCount              int       `json:"openTradeCount"`
	PendingOrderCount           int       `json:"pendingOrderCount"`
	Pl                          float64   `json:"pl,string"`
	PositionValue               float64   `json:"positionValue,string"`
	ResettablePL                float64   `json:"resettablePL,string"`
	UnrealizedPL                float64   `json:"unrealizedPL,string"`
	WithdrawalLimit             float64   `json:"withdrawalLimit,string"`
}

type AccountInstruments struct {
	Instruments []Instrument `json:"instruments"`
}

type Instrument struct {
	DisplayName                 string  `json:"displayName"`
	DisplayPrecision            int     `json:"displayPrecision"`
	MarginRate                  float64 `json:"marginRate,string"`
	MaximumOrderUnits           float64 `json:"maximumOrderUnits,string"`
	MaximumPositionSize         float64 `json:"maximumPositionSize,string"`
	MaximumTrailingStopDistance float64 `json:"maximumTrailingStopDistance,string"`
	MinimumTradeSize            float64 `json:"minimumTradeSize,string"`
	MinimumTrailingStopDistance float64 `json:"minimumTrailingStopDistance,string"`
	Name                        string  `json:"name"`
	PipLocation                 int     `json:"pipLocation"`
	TradeUnitsPrecision         int     `json:"tradeUnitsPrecision"`
	Type                        string  `json:"type"`
}

func (c *OandaClient) GetAccounts() (Accounts, error) {

	endpoint := "/accounts"

	response, err := c.get(endpoint)

	if err != nil {
		return Accounts{}, nil
	}

	data := Accounts{}
	err = json.Unmarshal(response, &data)

	if err != nil {
		return Accounts{}, err
	}

	return data, nil
}

func (c *OandaClient) GetAccountSummary(accountID string) (AccountSummaryResponse, error) {

	endpoint := "/accounts/" + accountID + "/summary"

	response, err := c.get(endpoint)

	if err != nil {
		return AccountSummaryResponse{}, nil
	}

	data := AccountSummaryResponse{}
	err = json.Unmarshal(response, &data)

	if err != nil {
		return AccountSummaryResponse{}, err
	}

	return data, nil
}

func (c *OandaClient) GetAccountInstruments(accountID string) (AccountInstruments, error) {

	endpoint := "/accounts/" + accountID + "/instruments"

	response, err := c.get(endpoint)

	if err != nil {
		return AccountInstruments{}, err
	}

	data := AccountInstruments{}
	err = json.Unmarshal(response, &data)

	if err != nil {
		return AccountInstruments{}, err
	}

	return data, nil
}
