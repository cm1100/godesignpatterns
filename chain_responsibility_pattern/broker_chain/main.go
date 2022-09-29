package main

import (
	"fmt"
	"sync"
)

// Cor , Mediator , observer , CQS

type Argument int

const (
	Attack Argument = iota
	Defence
)

type Query struct {
	CreatureName string
	WhatToQuery  Argument
	Value        int
}

type Observer interface {
	Handle(q *Query)
}

type Observable interface {
	Suscribe(o Observer)
	UnSuscribe(o Observer)
	Fire(q *Query)
}

type Game struct {
	observers sync.Map
}

func (g *Game) Suscribe(o Observer) {
	g.observers.Store(o, struct{}{})
}

func (g *Game) UnSuscribe(o Observer) {
	g.observers.Delete(o)
}

func (g *Game) Fire(q *Query) {
	g.observers.Range(func(key, value any) bool {
		if key == nil {
			return false
		}
		key.(Observer).Handle(q)
		return true
	})
}

type Creature struct {
	game            *Game
	Name            string
	attack, defence int
}

func NewCreature(game *Game, name string, attack int, defence int) *Creature {
	return &Creature{game: game, Name: name, attack: attack, defence: defence}
}

func (c *Creature) Attack() int {
	q := Query{c.Name, Attack, c.attack}
	c.game.Fire(&q)
	fmt.Println(q.Value)
	return q.Value
}

func (c *Creature) Defence() int {
	q := Query{c.Name, Defence, c.defence}
	c.game.Fire(&q)
	return q.Value
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.Defence(), c.Attack())
}

type CreatureModifier struct {
	game     *Game
	creature *Creature
}

func (c CreatureModifier) Handle(q *Query) {
	//
}

type DoubleAttackModifier struct {
	CreatureModifier
}

func (d *DoubleAttackModifier) Close() error {
	d.game.UnSuscribe(d)
	return nil
}

func NewDoubleAttackModifier(c *Creature, g *Game) *DoubleAttackModifier {
	d := &DoubleAttackModifier{CreatureModifier{g, c}}
	g.Suscribe(d)
	return d
}

func (d *DoubleAttackModifier) Handle(q *Query) {
	fmt.Println("in here")
	if q.CreatureName == d.creature.Name && q.WhatToQuery == Attack {
		q.Value *= 2
	}

}

func main() {

	game := &Game{sync.Map{}}
	goblin := NewCreature(game, "Strong Goblin", 2, 2)
	fmt.Println(goblin.String())

	{
		m := NewDoubleAttackModifier(goblin, game)
		fmt.Println(goblin.String())
		m.Close()
	}

	fmt.Println(goblin.String())
}
