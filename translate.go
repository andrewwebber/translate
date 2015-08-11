package translate

import "errors"

const LANG_ENGLISH = 0
const LANG_GERMAN = 1

var ERR_NOTFOUND = errors.New("not found")
var ERR_NOTSUPPORTED = errors.New("not supported language")

type Repo interface {
	Translate(key string, destinationLanguage string) (string, error)
}

type SimpleRepo struct {
	data map[string]string
}

func NewSimpleRepro(data map[string]string) SimpleRepo {
	return SimpleRepo{data: data}
}

func (s SimpleRepo) Translate(key string, destinationLanguage int) (string, error) {
	switch destinationLanguage {
	case LANG_ENGLISH:
		if value, ok := s.data[key]; ok {
			return value, nil
		}

		return "", ERR_NOTFOUND
	case LANG_GERMAN:
		return "", ERR_NOTSUPPORTED
	}

	return "", ERR_NOTSUPPORTED
}
