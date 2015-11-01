package api

// An Assumption defines a single assumption, typically about the meaning of a
// word or phrase in the query, and a series of possible other values.
type Assumption struct {
	// The assumption type
	Type string `xml:"type,attr"`

	// The word or phrase to which the assumption is applied
	Word string `xml:"word,attr"`

	// The possible assumption values (the first is the assumed value)
	Values []AssumptionValue `xml:"value"`
}
