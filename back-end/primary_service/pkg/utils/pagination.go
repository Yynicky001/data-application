package utils

// Pagination 分页响应结构
type Pagination struct {
	Limit      int         `json:"limit"`
	Page       int         `json:"page"`
	TotalRows  int64       `json:"totalRows"`
	TotalPages int         `json:"totalPages"`
	Rows       interface{} `json:"rows"`
}
