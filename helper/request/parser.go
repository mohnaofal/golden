package request

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/form"
)

const (
	tagName = "json"
	maxSize = 4 << 20
)

func Form(r *http.Request, i interface{}) error {
	if r.Body == nil {
		return nil
	}
	// nolint
	defer r.Body.Close()

	ct := r.Header.Get("Content-Type")
	if strings.HasPrefix(ct, "application/json") || strings.HasPrefix(ct, "text/json") {
		if err := json.NewDecoder(r.Body).Decode(i); err != nil {
			return err
		}
	} else if strings.HasPrefix(ct, "multipart/form-data") {
		if err := r.ParseMultipartForm(maxSize); err != nil {
			return err
		}
	} else {
		// default parse is urlencoded form type
		if err := r.ParseForm(); err != nil {
			return err
		}
	}

	d := form.NewDecoder()
	d.SetTagName(tagName)

	return d.Decode(i, r.Form)
}

func Query(r *http.Request, i interface{}) error {
	d := form.NewDecoder()
	d.SetTagName(tagName)

	return d.Decode(i, r.URL.Query())
}
