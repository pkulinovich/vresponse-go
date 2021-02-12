package response

func NewPagination(page, limit, total uint64) *Pagination {
	return &Pagination{
		Limit:       limit,
		PageNumber:  page,
		CurrentPage: page,
		Offset:      limit * (page - 1),
		TotalRecord: total,
		TotalPage:   getLastPageNumber(limit, total),
	}
}

type Pagination struct {
	TotalRecord uint64 `json:"totalRecord"`
	TotalPage   uint64 `json:"totalPage"`
	Offset      uint64 `json:"offset"`
	Limit       uint64 `json:"limit"`
	PageNumber  uint64 `json:"pageNumber"`
	CurrentPage uint64 `json:"currentPage"`
}
