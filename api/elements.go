package api

import (
	"encoding/xml"
	"net/url"
	"strings"
)

// An Assumption defines a single assumption, typically about the meaning of a
// word or phrase in the query, and a series of possible other values.
//
// For instance, Wolfram Alpha will assume that the query "pi" refers to the
// mathematical constant, as opposed to the Greek character, the movie, etc.
// The Result for the query "pi", then, will have an Assumption with this
// information.
type Assumption struct {
	// The tag name
	XMLName struct{} `xml:"assumption"`

	// The assumption type
	Type string `xml:"type,attr"`

	// The word or phrase to which the assumption is applied
	Word string `xml:"word,attr"`

	// The possible assumption values (the first is the assumed value)
	Values []AssumptionValue `xml:"value"`

	// A template for generating a message to display to the user
	Template string `xml:"template,attr"`
}

// An AssumptionValue defines a possible value for an assumption.
//
// In the Assumption example above, there would be an AssumptionValue for the
// mathematical constant 3.14159..., for the Greek character π, for the movie Pi, etc.
type AssumptionValue struct {
	// The tag name
	XMLName struct{} `xml:"value"`

	// The internal identifier for the assumption value
	Name string `xml:"name,attr"`

	// A description of the assumption suitable for display to the user
	Description string `xml:"desc,attr"`

	// The query value needed to invoke this assumption in a subsequent query
	Input string `xml:"input,attr"`
}

// An Error occurs when something goes wrong with the request.
//
// This might happen when the query as a whole fails (e.g., the input parameters
// are invalid, the App ID is incorrect, or Wolfram Alpha experiences an
// internal error). In such cases, the Result will have the Error.
//
// Wolfram Alpha might also fail to process an individual pod, even if the query
// as a whole succeeds. When this happens, the Pod that failed will have the
// Error.
type Error struct {
	// The tag name
	XMLName struct{} `xml:"error"`

	// The error code
	Code int `xml:"code"`

	// A short message describing the error
	Message string `xml:"msg"`
}

// An ExamplePage occurs when a query cannot be meaningfully computed, but is
// recognized as a topic for which a set of example queries has already been
// prepared.
//
// For example, the Result for the query "calculus" would include an ExamplePage
// linking to http://www.wolframalpha.com/examples/Calculus-content.html.
type ExamplePage struct {
	// The tag name
	XMLName struct{} `xml:"examplepage"`

	// The topic name
	Topic string `xml:"category,attr"`

	// The address of the web page with example queries
	URL string `xml:"url,attr"`
}

// A FutureTopic occurs when a query cannot be meaningfully computed, but is
// recognized as a topic under development.
//
// For example, the Result for the query "microsoft windows" would include a
// FutureTopic indicating that the topic "Operating Systems" is under
// investigation.
type FutureTopic struct {
	// The tag name
	XMLName struct{} `xml:"futuretopic"`

	// The topic name
	Topic string `xml:"topic,attr"`

	// A short message explaining why there is no data for the topic
	// (usually "Development of this topic is under investigation...")
	Message string `xml:"msg,attr"`
}

// An Image occurs within a Subpod when image results are requested. They point
// to stored image files (usually GIFs, but sometimes JPEGs) giving a formatted
// visual representation of a single subpod.
//
// If requested, almost all subpods will include an Image representation—even
// textual subpods. (The image in textual subpods will just point to a picture
// of text.)
type Image struct {
	// The tag name
	XMLName struct{} `xml:"img"`

	// The image URL
	URL string `xml:"src,attr"`

	// The image alt text
	Alt string `xml:"alt,attr"`

	// The image title
	Title string `xml:"title,attr"`

	// The image width, in pixels
	Width int `xml:"width,attr,omitempty"`

	// The image height, in pixels
	Height int `xml:"height,attr,omitempty"`
}

// HTML returns an HTML string for displaying the image in a webpage.
func (img Image) HTML() string {
	x, err := xml.Marshal(&img)
	if err != nil {
		panic(err)
	}
	return strings.Replace(string(x), "></img>", "/>", 1)
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
	// The tag name
	XMLName struct{} `xml:"languagemsg"`

	// The message in English
	English string `xml:"english,attr"`

	// The message in the same language as the query
	Other string `xml:"other,attr"`
}

// MathML occurs within a Subpod when MathML results are requested. MathML is a
// low-level specification for mathematical and scientific content on the Web
// and beyond. See http://www.w3.org/Math/ for the specification.
//
// Though it only has one field, this type is needed for unmarshaling the inner
// XML of MathML elements due to limitations of the encoding/xml package. See
// http://grokbase.com/t/gg/golang-nuts/149neksqjs/go-nuts-xml-package-problems
// for details.
type MathML struct {
	// The tag name
	XMLName struct{} `xml:"mathml"`

	// The MathML content
	Xml string `xml:",innerxml"`
}

// A Pod corresponds roughly to one category of result. All queries will return
// a Result with at least two Pods (the "Input interpretation" and "Result");
// many queries will return more.
//
// For example, the Result for the query "amanita" contains seven pods, which
// have titles like "Scientific name", "Taxonomy", and "Image", among others.
type Pod struct {
	// The tag name
	XMLName struct{} `xml:"pod"`

	// The pod title
	Title string `xml:"title,attr"`

	// The pod subpods
	Subpods []Subpod `xml:"subpod"`

	// The internal identifier for the pod type
	ID string `xml:"id,attr"`

	// The name of the scanner that produced the pod
	Scanner string `xml:"scanner,attr"`

	// A number indicating the intended position of the pod in a visual display
	// (the uppermost pod should have the lowest position)
	Position int `xml:"position,attr"`

	// Whether the pod couldn't be processed
	Errored bool `xml:"error,attr"`

	// Whether the pod is the query's primary pod
	Primary bool `xml:"primary,attr"`
}

