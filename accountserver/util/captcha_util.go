package util

import "github.com/dchest/captcha"

func VerifyCaptcha(captchaId string, value string) bool {
	return captcha.VerifyString(captchaId, value)
}
