package api

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssumption(t *testing.T) {
	var assum Assumption
	const assumXML = `<assumption type='Clash' word='pi' count='4'>
	                    <value name='NamedConstant'
	                           desc='a mathematical constant'
	                           input='*C.pi-_*NamedConstant-' />
	                    <value name='Character'
	                           desc='a character'
	                           input='*C.pi-_*Character-' />
	                    <value name='MathWorld'
	                           desc='referring to a definition'
	                           input='*C.pi-_*MathWorld-' />
	                  </assumption>`
	xml.Unmarshal([]byte(assumXML), &assum)
	assert.EqualValues(t, Assumption{
		Type: "Clash",
		Word: "pi",
		Values: []AssumptionValue{
			{
				Name:        "NamedConstant",
				Description: "a mathematical constant",
				Input:       "*C.pi-_*NamedConstant-",
			},
			{
				Name:        "Character",
				Description: "a character",
				Input:       "*C.pi-_*Character-",
			},
			{
				Name:        "MathWorld",
				Description: "referring to a definition",
				Input:       "*C.pi-_*MathWorld-",
			},
		},
	}, assum)
}

func TestError(t *testing.T) {
	var err Error
	const errXML = `<error>
	                  <code>2</code>
	                  <msg>Appid missing</msg>
	                </error>`
	xml.Unmarshal([]byte(errXML), &err)
	assert.EqualValues(t, Error{
		Code:    2,
		Message: "Appid missing",
	}, err)
}

func TestExamplePage(t *testing.T) {
	var expg ExamplePage
	const expgXML = `<examplepage category='ChemicalCompounds'
	                              url='http://wolframalpha.com/examples/....htm'/>`
	xml.Unmarshal([]byte(expgXML), &expg)
	assert.EqualValues(t, ExamplePage{
		Topic: "ChemicalCompounds",
		URL:   "http://wolframalpha.com/examples/....htm",
	}, expg)
}

func TestImage(t *testing.T) {
	var img Image
	const imgXML = `<img src="http://wolframalpha.com/53?MSPStoreType=image/gif"
		                   alt="x = 0"
		                   title="x = 0"
		                   width="36"
		                   height="18"/>`
	xml.Unmarshal([]byte(imgXML), &img)
	assert.EqualValues(t, Image{
		URL:    "http://wolframalpha.com/53?MSPStoreType=image/gif",
		Alt:    "x = 0",
		Title:  "x = 0",
		Width:  36,
		Height: 18,
	}, img)
}

func TestLanguageMessage(t *testing.T) {
	var langmsg LanguageMessage
	const langmsgXML = `<languagemsg english='Wolfram|Alpha does not yet support German.'
	                                 other='Wolfram|Alpha versteht noch kein Deutsch.' />`
	xml.Unmarshal([]byte(langmsgXML), &langmsg)
	assert.EqualValues(t, LanguageMessage{
		English: "Wolfram|Alpha does not yet support German.",
		Other:   "Wolfram|Alpha versteht noch kein Deutsch.",
	}, langmsg)
}

func TestReinterpretation(t *testing.T) {
	var rei Reinterpretation
	const reiXML = `<reinterpret text='Using closest Wolfram|Alpha interpretation:'
	                             new='mustang moon'
	                             score='0.705882'
	                             level='high'>
	                  <alternative score='0.386685' level='medium'>blue moon</alternative>
	                </reinterpret>`
	xml.Unmarshal([]byte(reiXML), &rei)
	assert.EqualValues(t, Reinterpretation{
		Query:   "mustang moon",
		Message: "Using closest Wolfram|Alpha interpretation:",
		Score:   0.705882,
		Level:   "high",
	}, rei)
}

func TestSource(t *testing.T) {
	var src Source
	const srcXML = `<source url='http://www.wolframalpha.com/sources/...'
	                        text='City data' />`
	xml.Unmarshal([]byte(srcXML), &src)
	assert.EqualValues(t, Source{
		URL:         "http://www.wolframalpha.com/sources/...",
		Description: "City data",
	}, src)
}

func TestPod(t *testing.T) {
	var pod Pod
	const podXML = `<pod title="Input interpretation"
	                     scanner="Identity"
	                     id="Input"
	                     position="100"
	                     error="false"
	                     numsubpods="1">
	                  <subpod title="">
	                    <plaintext>convert 10 feet to meters</plaintext>
	                  </subpod>
	                </pod>`
	xml.Unmarshal([]byte(podXML), &pod)
	assert.EqualValues(t, Pod{
		Title:    "Input interpretation",
		Scanner:  "Identity",
		ID:       "Input",
		Position: 100,
		Errored:  false,
		Primary:  false,
		Subpods: []Subpod{
			{
				Title:             "",
				Plaintext:         "convert 10 feet to meters",
				MathematicaInput:  "",
				MathematicaOutput: "",
				Image:             nil,
			},
		},
	}, pod)
}

func TestSubpod(t *testing.T) {
	var subpod Subpod
	const subpodXML = `<subpod title="The Gods! The Gods!" primary="true">
	                     <plaintext>d/x (4 x^2)</plaintext>
	                     <img src="http://www.wolframalpha.com/ag659?MSPStoreType=image/gif&amp;s=16"
	                          alt="d/x (4 x^2)"
	                          title="d/x (4 x^2)"
	                          width="51" height="37"/>
	                     <minput>(d/x) (4 x^2)</minput>
	                     <moutput>D[4 x^2, x]</moutput>
	                     <mathml><math>4 dx</math></mathml>
	                   </subpod>`
	xml.Unmarshal([]byte(subpodXML), &subpod)
	assert.EqualValues(t, Subpod{
		Title:     "The Gods! The Gods!",
		Plaintext: "d/x (4 x^2)",
		Image: &Image{
			URL:    "http://www.wolframalpha.com/ag659?MSPStoreType=image/gif&s=16",
			Alt:    "d/x (4 x^2)",
			Title:  "d/x (4 x^2)",
			Width:  51,
			Height: 37,
		},
		MathematicaInput:  "(d/x) (4 x^2)",
		MathematicaOutput: "D[4 x^2, x]",
		MathML: &MathML{
			Xml: `<math>4 dx</math>`,
		},
		Primary: true,
	}, subpod)
}
