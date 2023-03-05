package apize

import "fmt"

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	DocsUrl string `json:"docs_url"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}
