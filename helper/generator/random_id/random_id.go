package randomid

import (
	"github.com/teris-io/shortid"
)

func GenerateID() string {
	unique := ``
	sid, err := shortid.New(1, shortid.DefaultABC, 2342)
	if err != nil {
		return unique
	}

	unique, err = sid.Generate()
	if err != nil {
		return unique
	}

	return unique
}
