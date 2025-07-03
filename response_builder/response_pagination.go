package responsebuilder

type ResponsePagination struct {
	TotalPage int `json:"total_page,omitempty"`
	TotalRows int `json:"total_rows,omitempty"`
	Page      int `json:"page,omitempty"`
	Size      int `json:"size,omitempty"`
}

func NewResponsePagination(totalRows int, page int, size int) *ResponsePagination {
	if size <= 0 {
		size = 1
	}

	return &ResponsePagination{
		TotalRows: totalRows,
		TotalPage: (totalRows + size - 1) / size,
		Page:      page,
		Size:      size,
	}
}
