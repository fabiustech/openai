# openai
[![GoDoc](http://img.shields.io/badge/GoDoc-Reference-blue.svg)](https://godoc.org/github.com/fabiustech/openai)
[![Go Report Card](https://goreportcard.com/badge/github.com/sashabaranov/go-gpt3)](https://goreportcard.com/report/github.com/fabiustech/openai)

Zero dependency Go Client for [OpenAI](https://beta.openai.com/) API endpoints. Built upon the great work done [here](https://github.com/sashabaranov/go-gpt3).

Installation:
```
go get github.com/fabiustech/openai
```

Example usage:

```go
package main

import (
	"context"
	"fmt"

	"github.com/fabiustech/openai"
	"github.com/fabiustech/openai/models"
)

func main() {
	var c = openai.NewClient("your token")
	
	var resp, err = c.CreateCompletion(context.Background(), &openai.CompletionRequest{
		Model:     models.TextDavinci003,
		MaxTokens: 100,
		Prompt:    "Lorem ipsum",
	})
	if err != nil {
		return
	}
	
	fmt.Println(resp.Choices[0].Text)
}
```
