package response

type SysCaptchaResponse struct {
	CaptchaId         string `json:"captchaId"`
	PicPath           string `json:"picPath"`
	CaptchaLength     int    `json:"captchaLength"`
	OpenCaptcha       bool   `json:"openCaptcha"`
	HasInitial        bool   `json:"hasInitial"`
	AutoHrefToInitial bool   `json:"autoHrefToInitial"`
}
