package utils

import "math"

// DefaultSize :nodoc:
const DefaultSize int = 10

// DefaultPage :nodoc:
const DefaultPage int = 1

// ResponseWithPagination :nodoc:
type ResponseWithPagination struct {
	Records     interface{}            `json:"records"`
	PageSummary map[string]interface{} `json:"pageSummary"`
}

// BuildPagination creates response with pagination
func BuildPagination(result interface{}, total, page, size int) ResponseWithPagination {
	if page == 0 {
		page = DefaultPage
	}

	if size == 0 {
		size = DefaultSize
	}

	totalPages := math.Ceil(float64(total / size))

	var hasNext bool
	if page < int(totalPages) {
		hasNext = true
	}

	if !hasNext {
		page = 0
	} else {
		page = page + 1
	}

	return ResponseWithPagination{
		Records: result,
		PageSummary: map[string]interface{}{
			"size":    size,
			"page":    page,
			"hasNext": hasNext,
			"total":   total,
		},
	}
}
