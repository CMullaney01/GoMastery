package idiomatic

import (
	"fmt"
	"net"
	"sync"
)

//const declarations not all CAPS if we ant to export capitalise
const (
	scalar  = 0.1
	Version = 2
)

// how to constructor
type Order struct {
	Size float64
}

func NewOrder(size float64) *Order {
	return &Order{
		Size: size,
	}
}

// how to replicate enums (kinda) if we need it
type Suit byte

const (
	SuitHearts Suit = iota
	SuitClubs
	SuitDiamonds
	SuitSpades
)

///struct initialisations
type Vector struct {
	x int
	y int
}

// variable grouping
func Foo() int {
	var (
		x   = 100
		y   = 2
		foo = " "
	)

	fmt.Println(foo)
	return x + y
}

// if a function panics we must perfix with Must
func MustParseIntFromString(s string) int {
	//logic
	// if a function panics instead of error
	panic("oops")
	return 10, nil
}

//mutex grouping -- notice we can easily tell what mutex is protecting what
type Server struct {
	listAddr  string
	isRunning bool

	mu    sync.RWMutex
	peers map[string]net.Conn

	otherMu sync.RWMutex
	other   map[string]net.Conn
}

//interface declaration and namings interfaces should end in er storER writER readER
// interface composition make your interface do small amounts of things
type Storer interface {
	Getter
	Putter
}

type Getter interface {
	Get()
}

type Putter interface {
	Put()
}

//function grouping -- more important == higher up contants at very top (same logic with variables)
//very important Exported first
func veryImportantFuncExported() {}

// now inmportant functions but are local
func veryImportantFunc() {}

//simple util demoted to bottom as fi someone is erading they dont need to know a simple util first
func simpleUtil() {}

func handleGetUserById() {}

func handleResizeImage()

func Idiomatic() {
	//never do this:
	// vector := Vector{10, 20}
	// always do this
	vector := Vector{
		x: 10,
		y: 20,
	}

}
