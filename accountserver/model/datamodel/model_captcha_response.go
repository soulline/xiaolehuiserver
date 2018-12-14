package datamodel

type CaptchaResponse struct {
	CaptchaId string `json:"captchaId"`

	ImageUrl string `json:"imageUrl"`
}
