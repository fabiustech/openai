package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/fabiustech/openai/images"
	"github.com/fabiustech/openai/models"
	"github.com/fabiustech/openai/objects"
	"github.com/fabiustech/openai/params"
)

/*
This test suite has been ported from the original repo: https://github.com/sashabaranov/go-gpt3.
It is incomplete, and it's usefulness is questionable.

TODO: Cover all endpoints.
*/

const (
	testToken = "this-is-my-secure-token-do-not-steal!!"
)

func TestAPI(t *testing.T) {
	var token, ok = os.LookupEnv("OPENAI_TOKEN")
	if !ok {
		t.Skip("Skipping testing against production OpenAI API. Set OPENAI_TOKEN environment variable to enable it.")
	}

	var c = NewClient(token)
	var ctx = context.Background()
	var _, err = c.ListEngines(ctx)
	if err != nil {
		t.Fatalf("ListEngines error: %v", err)
	}

	_, err = c.GetEngine(ctx, "davinci")
	if err != nil {
		t.Fatalf("GetEngine error: %v", err)
	}

	var fl *List[*File]
	fl, err = c.ListFiles(ctx)
	if err != nil {
		t.Fatalf("ListFiles error: %v", err)
	}

	if len(fl.Data) > 0 {
		_, err = c.RetrieveFile(ctx, fl.Data[0].ID)
		if err != nil {
			t.Fatalf("RetrieveFile error: %v", err)
		}
	}

	_, err = c.CreateEmbeddings(ctx, &EmbeddingRequest{
		Input: []string{
			"The food was delicious and the waiter",
			"Other examples of embedding request",
		},
		Model: models.AdaEmbeddingV2,
	})
	if err != nil {
		t.Fatalf("Embedding error: %v", err)
	}
}

func newTestClient(u string) (*Client, error) {
	var h, err = url.Parse(u)
	if err != nil {
		return nil, err
	}

	return &Client{
		token:  testToken,
		host:   h.Host,
		scheme: "http",
	}, nil
}

// TestCompletions Tests the completions endpoint of the API using the mocked server.
func TestCompletions(t *testing.T) {
	var ts = OpenAITestServer()
	ts.Start()
	defer ts.Close()

	var client, _ = newTestClient(ts.URL)

	var _, err = client.CreateCompletion(context.Background(), &CompletionRequest[models.Completion]{
		Prompt:    "Lorem ipsum",
		Model:     models.TextDavinci003,
		MaxTokens: 5,
	})
	if err != nil {
		t.Fatalf("CreateCompletion error: %v", err)
	}
}

// TestEdits Tests the edits endpoint of the API using the mocked server.
func TestEdits(t *testing.T) {
	var ts = OpenAITestServer()
	ts.Start()
	defer ts.Close()

	var client, _ = newTestClient(ts.URL)

	var n = 3
	var resp, err = client.CreateEdit(context.Background(), &EditsRequest{
		Model: models.TextDavinciEdit001,
		Input: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, " +
			"sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim" +
			" ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip" +
			" ex ea commodo consequat. Duis aute irure dolor in reprehe",
		Instruction: "test instruction",
		N:           n,
	})
	if err != nil {
		t.Fatalf("Edits error: %v", err)
	}
	if len(resp.Choices) != n {
		t.Fatalf("edits does not properly return the correct number of choices")
	}
}

func TestEmbedding(t *testing.T) {
	embeddedModels := []models.Embedding{
		models.AdaEmbeddingV2,
	}
	for _, model := range embeddedModels {
		embeddingReq := &EmbeddingRequest{
			Input: []string{
				"The food was delicious and the waiter",
				"Other examples of embedding request",
			},
			Model: model,
		}
		// marshal embeddingReq to JSON and confirm that the model field equals
		// the AdaSearchQuery type
		marshaled, err := json.Marshal(embeddingReq)
		if err != nil {
			t.Fatalf("Could not marshal embedding request: %v", err)
		}
		if !bytes.Contains(marshaled, []byte(`"model":"`+model.String()+`"`)) {
			t.Fatalf("Expected embedding request to contain model field")
		}
	}
}

// getEditBody Returns the body of the request to create an edit.
func getEditBody(r *http.Request) (*EditsRequest, error) {
	edit := &EditsRequest{}
	// read the request body
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(reqBody, &edit)
	if err != nil {
		return nil, err
	}
	return edit, nil
}

// handleEditEndpoint Handles the edit endpoint by the test server.
func handleEditEndpoint(w http.ResponseWriter, r *http.Request) {
	var err error
	var resBytes []byte

	// edits only accepts POST requests
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	var editReq *EditsRequest
	editReq, err = getEditBody(r)
	if err != nil {
		http.Error(w, "could not read request", http.StatusInternalServerError)
		return
	}
	// create a response
	res := &EditsResponse{
		Object:  objects.Edit,
		Created: uint64(time.Now().Unix()),
	}
	// edit and calculate token usage
	editString := "edited by mocked OpenAI server :)"
	inputTokens := numTokens(editReq.Input+editReq.Instruction) * editReq.N
	completionTokens := int(float32(len(editString))/4) * editReq.N
	for i := 0; i < editReq.N; i++ {
		// instruction will be hidden and only seen by OpenAI
		res.Choices = append(res.Choices, &EditsChoice{
			Text:  editReq.Input + editString,
			Index: i,
		})
	}
	res.Usage = &Usage{
		PromptTokens:     inputTokens,
		CompletionTokens: completionTokens,
		TotalTokens:      inputTokens + completionTokens,
	}
	resBytes, _ = json.Marshal(res)
	fmt.Fprint(w, string(resBytes))
}

