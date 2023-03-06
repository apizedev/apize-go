package apize

import "context"

type TextClient interface {
	// Measure counts the number of words in a given text input. Because several of the Text Intelligence APIs are priced
	// in terms of word count, this API is useful to see exactly how large Apize considers an input to be.
	Measure(ctx context.Context, req *MeasureRequest) (*MeasureResponse, error)

	// Summarize generates an AI summary of a given text input.
	Summarize(ctx context.Context, req *SummarizeRequest) (*SummarizeResponse, error)

	// Sentiment runs sentiment analysis on a given text input and rates the sentiment on a scale from "very_negative" to
	// "very_positive".
	Sentiment(ctx context.Context, req *SentimentRequest) (*SentimentResponse, error)

	// ContentMatching determines how closely a given text input is relates to a list of content categories.
	ContentMatching(ctx context.Context, req *ContentMatchingRequest) (*ContentMatchingResponse, error)

	// Grammar corrects grammar in a given text input, and optionally explains the grammatical errors.
	Grammar(ctx context.Context, req *GrammarRequest) (*GrammarResponse, error)
}

type MeasureRequest struct {
	Text string `json:"text"`
}

type MeasureResponse struct {
	Words int `json:"words"`
}

type SummarizeRequest struct {
	Text      string `json:"text"`
	Sentences *int   `json:"sentences"`
}

type SummarizeResponse struct {
	Summary string `json:"summary"`
}

type SentimentRequest struct {
	Text string `json:"text"`
}

type SentimentResponse struct {
	Sentiment string `json:"sentiment"`
}

type ContentMatchingRequest struct {
	Text       string   `json:"text"`
	Categories []string `json:"categories"`
}

type ContentMatchingResponse struct {
	Categories map[string]float32 `json:"categories"`
}

type GrammarRequest struct {
	Text    string `json:"text"`
	Explain *bool  `json:"explain"`
}

type GrammarResponse struct {
	Text        string  `json:"text"`
	Explanation *string `json:"explanation"`
}

func (c *client) Measure(ctx context.Context, req *MeasureRequest) (*MeasureResponse, error) {
	return httpPost[MeasureResponse](ctx, c, "/solutions/text/v1/measure", req)
}

func (c *client) Summarize(ctx context.Context, req *SummarizeRequest) (*SummarizeResponse, error) {
	return httpPost[SummarizeResponse](ctx, c, "/solutions/text/v1/summarize", req)
}

func (c *client) Sentiment(ctx context.Context, req *SentimentRequest) (*SentimentResponse, error) {
	return httpPost[SentimentResponse](ctx, c, "/solutions/text/v1/sentiment", req)
}

func (c *client) ContentMatching(ctx context.Context, req *ContentMatchingRequest) (*ContentMatchingResponse, error) {
	return httpPost[ContentMatchingResponse](ctx, c, "/solutions/text/v1/content-matching", req)
}

func (c *client) Grammar(ctx context.Context, req *GrammarRequest) (*GrammarResponse, error) {
	return httpPost[GrammarResponse](ctx, c, "/solutions/text/v1/grammar", req)
}
