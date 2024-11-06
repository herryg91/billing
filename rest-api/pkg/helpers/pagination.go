package helpers

import "math"

type Pagination struct {
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
	From        int `json:"from"`
	To          int `json:"to"`
	FirstPage   int `json:"first_page"`
	PerPage     int `json:"per_page"`
	Total       int `json:"total"`
}

func (Pagination) New(page, size, total_data int) Pagination {
	return Pagination{
		CurrentPage: page,
		LastPage:    int(math.Ceil(float64(total_data) / float64(size))),
		From:        (page-1)*size + 1,
		To:          page * size,
		FirstPage:   1,
		PerPage:     size,
		Total:       total_data,
	}
}

func (p Pagination) NextPage() int {
	next_page := p.CurrentPage + 1
	if next_page > p.LastPage {
		next_page = p.LastPage
	}
	return next_page
}
func (p Pagination) PrevPage() int {
	prev_page := p.CurrentPage - 1
	if prev_page < 1 {
		prev_page = 1
	}
	return prev_page
}
