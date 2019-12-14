package main

import "fmt"
import "github.com/adamvinueza/sortutil"

var (
	names     = []string{"Adam", "Eric", "Carol", "Bob", "Daniel"}
	ages      = []int{21, 33, 14, 9, 39}
	addresses = []string{"123 Main", "88 Farview", "11 University", "Johnson Lane", "98 Varsity"}
)

func sortEntries(entries []Entry, name string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered")
			return
		}
	}()
	sortutil.AscByCiField(entries, name)
}
func main() {
	var entries []Entry
	for i := 0; i < len(names); i++ {
		entries = append(entries, NewEntry(names[i], ages[i], addresses[i]))
	}
	sortEntries(entries, "ae")
	for _, e := range entries {
		fmt.Println(e)
	}
}

type Entry struct {
	Name    string
	Age     int
	Address string
}

func (e Entry) String() string {
	return fmt.Sprintf(`Name: %s, Age: %d, Address: %s`, e.Name, e.Age, e.Address)
}

func NewEntry(name string, age int, address string) Entry {
	return Entry{
		Name:    name,
		Age:     age,
		Address: address,
	}
}
