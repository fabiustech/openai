package openai

import (
	"context"
	"os"

	"github.com/fabiustech/openai/audio"
	"github.com/fabiustech/openai/models"
)

// AudioTranscriptionRequest is the request body for the audio/transcriptions endpoint.
type AudioTranscriptionRequest struct {
	// File is the audio file object (not file name) to transcribe, in one of these formats:
	// mp3, mp4, mpeg, mpga, m4a, wav, or webm.
	File *os.File
	// Model is the ID of the model to use. Only whisper-1 is currently available.
	Model models.Audio
	// Prompt is optional text to guide the model's style or continue a previous audio segment. The prompt should match
	// the audio language.
	Prompt *string
	// ResponseFormat is the format of the transcript output, in one of these options:
	// json, text, srt, verbose_json, or vtt.
	ResponseFormat *audio.Format
	// Temperature is he sampling temperature, between 0 and 1. Higher values like 0.8 will make the output more random,
	// while lower values like 0.2 will make it more focused and deterministic. If set to 0, the model will use log
	// probability to automatically increase the temperature until certain thresholds are hit.
	Temperature *float64
	// Language is the language of the input audio. Supplying the input language in ISO-639-1 format will improve
	// accuracy and latency.
	Language *string
}

// TranscribeAudioFile creates a new audio file transcription request. File uploads are currently limited to 25 MB
// and the following input file types are supported:mp3, mp4, mpeg, mpga, m4a, wav, and webm.
// The returned []byte is the raw response from the API (as the response format changes depending on the contents of
// the request).
func (c *Client) TranscribeAudioFile(ctx context.Context, ar *AudioTranscriptionRequest) ([]byte, error) {
	return c.postAudio(ctx, ar)
}
