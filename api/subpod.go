package api

// A Subpod represents a <subpod> element, an element used by the Wolfram Alpha
// API to hold some kind of information in the results from a query. Subpods
// include various representations of a single datum.
//
// Depending on the query and what type of results are specified, these
// representations might include a textual representation, an image, MathML, or
// Mathematica input/output.
type Subpod struct {
	// The subpod title
	Title string `xml:"title,attr"`

	// The subpod image
	Image Image `xml:"img"`

	// The subpod plaintext representation
	Plaintext string `xml:"plaintext"`

	// The subpod MathML representation
	MathML string `xml:"mathml,innerxml"`

	// The Mathematica input, if available
	MathematicaInput string `xml:"minput"`

	// The Mathematica output, if available
	MathematicaOutput string `xml:"moutput"`

	// Whether the subpod is the query's primary subpod
	Primary bool `xml:"primary,attr"`
}
