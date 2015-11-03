package api

import (
	"errors"
	"fmt"
	"net/url"
)

// An Assumption defines a single assumption, typically about the meaning of a
// word or phrase in the query, and a series of possible other values.
//
// For instance, Wolfram Alpha will assume that the query "pi" refers to the
// mathematical constant, as opposed to the Greek character, the movie, etc. The
// Result for this query will have an Assumption with this information.
type Assumption struct {
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
// mathematical constant pi, for the Greek character pi, for the movie Pi, etc.
type AssumptionValue struct {
	// The internal identifier for the assumption value
	Name string `xml:"name,attr"`

	// A description of the assumption suitable for display to the user
	Description string `xml:"desc,attr"`

	// The query value needed to invoke this assumption in a subsequent query
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
type Pod struct {
	// The pod title
	Title string `xml:"title,attr"`

	// The name of the scanner that produced the pod
	Scanner string `xml:"scanner,attr"`

	// The internal identifier for the pod type
	ID string `xml:"id,attr"`

	// An integer (often a multiple of 100) indicating the intended position of
	// the pod in a visual display. Pods with a smaller position should be
	// displayed above pods with a greater position.
	Position int `xml:"position,attr"`

	// Whether the pod failed to be processed
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
	// (usually "Using closest Wolfram|Alpha interpretation:")
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

// A Result represents the Wolfram Alpha API's response to a single query.
// Results are returned from a Client when a query is made.
type Result struct {
	// The internal identifier for the result
	ID string `xml:"id,attr"`

	// Whether the input was understood
	Succeeded bool `xml:"success,attr"`

	// Whether a serious processing error occured
	Errored bool `xml:"error,attr"`

	// The error, if a serious processing error occured
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

	// The result pods
	Pods []Pod `xml:"pod"`

	// The query assumptions, if any were made
	Assumptions []Assumption `xml:"assumption"`

	// The example page, if the query referred to a general topic
	ExamplePage ExamplePage `xml:"examplepage"`

	// The topic name, if the query concerned a topic under development
	FutureTopic string `xml:"futuretopic>topic,attr"`

	// The language message, if the query was not in English
	LanguageMessage LanguageMessage `xml:"languagemsg"`

	// The query reinterpretation, if the query was reinterpreted
	Reinterpretation Reinterpretation `xml:"reinterpret"`

	// Alternative queries, close in spelling or meaning to the original, if any
	Suggestions []string `xml:"didyoumean"`

	// Tips (e.g., "Check your spelling and use English") for the user, if any
	Tips []string `xml:"tips>tip>text,attr"`

	// The sources used to compute the result, if any
	Sources []Source `xml:"source"`

	// The API version
	Version string `xml:"version,attr"`
}

// PrimaryText returns the first primary pod's plaintext representation
func (res Result) PrimaryText() (text string, err error) {
	for pod := range res.Pods {
		if pod.Primary {
			if len(pod.Subpods) == 0 {
				return "", errors.New("no subpods in first primary pod")
			}
			return pod.Subpods[0].Plaintext
		}
	}
	return "", errors.New("no primary pods")
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
	// The address of the web page with source information
	URL string `xml:"url,attr"`

	// A short description of the source
	Description string `xml:"text,attr"`
}

// A Subpod contains a distinct result or image for a Pod. Each Subpod may
// include various representations of the result, depending on what formats
// were requested and what is relevant to the query.
//
// At the very least, all Subpods will have an image (possibly a picture of
// text).
type Subpod struct {
	// The subpod title, usually an empty string
	Title string `xml:"title,attr"`

	// Whether the subpod is the query's primary subpod
	Primary bool `xml:"primary,attr"`

	// The subpod plaintext representation, if available
	Plaintext string `xml:"plaintext"`

	// The subpod image, if available
	Image Image `xml:"img"`

	// The subpod MathML representation, if available
	MathML string `xml:"mathml,innerxml"`

	// The Mathematica input, if available
	MathematicaInput string `xml:"minput"`

	// The Mathematica output, if available
	MathematicaOutput string `xml:"moutput"`
}
