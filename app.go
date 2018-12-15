package roku

// App has all the information of an application.
type App struct {
	Name    string `xml:",chardata"`
	ID      string `xml:"id,attr"`
	Type    string `xml:"type,attr"`
	SubType string `xml:"subtype,attr"`
	Version string `xml:"version,attr"`
}
