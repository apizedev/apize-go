package apize

import "context"

type TextClient interface {
	// Measure counts the number of words in a given text input. Because several of the Text Intelligence APIs are priced
	// in terms of word count, this API is useful to see exactly how large Apize considers an input to be.
	Measure(ctx context.Context, req *MeasureRequest) (*MeasureResponse, error)

	// Summarize generates an AI summary of a given text input.
	Summarize(ctx context.Context, req *SummarizeRequest) (*SummarizeResponse, error)
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

func (c *client) Measure(ctx context.Context, req *MeasureRequest) (*MeasureResponse, error) {
	return httpPost[MeasureResponse](ctx, c, "/solutions/text/v1/measure", req)
}

func (c *client) Summarize(ctx context.Context, req *SummarizeRequest) (*SummarizeResponse, error) {
	return httpPost[SummarizeResponse](ctx, c, "/solutions/text/v1/summarize", req)
}
