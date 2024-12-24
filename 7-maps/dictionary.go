package main

const (
	ErrNotFound     = DictionaryError("not found")
	ErrAlreadyFound = DictionaryError("already found")
)

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(term string) (string, error) {
	val, ok := d[term]

	if !ok {
		return "", ErrNotFound
	}

	return val, nil
}

func (d Dictionary) Add(term, val string) error {
	_, err := d.Search(term)

	switch err {
	case ErrNotFound:
		d[term] = val
	case nil:
		return ErrAlreadyFound
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(term, val string) error {
	_, err := d.Search(term)

	if err != nil {
		return err
	}

	d[term] = val

	return nil
}

func (d Dictionary) Delete(term string) error {
	_, err := d.Search(term)

	if err != nil {
		return err
	}

	delete(d, term)

	return nil
}
