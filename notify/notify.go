package notify

import (
	"encoding/base64"
	"github.com/go-resty/resty/v2"
	"strings"
)

func ToQYWechat(webhook, msg string) {
	if webhook == "" {
		return
	}
	body := `
		{
			"msgtype": "markdown",
			"markdown": {
				"content": "_CONTENT_"
			}
		}
		`
	body = strings.ReplaceAll(body, "_CONTENT_", msg)
	_, _ = resty.New().R().SetHeader("Content-Type", "application/json").SetBody(body).Post(webhook)
}

func ToQYWechatBase64(webhook, msg string) {
	if webhook == "" {
		return
	}
	msg = base64.StdEncoding.EncodeToString([]byte(msg))
	body := `
		{
			"msgtype": "markdown",
			"markdown": {
				"content": "_CONTENT_"
			}
		}
		`
	body = strings.ReplaceAll(body, "_CONTENT_", msg)
	_, _ = resty.New().R().SetHeader("Content-Type", "application/json").SetBody(body).Post(webhook)
}
