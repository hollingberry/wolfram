package api

import (
	"fmt"
	"net/url"
)

// An Assumption defines a single assumption, typically about the meaning of a
// word or phrase in the query, and a series of possible other values.
//
// For instance, Wolfram Alpha will assume that the query "pi" refers to the
// mathematical constant, as opposed to the Greek character, the movie, etc. The
// Result for this query will have an Assumption with this information.
//
// TODO: Add template field (see assumptions.xml)
type Assumption struct {
	// The assumption type
	Type string `xml:"type,attr"`

	// The word or phrase to which the assumption is applied
	Word string `xml:"word,attr"`

	// The possible assumption values (the first is the assumed value)
	Values []AssumptionValue `xml:"value"`
}

// An AssumptionValue defines a possible value for an assumption.
//
// In the Assumption example above, there would be an AssumptionValue for the
// mathematical constant pi, for the Greek character pi, for the movie Pi, etc.
type AssumptionValue struct {
	// The unique internal identifier for the assumption value
	Name string `xml:"name,attr"`

	// A textual description of the assumption suitable for display to users
	Description string `xml:"desc,attr"`

	// The parameter value needed to invoke this assumption in a subsequent query
	Input string `xml:"input,attr"`
}

// An Error occurs when something goes wrong with the request.
//
// This might occur when the query as a whole fails (e.g., the input parameters
// are invalid, the App ID is incorrect, or Wolfram Alpha experiences an
// internal error). In such cases, the Result will have the Error.
//
// Wolfram Alpha might also fail to process an individual pod, even if the query
// as a whole succeeds. When this happens, the Pod that failed will have the
// Error.
type Error struct {
	// The error code
	Code int `xml:"code"`

	// A short message describing the error
	Message string `xml:"msg"`
}

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

// An Image occurs within a Subpod when image results are requested. They point
// to stored image files (usually GIFs, but sometimes JPEGs) giving a formatted
// visual representation of a single subpod.
//
// If requested, almost all subpods will include an Image representation—even
// textual subpods. That is, the Image in textual subpods will just point to a
// picture of text.
type Image struct {
	// The image URL
	URL string `xml:"src,attr"`

	// The image alt text
	Alt string `xml:"alt,attr"`

	// The image title
	Title string `xml:"title,attr"`

	// The image width, in pixels
	Width int `xml:"width,attr"`

	// The image height, in pixels
	Height int `xml:"height,attr"`
}

// HTML returns an HTML string for displaying the image in a webpage.
func (img Image) HTML() string {
	return fmt.Sprintf(
		`<img src="%s" alt="%s" title="%s" width="%d" height="%d"/>`,
		img.URL,
		img.Alt,
		img.Title,
		img.Width,
		img.Height,
	)
}

// Mime returns the image MIME type, or an empty string if the MIME type cannot
// be guessed.
func (img Image) Mime() string {
	u, err := url.Parse(img.URL)
	if err != nil {
		return ""
	}
	return u.Query().Get("MSPStoreType")
}

// A LanguageMessage occurs when a query is in a foreign language.
//
// For instance, the Result for the query "wo noch nie" will contain a
// LanguageMessage explaining that Wolfram Alpha does not yet support German.
type LanguageMessage struct {
	// The message in English
	English string `xml:"english,attr"`

	// The message in the same language as the query
	Other string `xml:"other,attr"`
}

// A Pod corresponds roughly to one category of result. The Result for all
// queries will contain at least two Pods (the "Input interpretation" and
// "Result"); many queries return more.
//
// For example, the query "amanita" produces seven pods, which have titles like
// "Scientific name", "Taxonomy", and "Image", among others.
//
// TODO: Make the field comments prettier
type Pod struct {
	// The pod title
	Title string `xml:"title,attr"`

	// The name of the scanner that produced the pod
	Scanner string `xml:"scanner,attr"`

	// The unique internal identifier for the pod type
	ID string `xml:"id,attr"`

	// An integer (often a multiple of 100) indicating the intended position of
	// the pod in a visual display. Pods with a smaller position should be
	// displayed above pods with a greater position.
	Position int `xml:"position,attr"`

	// Whether a serious processing error occurred with this specific pod
	Error bool `xml:"error,attr"`

	// Whether the pod is the query's primary pod
	Primary bool `xml:"primary,attr"`

	// The pod subpods
	Subpods []Subpod `xml:"subpod"`
}

// A Reinterpretation occurs when Wolfram Alpha cannot understand a query and
// replaces it with a new query that seems close in meaning to the original.
//
// For example, the nonsensical query "blue mustang moon" might be reinterpreted
// as "mustang moon," the name of a 2002 book by Terri Farley.
type Reinterpretation struct {
	// A message that could be displayed to the user before showing the new query
	// (e.g., "Using closest Wolfram|Alpha interpretation:")
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

// A Result represents a <queryresult> element, the top-level element in queries
// to the Wolfram Alpha API.
//
// TODO: Add field comments
type Result struct {
	// A comma-separated list of the categories and types of data represented in
	// the results
	Datatypes string `xml:"datatypes,attr"`

	// True or false depending on whether a serious processing error occurred,
	// such as a missing required parameter. If true there will be no pod
	// content, just an error.
	// TODO: Rename
	ErrorSTATUS bool `xml:"error,attr"`

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

// A Source provides a link to a web page with source information. Sources are
// found in Results that do not exclusively use "common knowledge" or purely
// mathematical computation.
//
// Note that the URL links not directly to a source, but rather to a Wolfram
// Alpha webpage with source information for a general topic. See
// http://www.wolframalpha.com/sources/GivenNameDataSourceInformationNotes.html
// as an example.
type Source struct {
	// The address of the web page with source information
	URL string `xml:"url,attr"`

	// A short description of the source
	Description string `xml:"text,attr"`
}

// A Subpod represents a <subpod> element, an element used by the Wolfram Alpha
// API to hold some kind of information in the results from a query. Subpods
// include various representations of a single datum.
//
// Depending on the query and what type of results are specified, these
// representations might include a textual representation, an image, MathML, or
// Mathematica input/output.
type Subpod struct {
	// The subpod title, usually an empty string
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
