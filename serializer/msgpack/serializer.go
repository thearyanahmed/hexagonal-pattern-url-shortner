package msgpack

import (
	"github.com/vmihailenco/msgpack"
	"github.com/pkg/errors"
	"github.com/thearyanahmed/url-shortener/shortener"
)

type Redirect struct{}

func (r *Redirect) Decode(input []byte) (* shortener.Redirect, error) {
	redirect := &shortener.Redirect{}

	if err := msgpack.Unmarshal(input,redirect); err != nil {
		return nil, errors.Wrap(err,"serializer.Redirect.Decode")
	}

	return redirect, nil
}

func (r *Redirect) Encode(input *shortener.Redirect) ([]byte, error) {
	bytes, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err,"serializer.Redirect.Endcode")
	}

	return bytes, nil
}
