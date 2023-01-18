package openai

import (
	"context"
	"encoding/json"
	"path"

	"github.com/fabiustech/openai/models"
	"github.com/fabiustech/openai/objects"
	"github.com/fabiustech/openai/routes"
)

// FineTuneRequest contains all relevant fields for requests to the fine-tunes endpoints.
type FineTuneRequest struct {
	// TrainingFile specifies the ID of an uploaded file that contains training data. See upload file for how to upload
	// a file.
	//
	// https://beta.openai.com/docs/api-reference/files/upload
	//
	// Your dataset must be formatted as a JSONL file, where each training example is a JSON object with the keys
	// "prompt" and "completion". Additionally, you must upload your file with the purpose fine-tune. See the
	// fine-tuning guide for more details:
	//
	// https://beta.openai.com/docs/guides/fine-tuning/creating-training-data
	TrainingFile string `json:"training_file"`
	// ValidationFile specifies the ID of an uploaded file that contains validation data. If you provide this file, the
	// data is used to generate validation metrics periodically during fine-tuning. These metrics can be viewed in the
	// fine-tuning results file.
	//
	// https://beta.openai.com/docs/guides/fine-tuning/analyzing-your-fine-tuned-model
	//
	// Your train and validation data should be mutually exclusive. Your dataset must be formatted as a JSONL file,
	// where each validation example is a JSON object with the keys "prompt" and "completion". Additionally, you must
	// upload your file with the purpose fine-tune. See the fine-tuning guide for more details:
	//
	// https://beta.openai.com/docs/guides/fine-tuning/creating-training-data
	ValidationFile *string `json:"validation_file,omitempty"`
	// Model specifies the name of the base model to fine-tune. You can select one of "ada", "babbage", "curie",
	// "davinci", or a fine-tuned model created after 2022-04-21. To learn more about these models, see the Models
	// documentation.
	// Defaults to "curie".
	Model *models.FineTune `json:"model,omitempty"`
	// NEpochs specifies the number of epochs to train the model for. An epoch refers to one full cycle through
	// the training dataset.
	// Defaults to 4.
	NEpochs *int `json:"n_epochs,omitempty"`
	// BatchSize specifies the batch size to use for training. The batch size is the number of training examples used
	// to train a single forward and backward pass. By default, the batch size will be dynamically configured to be
	// ~0.2% of the number of examples in the training set, capped at 256 - in general, we've found that larger batch
	// sizes tend to work better for larger datasets.
	// Defaults to null.
	BatchSize *int `json:"batch_size,omitempty"`
	// LearningRateMultiplier specifies the learning rate multiplier to use for training. The fine-tuning learning rate
	// is the original learning rate used for pretraining multiplied by this value. By default, the learning rate
	// multiplier is the 0.05, 0.1, or 0.2 depending on final batch_size (larger learning rates tend to perform better
	// with larger batch sizes). We recommend experimenting with values in the range 0.02 to 0.2 to see what produces
	// the best results.
	// Defaults to null.
	LearningRateMultiplier *int `json:"learning_rate_multiplier,omitempty"`
	// PromptLossWeight specifies the weight to use for loss on the prompt tokens. This controls how much the model
	// tries to learn to generate the prompt (as compared to the completion which always has a weight of 1.0), and can
	// add a stabilizing effect to training when completions are short. If prompts are extremely long (relative to
	// completions), it may make sense to reduce this weight so as to avoid over-prioritizing learning the prompt.
	// Defaults to 0.01.
	PromptLossWeight *int `json:"prompt_loss_weight,omitempty"`
	// ComputeClassificationMetrics calculates classification-specific metrics such as accuracy and F-1 score using the
	// validation set at the end of every epoch if set to true. These metrics can be viewed in the results file.
	//
	// https://beta.openai.com/docs/guides/fine-tuning/analyzing-your-fine-tuned-model
	//
	// In order to compute classification metrics, you must provide a ValidationFile. Additionally, you must specify
	// ClassificationNClasses for multiclass classification or ClassificationPositiveClass for binary classification.
	ComputeClassificationMetrics bool `json:"compute_classification_metrics,omitempty"`
	// ClassificationNClasses specifies the number of classes in a classification task. This parameter is required for
	// multiclass classification.
	// Defaults to null.
	ClassificationNClasses *int `json:"classification_n_classes,omitempty"`
	// ClassificationPositiveClass specifies the positive class in binary classification. This parameter is needed to
	// generate precision, recall, and F1 metrics when doing binary classification.
	// Defaults to null.
	ClassificationPositiveClass *string `json:"classification_positive_class,omitempty"`
	// ClassificationBetas specifies that if provided, we calculate F-beta scores at the specified beta values. The
	// F-beta score is a generalization of F-1 score. This is only used for binary classification. With a beta of 1
	// (i.e. the F-1 score), precision and recall are given the same weight. A larger beta score puts more weight on
	// recall and less on precision. A smaller beta score puts more weight on precision and less on recall.
	// Defaults to null.
	ClassificationBetas []float32 `json:"classification_betas,omitempty"`
	// Suffix specifies a string of up to 40 characters that will be added to your fine-tuned model name. For example,
	// a suffix of "custom-model-name" would produce a model name like
	// ada:ft-your-org:custom-model-name-2022-02-15-04-21-04.
	Suffix string `json:"suffix,omitempty"`
}

