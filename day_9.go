package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Position struct {
	x int
	y int
}

type RopeNode struct {
	pos     Position
	visited map[Position]bool
	tail    *RopeNode
}

type Move struct {
	dir   string
	steps int
}

var dirs = map[string]Position{
	"U": {0, -1},
	"D": {0, 1},
	"L": {-1, 0},
	"R": {1, 0},
}

func parseMoves(input string) (moves []Move) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		var m Move
		fmt.Sscanf(line, "%s %d", &m.dir, &m.steps)
		moves = append(moves, m)
	}
	return moves
}

func makeRope(count int) *RopeNode {
	node := newNode()
	prev := node
	for i := 1; i < count; i++ {
		prev.tail = newNode()
		prev = prev.tail
	}
	return node
}

func newNode() *RopeNode {
	k := &RopeNode{visited: make(map[Position]bool)}
	k.move(Position{})
	return k
}

func (node *RopeNode) getNth(count int) *RopeNode {
	if count == 0 {
		return node
	}
	return node.tail.getNth(count - 1)
}

func runMoves(knot *RopeNode, moves []Move) *RopeNode {
	for _, m := range moves {
		for s := 0; s < m.steps; s++ {
			knot.move(dirs[m.dir])
		}
	}
	return knot
}

func (node *RopeNode) move(v Position) {
	node.pos = node.pos.updatePosition(v)
	node.visited[node.pos] = true
	if node.tail != nil {
		node.tail.follow(node)
	}
}

func (p Position) updatePosition(v Position) Position {
	return Position{p.x + v.x, p.y + v.y}
}

func (tail *RopeNode) follow(head *RopeNode) {
	if tail.pos.touching(head.pos) {
		return
	}
	dx := (head.pos.x-tail.pos.x)/2 + (head.pos.x-tail.pos.x)%2
	dy := (head.pos.y-tail.pos.y)/2 + (head.pos.y-tail.pos.y)%2
	tail.move(Position{dx, dy})
}

func (node1 Position) touching(node2 Position) bool {
	return math.Abs(float64(node1.x)-float64(node2.x)) < 1.5 && math.Abs(float64(node1.y)-float64(node2.y)) < 1.5
}

func main() {
	input, _ := os.ReadFile("input_d9.txt")
	moves := parseMoves(string(input))
	head := makeRope(10)
	head = runMoves(head, moves)
	fmt.Println("Head and Tail:", len(head.getNth(1).visited))
	fmt.Println("I am tired of this motherfucking Snake on this motherfucking Puzzle:", len(head.getNth(9).visited))
}
