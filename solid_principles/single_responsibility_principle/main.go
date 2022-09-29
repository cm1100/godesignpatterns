package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

var entryCount = 0

type Journal struct {
	enteries []string
}

func (j *Journal) String() string {
	return strings.Join(j.enteries, "\n")
}

func (j *Journal) AddEntry(text string) int {

	entryCount++
	entry := fmt.Sprintf("%d %s", entryCount, text)
	j.enteries = append(j.enteries, entry)
	return entryCount
}

func (j *Journal) RemoveENtry(index int) {
	//...
}

// seperation of concerns
// God Object
// Responsibility of journal is not to deal with persistence
// Persistence can be handkeked by another package

func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {

}

func (j *Journal) LoadFromWeb(url *url.URL) {

}

func main() {

	j := Journal{}
	j.AddEntry("i played today")
	j.AddEntry("i ate a bug")
	fmt.Println(j.String())

	p := Persistent{Lineseperator: "\n"}
	p.SaveToFile(&j, "journal.txt")
}
