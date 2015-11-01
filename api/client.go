package api

// A Format defines a format in which results will be returned. Multiple formats
// can be requested for a single request, although not all requested formats
// will necessarily be present in each pod.
type Format int

const (
	Plaintext Format = iota
	Image
	MathematicaInput
	MathematicaOutput
	Cell
	MathML
	ImageMap
	Sound
	Wav
)

// A UnitSystem defines a system of units.
type UnitSystem int

const (
	// The imperial (or U.S.) system of units
	Imperial UnitSystem = iota

	// The metric system
	Metric

	// The system of units used in your location
	Location
)

type Client struct {
	// The AppID for your application
	AppID string

	// The desired output formats for each pod
	Formats []Format

	// The optimal width, in pixels, for pod images. Wolfram Alpha will try to
	// keep the widths of images under this value, but will make the images wider
	// (up to ImageMaxWidth) if ugly line breaks would be used at the smaller
	// size.
	ImageWidth int

	// The maximum width, in pixels, for pod images
	ImageMaxWidth int

	// The magnification of pod images. Magnification does not affect the pixel
	// width of images, but rather the size of the content in them.
	ImageMagnification int

	// Optimal width, in pixels, for plots and other graphics. There are many
	// graphics in Wolfram Alpha that are deliberately rendered at larger sizes to
	// accommodate their content. Specifying plot width is currently an
	// experimental feature that does not yet affect many type of graphics.
	ImagePlotWidth int

	// The user's IP address (for queries that use location data). Use this option
	// to override what Wolfram Alpha thinks your current IP address is.
	IPAddress string

	// The user's latitude/longitude (for queries that use location data). This
	// should be a comma-separated value like "40.42,-3.71".
	LatLong string

	// The user's location (for queries that use location data). This should be a
	// place name like "Los Angeles, CA" or "Madrid".
	Location string

	// If true, then Wolfram Alpha will try to reinterpret queries that it cannot
	// understand.
	Reinterpret bool

	// The user's preferred measurement system.
	Units UnitSystem
}

func NewClient(id string) {
	return Client{
		AppID: id,
	}
}

func (c *Client) Query(input string) Result {
}

func (c *Client) Validate(input string) Result {
}

func (c *Client) Ask(input string) string {
}
