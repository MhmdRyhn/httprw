package httprw

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type RequestParser struct {
	req *http.Request
}

func NewRequestParser(req *http.Request) *RequestParser {
	return &RequestParser{req: req}
}

/*
Ideally, `loadingObject` should be reference to a struct object
where the request body is to be loaded.
*/
func (rp *RequestParser) Body(loadingObject any) error {
	bodyBytes, err := ioutil.ReadAll(rp.req.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bodyBytes, loadingObject)
	return err
}

func (rp *RequestParser) Headers() map[string][]string {
	return rp.req.Header
}

func (rp *RequestParser) QueryParams() map[string][]string {
	return rp.req.URL.Query()
}
