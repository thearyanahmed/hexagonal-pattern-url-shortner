package json

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/thearyanahmed/url-shortener/shortener"
)

type Redirect struct{}

func (r *Redirect) Decode(input []byte) (* shortener.Redirect, error) {
	redirect := &shortener.Redirect{}

	if err := json.Unmarshal(input,redirect); err != nil {
		return nil, errors.Wrap("serializer.Redirect.Decode")
	}

	return redirect, nil
}

func (r *Redirect) Encode(input *Redirect) ([]byte, error) {
	bytes, err := json.Marshal(input) ; err != nil {
		return nil, errors.Wrap("serializer.Redirect.Endcode")
	}

	return bytes, nil
}
