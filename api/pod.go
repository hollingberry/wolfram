package api

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
