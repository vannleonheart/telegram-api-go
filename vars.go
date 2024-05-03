package telegram

var (
	DefaultParseMode = ParseModeHtml
)

type Client struct {
	Config *Config
	token  string
}

type Config struct {
	BaseUrl string `json:"base_url"`
	Token   string `json:"token"`
}

type SendMessageRequest struct {
	ChatId    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

type SendMessageResponse struct {
	Ok          bool               `json:"ok"`
	ErrorCode   *int               `json:"error_code"`
	Description *string            `json:"description"`
	Result      *SendMessageResult `json:"result"`
}

type SendMessageResult struct {
	MessageId int64  `json:"message_id"`
	From      From   `json:"from"`
	Chat      Chat   `json:"chat"`
	Date      int64  `json:"date"`
	Text      string `json:"text"`
}

type From struct {
	Id        int64  `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
}

type Chat struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}
