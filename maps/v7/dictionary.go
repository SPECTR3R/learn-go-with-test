package dictionary

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrAlreadyDefined   = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot updateee word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrAlreadyDefined
	}

	return nil
}

func (d Dictionary) Update(word string, definition string) error {
	_, err := d.Search(word)

	if err == nil {
		d[word] = definition
		return nil
	}
	return ErrWordDoesNotExist
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
