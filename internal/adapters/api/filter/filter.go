package filter

const (
	Eq               = "eq"
	NotEq            = "notEq"
	GreaterThan      = "gt"
	GreaterThanEqual = "gte"
	LowerThan        = "lt"
	LowerThanEqual   = "lte"
	Between          = "btw"
	Like             = "like"

	Key = "filter"
)

var Operators = []string{"eq", "neq", "gt", "gte", "lt", "lte", "btw", "like"}

type Options []Option

type Option struct {
	Field    string
	Operator string
	Value    string
	Min      string
	Max      string
}
