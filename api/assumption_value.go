package api

// An AssumptionValue defines a possible value for an assumption.
type AssumptionValue struct {
	// The unique internal identifier for the assumption value
	Name string `xml:"name,attr"`

	// A textual description of the assumption suitable for display to users
	Description string `xml:"desc,attr"`

	// The parameter value needed to invoke this assumption in a subsequent query
	Input string `xml:"input,attr"`
}
