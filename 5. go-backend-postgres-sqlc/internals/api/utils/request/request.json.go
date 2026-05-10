package request

import (
	"encoding/json"
	"net/http"
)

func Decode[T any](r *http.Request, dst *T) error {

	dec := json.NewDecoder(r.Body)

	dec.DisallowUnknownFields()

	return dec.Decode(dst)
}