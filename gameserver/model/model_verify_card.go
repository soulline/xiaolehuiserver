package cardmodel

import "xiaolehuigo/gameserver/enum"

type VerifyResult struct {
	IsCredit bool                `json:"isCredit"`
	CardType enum.CardTypeStatus `json:"cardType"`
}
