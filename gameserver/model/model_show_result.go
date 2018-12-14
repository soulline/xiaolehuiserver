package cardmodel

import "xiaolehuigo/gameserver/enum"

type ShowResult struct {
	ShowTime       int                 `json:"showTime"`
	ShowValue      []string            `json:"showValue,omitempty"`
	CompareValue   int                 `json:"compareValue"`
	CompareCount   int                 `json:"compareCount"`
	CardTypeStatus enum.CardTypeStatus `json:"cardTypeStatus"`
	ShowPlayer     int                 `json:"showPlayer"`
	IsPass         bool                `json:"isPass"`
}
