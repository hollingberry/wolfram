package api

// A Reinterpretation occurs when Wolfram Alpha cannot understand a query and
// replaces it with a new query that seems close in meaning to the original.
//
// For example, the nonsensical query "blue mustang moon" might be reinterpreted
// as "mustang moon," the name of a 2002 book by Terri Farley.
type Reinterpretation struct {
	// A message that could be displayed to the user before showing the new query.
	// This is almost always "Using closest Wolfram|Alpha interpretation:"
	Message string `xml:"text,attr"`

	// The new query
	Query string `xml:"new,attr"`

	// A value from 0 to 1 indicating how similar the new query is to the original
	// query
	Score float32 `xml:"score,attr"`

	// A description ("low", "medium", or "high") indicating how similar the new
	// query is to the original query
	Level string `xml:"level,attr"`
}
