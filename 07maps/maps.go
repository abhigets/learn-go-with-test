package _7pointersanderrors

import "errors"

func Search(dict map[string]string, valueToSearch string) string {
	return dict[valueToSearch]
}

type Dictionary map[string]string

var ErrNotFound = errors.New("error: cannot find string")

func (dict Dictionary) Search(SearchString string) (string, error) {
	value, ok := dict[SearchString]

	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}

func (dict Dictionary) AddWord(key string, value string) {
	dict[key] = value
}
