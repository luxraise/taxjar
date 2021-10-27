package taxjar

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

func getRequestBody(request interface{}) (r io.Reader, err error) {
	if request == nil {
		return
	}

	var bs []byte
	if bs, err = json.Marshal(request); err != nil {
		return
	}

	r = bytes.NewReader(bs)
	return
}

func handleResponse(r io.Reader, value interface{}) (err error) {
	if value == nil {
		return
	}

	if err = json.NewDecoder(r).Decode(value); err != nil {
		return fmt.Errorf("error encountered while attempting to decode response as JSON: %v", err)
	}

	return
}

func handleError(r io.Reader) (err error) {
	buf := bytes.NewBuffer(nil)
	if _, err = io.Copy(buf, r); err != nil {
		return
	}

	return errors.New(buf.String())
}
