package filter

const (
	Tag = "json"

	Eq               = "eq"
	NotEq            = "notEq"
	GreaterThan      = "gt"
	GreaterThanEqual = "gte"
	LowerThan        = "lt"
	LowerThanEqual   = "lte"
	Between          = "btw"
	Like             = "like"

	Limit  = "count"
	Offset = "start"
	Sort   = "sortBy"
	Desc   = "desc"

	Key = "filter"
)

var Operators = []string{"eq", "neq", "gt", "gte", "lt", "lte", "btw", "like"}

var GeneralFilters = []string{Limit, Offset, Sort, Desc}

type Options []Option

type Option struct {
	Field    string
	Operator string
	Value    string
	Min      string
	Max      string
}
