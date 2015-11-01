package api

// A Source represents a <source> element, an element used by the Wolfram Alpha
// API to provide a link to a web page of source information.  Source
// information is not always present, such as for a purely mathematical
// computation.
type Source struct {
	// The source URL
	URL string `xml:"url,attr"`

	// A short description of the source
	Description string `xml:"text,attr"`
}
