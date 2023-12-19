package main

import (
	"fmt"
	"sort"
)

type Dictionary struct {
	entries map[string]string
}

func NewDictionary() *Dictionary {
	return &Dictionary{
		entries: make(map[string]string),
	}
}

func (d *Dictionary) Add(word, definition string) {
	d.entries[word] = definition
}

func (d *Dictionary) Get(word string) (string, bool) {
	definition, ok := d.entries[word]
	return definition, ok
}

func (d *Dictionary) Remove(word string) {
	delete(d.entries, word)
}

func (d *Dictionary) List() []string {
	var result []string

	for word, definition := range d.entries {
		result = append(result, fmt.Sprintf("%s: %s", word, definition))
	}

	sort.Strings(result)

	return result
}

func main() {
	dictionary := NewDictionary()

	dictionary.Add("go", "aller")
	dictionary.Add("hello", "bonjour")
	dictionary.Add("world", "monde")

	definition, found := dictionary.Get("go")
	if found {
		fmt.Printf("Définition de 'go': %s\n", definition)
	} else {
		fmt.Println("Mot non trouvé.")
	}

	dictionary.Remove("world")

	wordList := dictionary.List()
	fmt.Println("Liste des mots et définitions :")
	for _, entry := range wordList {
		fmt.Println(entry)
	}
}
