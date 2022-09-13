package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

// StringListFlag is a flag type that appends each occurence to a list.
//
// For instance: `-f a -f b` => []string{"a", "b"}
type StringListFlag []string

func (s *StringListFlag) String() string {
	return "<StringListFlag>"
}

func (s *StringListFlag) Set(value string) error {
	*s = append(*s, value)
	return nil
}

var values StringListFlag

// keysValues splits a string array into a map of keys and values.
//
// We expect that each item in `strs` will have one of these forms:
// 1. <key>=<value>
// 2. <key>@<value>
//
// For (1), we will add an element <key>=<value>. For (2), we will open
// the file named <value> and include an element <key>=<contents of file <value>>.
func keysValues(strs []string) (map[string]string, error) {
	result := make(map[string]string)
	for i, s := range strs {
		elems := strings.Split(s, "=")
		if len(elems) != 2 {
			return nil, fmt.Errorf("[item %d] expected <key>=<value> but got: %s", i, s)
		}
		key, value := elems[0], elems[1]

		if key[len(key)-1] == '@' {
			cleanKey := key[0 : len(key)-1]
			// Case 2
			b, err := os.ReadFile(value)
			if err != nil {
				return nil, fmt.Errorf("[item %d] %s: %v", i, s, err)
			}
			result[cleanKey] = string(b)
		} else {
			// Case 1
			result[key] = value
		}
	}
	return result, nil
}

func main() {
	flag.Var(&values, "v", "Key/value args: <key>=<value> or <key>@=<filename>")
	flag.Parse()
	vars, err := keysValues(values)
	if err != nil {
		log.Fatal(err)
	}
	templates := flag.Args()
	for _, t := range templates {
		contents, err := os.ReadFile(t)
		if err != nil {
			log.Fatal(err)
		}
		// log.Printf("loading template: [%s]", string(contents))
		tmpl, err := template.New(t).Parse(string(contents))
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(os.Stdout, vars)
	}
}
