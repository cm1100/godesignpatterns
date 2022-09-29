package main

import (
	"io/ioutil"
	"strings"
)

type Persistent struct {
	Lineseperator string
}

func (p *Persistent) SaveToFile(j *Journal, filename string) {

	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.enteries, p.Lineseperator)), 0644)

}
