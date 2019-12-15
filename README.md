# sortutil

This is a fork of Patrick Mylund Nielsen's `sortutil` package. Its chief use
case is to sort structs by arbitrary field names.

## Rationale
When return JSON representing the results of queries against a database, it's
normal to return an array of JSON objects representing database rows. These rows
can be sorted by arbitrary column names, but if we want to sort the data in
variable ways, we should not have to repeatedly call the database and have it
sort the data; if we have many thousands of rows cached, we should be able to
just retrieve the cached rows, sort them appropriately, and return the results.
The rows are represented in our cache as slices of structs, so we need to be
able to sort by struct field names.

## Why Fork?
Forking this package was necessary for two reasons. First, Nielsen's package
causes a runtime panic when sorting a struct by a field whose name is not found;
we don't want to have to put in recovery code just so we can sort our data.
Second, it is common to ignore case when querying database tables and columns:
"SELECT col FROM mytable" should query the same data as "SELECT COL FROM
MYTABLE". Nielsen's package does not permit selecting struct field names in a
case-insensitive manner.

## Usage

```Go
import "sortutil"

var (
	names     = []string{"Adam", "Eric", "Carol", "Bob", "Daniel"}
	ages      = []int{21, 33, 14, 9, 39}
	addresses = []string{"123 Main", "88 Farview", "11 University", "Johnson Lane", "98 Varsity"}
)

struct Entry {
    Name string
    Age int
    Address string
}

func sortEntriesCaseInsensitive(entries []Entry, name string, order sortutil.Order) {
    if err := sortutil.ByCiField(entries, name, sortutil.Ascending); err == nil {
        for _, e := range entries {
            fmt.Println(e.Name) 
        }
    } else {
        fmt.Println(err.Error())
    }
}

func main() {
	var entries []Entry
	for i := 0; i < len(names); i++ {
		entries = append(entries, NewEntry(names[i], ages[i], addresses[i]))
	}
    sortEntriesCaseInsensitive(entries, "Name", sortutil.Ascending)
    /*
      Adam
      Bob
      Carol
      Daniel
      Eric
    */
    sortEntriesCaseInsensitive(entries, "NAME", sortutil.Descending)
    /*
      Eric
      Daniel
      Carol
      Bob
      Adam
    */
    sortEntriesCaseInsensitive(entries, "age", sortutil.Ascending)
    /*
      Bob
      Carol
      Adam
      Eric
      Daniel
    */
    sortEntriesCaseInsensitive(entries, "dog", sortutil.Ascending)
    /*
      error getting values for field "dog" from struct with fields "Name","Age","Address": reflected value is invalid
    */
}
```