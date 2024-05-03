### TELEGRAM API

#### Installation
```go
go get -u github.com/vannleonheart/telegram-api-go
```

#### Config
```go
telegramConfig := telegram.Config {
    BaseUrl: "https://api.telegram.org",
    Token:   "{your_bot_api_token}",
}
```
#### Create Client
```go
telegramClient := telegram.New(&telegramConfig)
```
#### Set Token Manually
```go
token := "{your_other_bot_api_token}"

telegramClient = telegramClient.WithToken(token)
```
#### Send Chat Message
```go
chatId := "{target_chat_id}"
message := "{your_message}"
parseMode := telegram.ParseModeHtml

result, err := telegramClient.SendMessage(chatId, message, &parseMode)

if err !=nil {
	// handle error
}

if !result.Ok {
	// handle error response from telegram
}

fmt.Println(result.Result.MessageId)
```