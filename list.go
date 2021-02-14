package response

func NewList(items interface{}) *List {
	return &List{
		Items: items,
	}
}

type List struct {
	Items interface{} `json:"items"`
}
