package api

import (
	"fmt"
	"strings"
)

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

// An AssumptionValue defines a possible value for an assumption.
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

// An Image represents an <img> element, an element returned by the
// Wolfram Alpha API when image results are requested. They point to stored
// image files (usually GIFs, but sometimes JPEGs) giving a formatted visual
// representation of a single subpod.
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

// Mime returns the image MIME type
func (img Image) Mime() string {
	if i := strings.Index(img.URL, "MSPStoreType="); i != -1 {
		return img.URL[i+len("MSPStoreType="):]
	}
	return ""
}

// HTML returns an HTML snippet for displaying the image in a webpage
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

// A LanguageMessage is used to represent a <languagemsg> element, an element
// used by Wolfram Alpha to provide details when it recognizes that your query
// is in a foreign language.
type LanguageMessage struct {
	// The message in English
	English string `xml:"english,attr"`

	// The message in the same language as the query
	Other string `xml:"other,attr"`
}

// A Pod is used to group related results.
//
// For example, the query "amanita" would produce several pods, including ones
// for the mushroom's scientific name, taxonomy, and image, among others.
type Pod struct {
	// The pod title
	Title string `xml:"title,attr"`

	// The name of the scanner that produced the pod
	Scanner string `xml:"scanner,attr"`

	// The pod ID
	ID string `xml:"id,attr"`

	// A number indicating the intended position of the pod in a visual display
	Position int `xml:"position,attr"`

	// Whether a serious processing error occurred with this specific pod
	Error bool `xml:"error,attr"`

	// True if the pod is the query's primary pod
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

// A Result represents a <queryresult> element, the top-level element in queries
// to the Wolfram Alpha API.
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
