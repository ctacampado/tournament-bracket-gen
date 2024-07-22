package main

import (
	"fmt"
	"math"
	"time"

	"github.com/google/uuid"
)

// Generate returns a bracket depending on the number of participants.
// The bracket is generated using a perfect binary tree. The minimum number
// of participants is BracketMin
func Generate(ps []*Participant, bo int) (*Bracket, error) {
	if len(ps) < BracketMin {
		return nil, ErrIncorrectParticipantCount
	}
	// to create a perfect binary tree, we need n to be power of 2.
	// if num of participants is not power of 2, increment until it is.
	n := toPow2(len(ps))
	bracket := initBracket(n, bo)
	// find initial match to assign participant
	for i, p := range ps {
		var m *Match
		// this makes the traversal path to alternate between left and right
		if i%2 == 0 {
			m = findMatch(bracket.Matches[0].Left)
		} else {
			m = findMatch(bracket.Matches[0].Right)
		}
		if m != nil {
			if m.ParticipantA == nil {
				m.ParticipantA = p
			} else if m.ParticipantB == nil {
				m.ParticipantB = p
				m.IsFull = true
			}
		}
	}
	return bracket, nil
}

// findMatch traverses the bracket.
// The bracket is represented by a perfect binary tree.
// The match that we are trying to find is a leaf node of the tree.
func findMatch(m *Match) *Match {
	var mm *Match
	if m.Left != nil {
		if !m.Left.IsFull {
			mm = findMatch(m.Left)
			// if found nothing, try right
			if mm == nil {
				mm = findMatch(m.Right)
			}
			return mm
		} else if !m.Right.IsFull {
			mm = findMatch(m.Right)
			// if found nothing, try left
			if mm == nil {
				mm = findMatch(m.Left)
			}
			return mm
		}
		// only middle nodes available
		return nil
	}
	// at the leaf node
	return m
}

// initBracket returns a bracket that is represented by a perfect binary tree.
// A node in a perfect binary tree is represented by a match and the number of
// matches will always be equal to n-1, where n = number of participants.
func initBracket(n int, bo int) *Bracket {
	numMatches := n - 1
	matches := make([]*Match, numMatches)
	now := time.Now()
	// initialize all nodes
	for i := range matches {
		// added the index as prefix to easily inspect the tree
		// when logging or viewing it as JSON data
		uuid := fmt.Sprintf("%d-%s", i, uuid.NewString())
		matches[i] = &Match{UUID: uuid, BestOf: bo, UpdatedAt: &now}
		if i == 0 {
			matches[i].IsFinal = true
		}
	}
	// assign children nodes
	for i := range matches {
		// left child position in the array
		left := i*2 + 1
		if left > len(matches)-1 {
			// all children has been assigned
			break
		}
		// right child position in the array
		right := i*2 + 2
		matches[i].Left = matches[left]
		matches[i].Right = matches[right]
	}
	return &Bracket{
		UUID:      uuid.NewString(),
		Matches:   matches,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func toPow2(n int) int {
	for {
		if isPow2(n) {
			break
		}
		n++
	}
	return n
}

func isPow2(n int) bool {
	return math.Mod(math.Log2(float64(n)), 1.0) == 0
}
