package domain

type Query struct {
	Bool Bool `json:"bool,omitempty"`
}

type Bool struct {
	Must   []interface{} `json:"must,omitempty"`
	Filter []interface{} `json:"filter,omitempty"`
	Should []interface{} `json:"should,omitempty"`
}

type CombinedFields struct {
	CombinedFields CombinedFieldsValue `json:"combined_fields"`
}

type CombinedFieldsValue struct {
	Query  string   `json:"query,omitempty"`
	Fields []string `json:"fields,omitempty"`
}

type Term struct {
	Term map[string]TermValue `json:"term,omitempty"`
}

type TermValue struct {
	Value interface{} `json:"value,omitempty"`
	Boost float64     `json:"boost,omitempty"`
}
