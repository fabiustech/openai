# openai
[![GoDoc](http://img.shields.io/badge/GoDoc-Reference-blue.svg)](https://godoc.org/github.com/fabiustech/openai)
[![Go Report Card](https://goreportcard.com/badge/github.com/sashabaranov/go-gpt3)](https://goreportcard.com/report/github.com/fabiustech/openai)

Zero dependency (Unofficial) Go Client for [OpenAI](https://beta.openai.com/) API endpoints. Built upon the great work done [here](https://github.com/sashabaranov/go-gpt3).

### Goals

Why did we bother to refactor the original library? We have 5 main goals:

1. Use more idiomatic Go style.
2. Better documentation.
3. Make request parameters whose Go default value is a valid parameter value (and differs from OpenAI's defaults) pointers
(See [here](https://github.com/fabiustech/openai/pull/1#:~:text=set%20values.%20(See-,here,-for%20more).) for more).
4. Have a consistent style throughout. 
5. Implement all endpoints.

We hope that by doing the above, future maintenance should also be trivial.
Read more on the original refactoring PR [here](https://github.com/fabiustech/openai/pull/1).

### Installation
```
go get github.com/fabiustech/openai
```

### Example Usage

```go
package main

import (
	"context"
	"fmt"
	"os"
	
	"github.com/fabiustech/openai"
	"github.com/fabiustech/openai/models"
	"github.com/fabiustech/openai/params"
)

func main() {
	var key, ok = os.LookupEnv("OPENAI_API_KEY")
	if !ok {
		panic("env variable OPENAI_API_KEY not set")
    }
	var c = openai.NewClient(key)

	var resp, err = c.CreateCompletion(context.Background(), &openai.CompletionRequest[models.Completion]{
		Model:       models.TextDavinci003,
		MaxTokens:   100,
		Prompt:      "Lorem ipsum",
		Temperature: params.Optional(0.0),
	})
	if err != nil {
		return
	}

	fmt.Println(resp.Choices[0].Text)
}
```

### Contributing

Contributions are welcome and encouraged! Feel free to report any bugs / feature requests as issues.
