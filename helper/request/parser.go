package request

import "net/http"

const (
	defaultTagName = "json"
)

type parser struct {
	tagName string
}

type Parser interface {
	Form(r *http.Request, i interface{}) error
	Query(r *http.Request, i interface{}) error
}

func NewDefaultParser() Parser {
	return &parser{
		tagName: defaultTagName,
	}
}

func (c *parser) Form(r *http.Request, i interface{}) error {
	return nil
}

func (c *parser) Query(r *http.Request, i interface{}) error {
	return nil
}
