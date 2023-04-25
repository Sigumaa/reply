package replygo

import (
	"math/rand"
	"strings"
	"time"
)

// Response represents a single response option for the Reply function.
type Response struct {
	Trigger      []string
	ReplyOptions []string
}

// Bot is the main type for this library.
type Bot struct {
	Responses []Response
}

// NewBot creates a new instance of Bot with the given response options.
func NewBot(responses []Response) *Bot {
	return &Bot{Responses: responses}
}

// GetReply returns a reply based on the input text.
func (b *Bot) GetReply(text string) string {

	for _, response := range b.Responses {
		for _, trigger := range response.Trigger {
			if strings.Contains(text, trigger) {
				return response.ReplyOptions[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(response.ReplyOptions))]
			}
		}
	}

	return "" // If no trigger is found, return empty string
}
