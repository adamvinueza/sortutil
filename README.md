# sortutil

This is a fork of Patrick Mylund Nielsen's `sortutil` package.

Its purpose is to enable sorting with greater flexibility than is allowed by
Go's built-in `sort` package. It is a _lot_ slower than Go's built-in sorting,
because it relies upon reflection. Its two major use cases are sorting slices of
structs: either by values of arbitrary named fields, or by values of indexed
fields.

Sorting of string values allows for ignoring case. Also, sorting by named fields
allows for ignoring case. (In that situation, your struct obviously should not
have fields whose names differ only in case.)

## Rationale
When returning JSON representing the results of queries against a database, it's
normal to return an array of JSON objects representing database rows. These rows
can be sorted by arbitrary column names, but if we want to sort the data in
variable ways, we should not have to repeatedly call the database and have it
sort the data; if we have many thousands of rows cached, we should be able to
just retrieve the cached rows, sort them appropriately, and return the results.
The rows are represented in our cache as slices of structs, so we need to be
able to sort by struct field names.

## Usage

To sort a struct by a field name:
```Go
import "sortutil"

type Entry struct {
    Name string
    Age int
    Address string
}

var entries []Entry{
    Entry{ Name: "Bob", Age: 21, Address: "100 Main St" },
    Entry{ Name: "Carol", Age: 22, Address: "200 Main St"},
}

err := sortutil.ByField(entries, "Name", Descending)
// Carol, Bob
err = sortutil.ByField(entries, "Age", Ascending)
// Bob, Carol
err = sortutil.ByCiField(entries, "address", Descending)
// Carol, Bob
err = sortutil.ByField(entries, "dog", Ascending)
// err is not nil
```

## Sample Program
```Go
package main

import (
    "fmt"
    "github.com/adamvinueza/sortutil"
)

var (
	names     = []string{"Adam", "Eric", "Carol", "Bob", "Daniel"}
	ages      = []int{21, 33, 14, 9, 39}
	addresses = []string{"123 Main", "88 Farview", "11 University", "Johnson Lane", "98 Varsity"}
)

type Entry struct {
    Name string
    Age int
    Address string
}

func NewEntry(name string, age int, addr string) Entry {
    return Entry{
        Name: name,
        Age: age,
        Address: addr,
    }
}

func sortEntriesCaseInsensitive(entries []Entry, name string, order sortutil.Order) {
    if err := sortutil.ByCiField(entries, name, order); err == nil {
        fmt.Printf("ENTRIES (entry.Name sorted by field \"%s\", %s):\n", name, order)
        for _, e := range entries {
            fmt.Printf("\t%s\n", e.Name)
        }
    } else {
        fmt.Printf("ERROR: %s\n", err.Error())
    }
}

func main() {
	var entries []Entry
	for i := 0; i < len(names); i++ {
		entries = append(entries, NewEntry(names[i], ages[i], addresses[i]))
	}
    sortEntriesCaseInsensitive(entries, "Name", sortutil.Ascending)
    sortEntriesCaseInsensitive(entries, "NAME", sortutil.Descending)
    sortEntriesCaseInsensitive(entries, "age", sortutil.Ascending)
    sortEntriesCaseInsensitive(entries, "dog", sortutil.Ascending)
}

/*
EXPECTED OUTPUT:
----------------
ENTRIES (entry.Name sorted by field "Name", Ascending):
	Adam
	Bob
	Carol
	Daniel
	Eric
ENTRIES (entry.Name sorted by field "NAME", Descending):
	Eric
	Daniel
	Carol
	Bob
	Adam
ENTRIES (entry.Name sorted by field "age", Ascending):
	Bob
	Carol
	Adam
	Eric
	Daniel
ERROR: error getting values for field "dog" from struct with fields "Name","Age","Address": reflected value is invalid
*/
```
