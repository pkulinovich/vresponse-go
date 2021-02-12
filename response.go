package response

import (
	"net/http"
)

func New() *Response {
	return new(Response)
}

type Response struct {
	Data       interface{} `json:"data,omitempty"`
	Links      interface{} `json:"links,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Errors     []*Error    `json:"errors,omitempty"`
}

func (r *Response) WithList(items interface{}) *Response {
	list := NewList(items)
	return r.SetData(list)
}

func (r *Response) WithError(status int, title, code, details string) *Response {
	return r.AddError(status, title, details, code)
}

func (r *Response) WithErrorStatus(status int) *Response {
	text := http.StatusText(status)
	return r.AddError(status, text, text, "")
}

func (r *Response) Paginate(page, limit, total uint64) *Response {
	r.Pagination = NewPagination(page, limit, total)
	return r
}

func (r *Response) PaginateLinks(path string, page, limit, total uint64) *Response {
	r.Links = NewLinks(path, page, limit, total)
	return r
}

func (r *Response) AddError(status int, title, details, code string) *Response {
	e := &Error{
		Title:   title,
		Details: details,
		Status:  status,
		Code:    code,
	}
	r.Errors = append(r.Errors, e)
	return r
}

func (r *Response) SetErrors(errors []*Error) *Response {
	r.Errors = errors
	return r
}

func (r *Response) SetData(data interface{}) *Response {
	r.Data = data
	return r
}
