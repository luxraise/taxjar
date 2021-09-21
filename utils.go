package taxjar

import (
	"io"
	"net/url"
	"strings"
)

func getRequestBody(request Request) (body io.Reader) {
	if request == nil {
		return
	}

	encoded := request.ToFormValues().Encode()
	body = strings.NewReader(encoded)
	return
}

type Request interface {
	ToFormValues() url.Values
}
