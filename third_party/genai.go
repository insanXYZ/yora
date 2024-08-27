package thirdparty

import (
	"context"
	"yora/throw"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GenaiAI struct {
	ctx   context.Context
	model *genai.GenerativeModel
}

func NewGenai(ctx context.Context, apikey string) (*GenaiAI, error) {

	client, err := genai.NewClient(ctx, option.WithAPIKey(apikey))
	if err != nil {
		return nil, throw.ClientGeminiKey()
	}

	model := client.GenerativeModel("gemini-1.5-flash")

	return &GenaiAI{
		ctx:   ctx,
		model: model,
	}, nil
}

func (g *GenaiAI) QuestionStream(text string) *genai.GenerateContentResponseIterator {
	res := g.model.GenerateContentStream(g.ctx, genai.Text(text))
	return res
}
