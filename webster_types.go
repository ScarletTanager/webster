package webster

import (
	"encoding/xml"
	"net/http"
)

type websterClient struct {
	client *http.Client
	url, query string
	port int
	lookup func(string) (error)
	currentWord string
	currentEntryIdx	int
	entries		[]Entry
}

type EntryList struct {
	XMLName		xml.Name	`xml:"entry_list"`
	Entry		[]Entry		`xml:"entry"`
	Suggestion	[]string 	`xml:"suggestion"`
}

type Entry struct {
	XMLName		xml.Name 	`xml:"entry"`
	Id			string		`xml:"id,attr"`
	Ew 			string 		`xml:"ew"`
	Hw 			string 		`xml:"hw"`
	Fl 			string 		`xml:"fl"`	// Functional label (e.g. part of speech)
	Pr 			string 		`xml:"pr"`
	Def 		Def 		`xml:"def"`
}

type Def struct {
	XMLName 	xml.Name 	`xml:"def"`
}
