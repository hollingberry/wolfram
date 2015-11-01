package api

// A Result represents a <queryresult> element, the top-level element in queries
// to the Wolfram Alpha API.
type Result struct {
	// A comma-separated list of the categories and types of data represented in
	// the results
	Datatypes string `xml:"datatypes,attr"`

	// True or false depending on whether a serious processing error occurred,
	// such as a missing required parameter. If true there will be no pod
	// content, just an error.
	Error bool `xml:"error,attr"`

	// The query ID
	ID string `xml:"id,attr"`

	ParseTimedOut    bool             `xml:"parsetimedout,attr"`
	ParseTiming      float32          `xml:"parsetiming,attr"`
	Recalculate      string           `xml:"recalculate,attr"`
	Success          bool             `xml:"success,attr"`
	TimedOut         string           `xml:"timedout,attr"` // arraylike
	Timing           float32          `xml:"timing,attr"`
	Version          string           `xml:"version,attr"`
	Error            Error            `xml:"error"`
	ExamplePage      ExamplePage      `xml:"examplepage"`
	LanguageMessage  LanguageMessage  `xml:"languagemsg"`
	Reinterpretation Reinterpretation `xml:"reinterpret"`
	Assumptions      []Assumption     `xml:"assumption"`
	Pods             []Pod            `xml:"pod"`
	Sources          []Source         `xml:"source"`
	Suggestions      []string         `xml:"didyoumean"`
}

// func (res Result) FutureTopic() {}
//
// func (res Result) PrimaryText() {}
//
// func (res Result) Reinterpreted() {}
//
// func (res Result) Tips() {}
