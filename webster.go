package webster

import (
	"errors"
	"io/ioutil"
	"net/http"
	"encoding/xml"
)

const (
	DICT_HOST = "www.dictionaryapi.com"
	DICT_PORT = 80
	DICT_API_URI = "/api/v1/references/collegiate/xml"
	DICT_API_QUERY_PARAM = "key"
)

var wc websterClient

/*func (e *Entry) EntryId(string) {
	return *e.Id
}*/

func lookup(word string) (err error) {
	var body []byte
	var resp *http.Response
	el := EntryList{}

	/*
	 * We only set currentWord on a succesful fetch, so if it's there, no need to refetch
	 */
	if(word == wc.currentWord) {
		return nil
	}


	resp, err = wc.client.Get(wc.url + word + wc.query)
	if(err != nil) {
		return errors.New("Error on connecting: " + err.Error())
	}

	body, err = ioutil.ReadAll(resp.Body)
	if(err != nil) {
		return errors.New("Error while reading response: " + err.Error())
	}

	err = xml.Unmarshal(body, &el)
	if(err != nil) {
		return errors.New("Error parsing XML response body: " + err.Error())
	}

	if(len(el.Entry) != 0) {
		wc.currentWord = word
		wc.currentEntryIdx = 0
		wc.entries = make([]Entry, len(el.Entry))
		copy(wc.entries, el.Entry)
	}

	return nil
}

func WordExists(word string) (exists bool, err error) {
	err = wc.lookup(word)

	if(err != nil) {
		return false, err
	}

	if(EntryCount() > 0) {
		return true, nil
	}

	return false, nil
}

func InitClient(key string, testConn bool) (err error) {
	var resp *http.Response

	if(len(key) == 0) {
		err = errors.New("API key cannot be zero length")
		return err
	}

	wc = websterClient{
		&http.Client{},
		"http://" + DICT_HOST+ DICT_API_URI + "/",
		"?" + DICT_API_QUERY_PARAM + "=" + key,
		DICT_PORT,
		lookup,
		"",
		0,
		[]Entry{},
	}

	if(testConn) {
		resp, err = wc.client.Get(wc.url + "dictionary" + wc.query)
		_ = resp
		return err
	}

	return nil
}

func Fetch(word string) (error) {
	err := wc.lookup(word)
	if(err != nil) {
		return err
	}

	return nil
}

func EntryCount() (int) {
	return len(wc.entries)
}

func CurrentWord() (string, error) {
	if(&wc == nil) {
		return "", errors.New("Webster client has not been initialized, must call webster.Init()")
	}

	return wc.currentWord, nil
}

func FirstEntry() (*Entry) {
	if(len(wc.entries) == 0) {
		return nil
	}

	wc.currentEntryIdx = 0
	return &(wc.entries[0])
}

func CurrentEntry() (*Entry) {
	return &(wc.entries[wc.currentEntryIdx])
}

func NextEntry() (*Entry) {
	wc.currentEntryIdx = wc.currentEntryIdx + 1
	if(wc.currentEntryIdx >= len(wc.entries)) {
		return nil
	}

	return &(wc.entries[wc.currentEntryIdx])
}