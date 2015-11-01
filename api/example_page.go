package api

// An ExamplePage occurs when a query cannot be meaningfully computed, but is
// recognized as a topic for which a set of example queries has already been
// prepared.
//
// For example, the query "calculus" would return a Result with an ExamplePage
// linking to http://www.wolframalpha.com/examples/Calculus-content.html.
type ExamplePage struct {
	// The topic name
	Topic string `xml:"category,attr"`

	// The address of the web page with example queries
	URL string `xml:"url,attr"`
}
