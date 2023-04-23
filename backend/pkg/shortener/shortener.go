package shortener

import (
	"encoding/base64"
	"encoding/hex"
	_ "encoding/hex"
)

var b64 = base64.NewEncoding("-_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func Encode(input []byte) ([]byte, error) {
	db := make([]byte, hex.DecodedLen(len(input)))
	_, err := hex.Decode(db, input)
	if err != nil {
		return nil, err
	}

	eb := make([]byte, b64.EncodedLen(len(db)))
	b64.Encode(eb, db)

	return eb, nil
}

func Decode(input []byte) ([]byte, error) {
	db := make([]byte, b64.DecodedLen(len(input)))
	_, err := b64.Decode(db, input)
	if err != nil {
		return nil, err
	}

	eb := make([]byte, hex.EncodedLen(len(db)))
	hex.Encode(eb, db)

	return eb, nil
}
