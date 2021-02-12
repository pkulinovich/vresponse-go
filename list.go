package response

func NewList(items interface{}) *List {
	return &List{
		Items: items,
	}
}

type List struct {
	HasMore bool        `json:"hasMore"`
	Items   interface{} `json:"items"`
}
