package telegram

import (
	"errors"
	"fmt"
	"github.com/vannleonheart/goutil"
)

func New(config *Config) *Client {
	return &Client{Config: config}
}

func (c *Client) WithToken(token string) *Client {
	c.token = token

	return c
}

func (c *Client) SendMessage(to, message string, parseMode *string) (*SendMessageResponse, error) {
	token := c.getToken()
	if len(token) == 0 {
		return nil, errors.New("token is empty")
	}

	targetUrl := fmt.Sprintf("%s/bot%s/sendMessage", c.Config.BaseUrl, token)

	var result SendMessageResponse

	useParseMode := DefaultParseMode
	if parseMode != nil && len(*parseMode) > 0 {
		useParseMode = *parseMode
	}

	reqBody := SendMessageRequest{
		ChatId:    to,
		Text:      message,
		ParseMode: useParseMode,
	}

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	_, err := goutil.SendHttpPost(targetUrl, reqBody, &reqHeaders, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) getToken() string {
	if len(c.token) > 0 {
		return c.token
	}

	if c.Config != nil && len(c.Config.Token) > 0 {
		return c.Config.Token
	}

	return ""
}
