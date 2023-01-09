package types

type Field struct {
	Name       string
	Type       string
	Unique     bool
	DefaultVal string
	References string
}

type Table struct {
	Name    string
	Fields  []Field
	Default []string
}

type Schema = map[string][]Table
