package main

// Dictionary a simple key value store
type Dictionary map[string]string

const (
	// ErrNotFound is used when the word is not found in the dictionary
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	// ErrWordExists is used when trying to add the same key to the dictionary
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	// ErrWordDoesNotExist used when trying to update a key that doesn't already exist in the dictionary
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

// DictionaryErr is a collection of errors
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Search finds a word in the dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add a word and defintion to the dictionay
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// Update a key with a new definition
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

// Delete an item from the dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
