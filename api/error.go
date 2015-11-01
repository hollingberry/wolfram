package api

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
	// A short message describing the error
	Message string `xml:"msg"`

	// The error code
	Code int `xml:"code"`
}
