package api

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var elementTests = []struct {
	Msg      string
	XML      string
	Value    interface{}
	Computed map[string]interface{}
}{
	{
		Msg: "Basic unmarshal of <error>",
		XML: `<error>
		        <code>2</code>
		        <msg>Appid missing</msg>
		      </error>`,
		Value: &Error{Code: 2, Message: "Appid missing"},
	},
	{
		Msg: "Basic unmarshal of <examplepage>",
		XML: `<examplepage category='ChemicalCompounds'
		        url='http://wolframalpha.com/examples/....htm'/>`,
		Value: &ExamplePage{
			Topic: "ChemicalCompounds",
			URL:   "http://wolframalpha.com/examples/....htm",
		},
	},
	{
		Msg: "Basic unmarshal of <img>",
		XML: `<img src="http://wolframalpha.com/53?MSPStoreType=image/gif"
		        alt="x = 0"
		        title="x = 0"
		        width="36"
		        height="18"/>`,
		Value: &Image{
			URL:    "http://wolframalpha.com/53?MSPStoreType=image/gif",
			Alt:    "x = 0",
			Title:  "x = 0",
			Width:  36,
			Height: 18,
		},
		Computed: map[string]interface{}{
			"Mime": "image/gif",
			"HTML": `<img src="http://wolframalpha.com/53?MSPStoreType=image/gif" ` +
				`alt="x = 0" title="x = 0" width="36" height="18"/>`,
		},
	},
}

func TestElements(t *testing.T) {
	for _, test := range elementTests {
		vt := reflect.TypeOf(test.Value)
		dest := reflect.New(vt.Elem()).Interface()
		xml.Unmarshal([]byte(test.XML), &dest)

		if test.Value != nil {
			assert.EqualValues(t, test.Value, dest, test.Msg)
		}

		if test.Computed != nil {
			for name, exp := range test.Computed {
				got := reflect.ValueOf(dest).
					MethodByName(name).
					Call([]reflect.Value{})[0].
					String()
				assert.EqualValues(t, exp, got, test.Msg)
			}
		}
	}
}