// Event represents an event related to a fine-tune request.
type Event struct {
	Object    objects.Object `json:"object"`
	CreatedAt uint64         `json:"created_at"`
	Level     string         `json:"level"`
	Message   string         `json:"message"`
}

// FineTuneResponse is the response from fine-tunes endpoints.
type FineTuneResponse struct {
	ID             string          `json:"id"`
	Object         objects.Object  `json:"object"`
	Model          models.FineTune `json:"model"`
	CreatedAt      uint64          `json:"created_at"`
	Events         []*Event        `json:"events,omitempty"`
	FineTunedModel *string         `json:"fine_tuned_model"`
	Hyperparams    struct {
		BatchSize              int     `json:"batch_size"`
		LearningRateMultiplier float64 `json:"learning_rate_multiplier"`
		NEpochs                int     `json:"n_epochs"`
		PromptLossWeight       float64 `json:"prompt_loss_weight"`
	} `json:"hyperparams"`
	OrganizationID  string   `json:"organization_id"`
	ResultFiles     []string `json:"result_files"`
	Status          string   `json:"status"`
	ValidationFiles []string `json:"validation_files"`
	TrainingFiles   []struct {
		ID        string         `json:"id"`
		Object    objects.Object `json:"object"`
		Bytes     int            `json:"bytes"`
		CreatedAt uint64         `json:"created_at"`
		Filename  string         `json:"filename"`
		Purpose   string         `json:"purpose"`
	} `json:"training_files"`
	UpdatedAt uint64 `json:"updated_at"`
}

// FineTuneDeletionResponse is the response from the fine-tunes/delete endpoint.
type FineTuneDeletionResponse struct {
	ID      string         `json:"id"`
	Object  objects.Object `json:"object"`
	Deleted bool           `json:"deleted"`
}

// CreateFineTune creates a job that fine-tunes a specified model from a given dataset. *FineTuneResponse includes
// details of the enqueued job including job status and the name of the fine-tuned models once complete.
func (c *Client) CreateFineTune(ctx context.Context, ftr *FineTuneRequest) (*FineTuneResponse, error) {
	var b, err = c.post(ctx, routes.FineTunes, ftr)
	if err != nil {
		return nil, err
	}

	var f = &FineTuneResponse{}
	if err = json.Unmarshal(b, f); err != nil {
		return nil, err
	}

	return f, nil
}

// ListFineTunes lists your organization's fine-tuning jobs.
func (c *Client) ListFineTunes(ctx context.Context) (*List[*FineTuneResponse], error) {
	var b, err = c.get(ctx, routes.FineTunes)
	if err != nil {
		return nil, err
	}

	var l = &List[*FineTuneResponse]{}
	if err = json.Unmarshal(b, l); err != nil {
		return nil, err
	}

	return l, nil
}

// RetrieveFineTune gets info about the fine-tune job.
func (c *Client) RetrieveFineTune(ctx context.Context, id string) (*FineTuneResponse, error) {
	var b, err = c.get(ctx, path.Join(routes.FineTunes, id))
	if err != nil {
		return nil, err
	}

	var f = &FineTuneResponse{}
	if err = json.Unmarshal(b, f); err != nil {
		return nil, err
	}

	return f, nil
}

// CancelFineTune immediately cancels a fine-tune job.
func (c *Client) CancelFineTune(ctx context.Context, id string) (*FineTuneResponse, error) {
	var b, err = c.post(ctx, path.Join(routes.FineTunes, id, "cancel"), nil)
	if err != nil {
		return nil, err
	}

	var f = &FineTuneResponse{}
	if err = json.Unmarshal(b, f); err != nil {
		return nil, err
	}

	return f, nil
}

// ListFineTuneEvents returns fine-grained status updates for a fine-tune job.
// TODO: Support streaming (in a different method).
func (c *Client) ListFineTuneEvents(ctx context.Context, id string) (*List[*Event], error) {
	var b, err = c.get(ctx, path.Join(routes.FineTunes, id, "events"))
	if err != nil {
		return nil, err
	}

	var l = &List[*Event]{}
	if err = json.Unmarshal(b, l); err != nil {
		return nil, err
	}

	return l, nil
}

// DeleteFineTune delete a fine-tuned model. You must have the Owner role in your organization.
func (c *Client) DeleteFineTune(ctx context.Context, id string) (*FineTuneDeletionResponse, error) {
	var b, err = c.delete(ctx, path.Join(routes.FineTunes, id))
	if err != nil {
		return nil, err
	}

	var f = &FineTuneDeletionResponse{}
	if err = json.Unmarshal(b, f); err != nil {
		return nil, err
	}

	return f, nil
}
