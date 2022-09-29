package main

import (
	"fmt"
	"strings"
)

type GraphObject struct {
	Name, Color string
	Children    []GraphObject
}

func (g *GraphObject) String() string {
	sb := strings.Builder{}
	g.print(&sb, 0)
	return sb.String()
}

func (g *GraphObject) print(sb *strings.Builder, depth int) {
	sb.WriteString(strings.Repeat("*", depth))
	if len(g.Color) > 0 {
		sb.WriteString(g.Color)
		sb.WriteRune(' ')
	}
	sb.WriteString(g.Name)
	sb.WriteRune('\n')

	for _, child := range g.Children {
		child.print(sb, depth+1)
	}
}

func newCircle(color string) *GraphObject {
	return &GraphObject{"Circle", color, nil}
}

func newSquare(color string) *GraphObject {
	return &GraphObject{"Square", color, nil}
}

func main() {
	drawing := GraphObject{"MyDrawing", "", nil}
	drawing.Children = append(drawing.Children, *newCircle("Red"))
	drawing.Children = append(drawing.Children, *newCircle("Yellow"))

	group := GraphObject{"Group1", "", nil}
	group.Children = append(group.Children, *newCircle("Blue"))
	group.Children = append(group.Children, *newSquare("Blue"))

	drawing.Children = append(drawing.Children, group)
	fmt.Println(drawing.String())
}