// A Reinterpretation occurs when Wolfram Alpha cannot understand a query and
// replaces it with a new query that seems close in meaning to the original.
//
// For example, the nonsensical query "blue mustang moon" might be replaced by
// the query "mustang moon," the name of a 2002 book by Terri Farley.
type Reinterpretation struct {
	// The tag name
	XMLName struct{} `xml:"reinterpret"`

	// The new query
	Query string `xml:"new,attr"`

	// A message that could be displayed to the user before showing the new query
	// (usually "Using closest Wolfram|Alpha interpretation:")
	Message string `xml:"text,attr"`

	// A value from 0 to 1 indicating how similar the new query is to the original
	// query
	Score float32 `xml:"score,attr"`

	// A description ("low", "medium", or "high") indicating how similar the new
	// query is to the original query
	Level string `xml:"level,attr"`
}

// A Result represents the Wolfram Alpha API's response to a single query.
// Results are returned from a Client when a query is made.
type Result struct {
	// The tag name
	XMLName struct{} `xml:"queryresult"`

	// The internal identifier for the result
	ID string `xml:"id,attr"`

	// The result pods
	Pods []Pod `xml:"pod"`

	// The query assumptions, if any were made
	Assumptions []Assumption `xml:"assumption"`

	// The example page, if the query referred to a general topic
	ExamplePage *ExamplePage `xml:"examplepage"`

	// The future topic, if the query concerned a topic under development
	FutureTopic *FutureTopic `xml:"futuretopic"`

	// The language message, if the query was not in English
	LanguageMessage *LanguageMessage `xml:"languagemsg"`

	// The query reinterpretation, if the query was reinterpreted
	Reinterpretation *Reinterpretation `xml:"reinterpret"`

	// Alternative queries, close in spelling or meaning to the original, if any
	Suggestions []string `xml:"didyoumean"`

	// Tips for the user, if any
	Tips []Tip `xml:"tips>tip"`

	// The sources used to compute the result, if any
	Sources []Source `xml:"source"`

	// Whether the input was understood
	Succeeded bool `xml:"success,attr"`

	// Whether the query couldn't be processed
	Errored bool `xml:"error,attr"`

	// The error, if the query couldn't be processed
	Error Error `xml:"error"`

	// A URL to recalculate the query and get more pods, if there were errors
	Recalculate string `xml:"recalculate,attr"`

	// A comma-separated list of the types of data represented in the result
	DataTypes string `xml:"datatypes,attr"`

	// The wall clock time to parse the query, in seconds
	ParseTiming float32 `xml:"parsetiming,attr"`

	// Whether the parsing stage timed out
	ParseTimedOut bool `xml:"parsetimedout,attr"`

	// The wall clock time to generate the result, in seconds
	Timing float32 `xml:"timing,attr"`

	// A comma-separated list of the IDs of pods that timed out
	TimedOut string `xml:"timedout,attr"`

	// The API version
	Version string `xml:"version,attr"`
}

// A Source provides a link to a web page with source information. Sources are
// found in Results that do not exclusively use "common knowledge" or purely
// mathematical computation.
//
// Note that the URL links not directly to a source, but rather to a Wolfram
// Alpha webpage with source information for a general topic. See
// http://www.wolframalpha.com/sources/GivenNameDataSourceInformationNotes.html
// as an example.
type Source struct {
	// The tag name
	XMLName struct{} `xml:"source"`

	// The address of the web page with source information
	URL string `xml:"url,attr"`

	// A short description of the source
	Description string `xml:"text,attr"`
}

// A Subpod contains a distinct result or image for a Pod. Each Subpod may
// include various representations of a single datum, depending on what formats
// were requested and what is relevant to the query.
//
// At the very least, all Subpods will have an image (possibly a picture of
// text).
type Subpod struct {
	// The tag name
	XMLName struct{} `xml:"subpod"`

	// The subpod title, usually an empty string
	Title string `xml:"title,attr"`

	// The subpod plaintext representation, if available
	Plaintext string `xml:"plaintext"`

	// The subpod image, if available
	Image *Image `xml:"img"`

	// The subpod MathML representation, if available
	MathML *MathML `xml:"mathml"`

	// The Mathematica input, if available
	MathematicaInput string `xml:"minput"`

	// The Mathematica output, if available
	MathematicaOutput string `xml:"moutput"`

	// Whether the subpod is the query's primary subpod
	Primary bool `xml:"primary,attr"`
}

// A Tip offers a suggestion to the user for improving future queries. Tips
// usually occur when Wolfram Alpha cannot understand the input. For example, a
// tip might suggest, "Check your spelling and use English."
//
// Though it only has one field, this type is needed for unmarshaling the text
// attribute of tip elements due to limitations of the encoding/xml package.
type Tip struct {
	// The tag name
	XMLName struct{} `xml:"tip"`

	// The tip message
	Message string `xml:"text,attr"`
}