// handleCompletionEndpoint Handles the completion endpoint by the test server.
func handleCompletionEndpoint(w http.ResponseWriter, r *http.Request) {
	var err error
	var resBytes []byte

	// completions only accepts POST requests
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	var completionReq *CompletionRequest[models.Completion]
	if completionReq, err = getCompletionBody(r); err != nil {
		http.Error(w, "could not read request", http.StatusInternalServerError)
		return
	}
	res := &CompletionResponse[models.Completion]{
		ID:      strconv.Itoa(int(time.Now().Unix())),
		Object:  objects.TextCompletion,
		Created: uint64(time.Now().Unix()),
		// would be nice to validate Model during testing, but
		// this may not be possible with how much upkeep
		// would be required / wouldn't make much sense
		Model: completionReq.Model,
	}
	// create completions
	for i := 0; i < completionReq.N; i++ {
		// generate a random string of length completionReq.Length
		completionStr := strings.Repeat("a", completionReq.MaxTokens)
		if completionReq.Echo {
			completionStr = completionReq.Prompt + completionStr
		}
		res.Choices = append(res.Choices, &CompletionChoice{
			Text:  completionStr,
			Index: i,
		})
	}
	inputTokens := numTokens(completionReq.Prompt) * completionReq.N
	completionTokens := completionReq.MaxTokens * completionReq.N
	res.Usage = &Usage{
		PromptTokens:     inputTokens,
		CompletionTokens: completionTokens,
		TotalTokens:      inputTokens + completionTokens,
	}
	resBytes, _ = json.Marshal(res)
	fmt.Fprintln(w, string(resBytes))
}

// handleImageEndpoint Handles the images endpoint by the test server.
func handleImageEndpoint(w http.ResponseWriter, r *http.Request) {
	// Images only accepts POST requests.
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	var ir, err = getImageBody(r)
	if err != nil {
		http.Error(w, "could not read request", http.StatusInternalServerError)
		return
	}

	var resp = &ImageResponse{
		Created: uint64(time.Now().Unix()),
	}

	// Handle default values.
	if ir.N == 0 {
		ir.N = 1
	}

	for i := 0; i < ir.N; i++ {
		var imageData = &ImageData{}
		switch ir.ResponseFormat {
		// Invalid is the go default value, and URL is the default API behavior.
		case images.FormatURL, images.FormatInvalid:
			imageData.URL = params.Optional("https://example.com/image.png")
		case images.FormatB64JSON:
			// This decodes to "{}" in base64.
			imageData.B64JSON = params.Optional("e30K")
		default:
			http.Error(w, "invalid response format", http.StatusBadRequest)
			return
		}
		resp.Data = append(resp.Data, imageData)
	}

	var b, _ = json.Marshal(resp)
	_, _ = w.Write(b)
}

// getCompletionBody Returns the body of the request to create a completion.
func getCompletionBody(r *http.Request) (*CompletionRequest[models.Completion], error) {
	var completion = &CompletionRequest[models.Completion]{}
	// read the request body
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(reqBody, &completion)
	if err != nil {
		return nil, err
	}
	return completion, nil
}

// getImageBody Returns the body of the request to create a image.
func getImageBody(r *http.Request) (*CreateImageRequest, error) {
	var image = &CreateImageRequest{}
	// read the request body
	var reqBody, err = io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(reqBody, &image)
	if err != nil {
		return nil, err
	}

	return image, nil
}

// numTokens Returns the number of GPT-3 encoded tokens in the given text.
// This function approximates based on the rule of thumb stated by OpenAI:
// https://beta.com/tokenizer
//
// TODO: implement an actual tokenizer for GPT-3 and Codex (once available)
func numTokens(s string) int {
	return int(float32(len(s)) / 4)
}

func TestImages(t *testing.T) {
	var ts = OpenAITestServer()
	ts.Start()
	defer ts.Close()

	var client, _ = newTestClient(ts.URL)
	var _, err = client.CreateImage(context.Background(), &CreateImageRequest{Prompt: "Lorem ipsum"})
	if err != nil {
		t.Fatalf("CreateImage error: %v", err)
	}
}

// OpenAITestServer Creates a mocked OpenAI server which can pretend to handle requests during testing.
func OpenAITestServer() *httptest.Server {
	return httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("received request at path %q\n", r.URL.Path)

		// check auth
		if r.Header.Get("Authorization") != "Bearer "+testToken {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// OPTIMIZE: create separate handler functions for these
		switch r.URL.Path {
		case "/v1/edits":
			handleEditEndpoint(w, r)
			return
		case "/v1/completions":
			handleCompletionEndpoint(w, r)
			return
		case "/v1/images/generations":
			handleImageEndpoint(w, r)
		// TODO: Implement the other endpoints.
		default:
			// the endpoint doesn't exist
			http.Error(w, "the resource path doesn't exist", http.StatusNotFound)
			return
		}
	}))
}
