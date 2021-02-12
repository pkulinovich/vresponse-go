package response

func NewLinks(path string, page, limit, total uint64) *Links {
	return &Links{
		First: buildFirstPath(path, limit),
		Last:  buildLastPath(path, limit, total),
		Next:  buildNextPath(path, page, limit, total),
		Prev:  buildPrevPath(path, page, limit, total),
		Self:  buildSelfPath(path, page, limit),
	}
}

type Links struct {
	Self  *string `json:"self"`
	Next  *string `json:"next"`
	Prev  *string `json:"prev"`
	First *string `json:"first"`
	Last  *string `json:"last"`
}
