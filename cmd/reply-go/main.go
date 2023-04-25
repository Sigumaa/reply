package main

import (
	r "github.com/Sigumaa/reply"
)

func main() {
	bot := r.NewBot([]r.Response{
		{
			Trigger:      []string{"what's up", "what's new"},
			ReplyOptions: []string{"Nothing much!", "Not much!", "The usual!"},
		},
	})

	println(bot.GetReply("what's up"))

	//You can also add more responses later
	bot.Responses = append(bot.Responses, r.Response{
		Trigger:      []string{"how are you"},
		ReplyOptions: []string{"I'm fine!", "I'm good!", "I'm doing well!"},
	})

	println(bot.GetReply("how are you"))

	unchi := r.Response{
		Trigger:      []string{"うんち"},
		ReplyOptions: []string{"きたねえ"},
	}

	bot.Responses = append(bot.Responses, unchi)
	println(bot.GetReply("うんち"))

	// 後からreplyoptionを追加
	bot.Responses[2].ReplyOptions = append(bot.Responses[2].ReplyOptions, "うんち")

	println(bot.GetReply("うんち"))
}
