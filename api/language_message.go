package api

// A LanguageMessage is used to represent a <languagemsg> element, an element
// used by Wolfram Alpha to provide details when it recognizes that your query
// is in a foreign language.
type LanguageMessage struct {
	// The message in English
	English string `xml:"english,attr"`

	// The message in the same language as the query
	Other string `xml:"other,attr"`
}
