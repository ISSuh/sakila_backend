package common

type Pagnation struct {
	Page   int    `form:"page" json:"page"`
	Offset int    `form:"offset" json:"offset"`
	Limit  int    `form:"limit" json:"limit"`
	Sort   string `form:"sort" json:"sort,omitempty"`
	SortBy string `form:"sortBy" json:"sort_by,omitempty"`
	Total  int64  `json:"total"`
	Item   any    `json:"item,omitempty"`
}

func (p Pagnation) IsValidate() bool {
	return (p.Page > 0) && (p.Limit > 0)
}

func (p Pagnation) IsOverThanTotal(total int) bool {
	return p.CalculateOffset() > total
}

func (p Pagnation) CalculateOffset() int {
	if p.Offset > 0 {
		return p.Offset
	}
	return ((p.Page - 1) * p.Limit)
}

func (p Pagnation) SortSQL() string {
	if p.Sort == "d" {
		return p.SortBy + " desc"
	}
	return p.SortBy + " asc"
}
