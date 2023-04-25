<samp>

reply-go
========

simple go library for replying to messages

install
-------

`go get github.com/Sigumaa/reply`

example
-------

```go
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
    
	hello := r.Response{
		Trigger:      []string{"hello", "hi"},
		ReplyOptions: []string{"Hello!", "Hi!", "Hey!"},
	}

	bot.Responses = append(bot.Responses, hello)

	println(bot.GetReply("hello"))
}
```

output:

```bash
$ go run main.go
The usual!
Hey!
```

</samp>