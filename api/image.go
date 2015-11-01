package api

import (
	"fmt"
	"strings"
)

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
