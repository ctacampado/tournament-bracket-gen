package main

import (
	"time"
)

const (
	BO1        int = 1
	BO3        int = 3
	BO5        int = 5
	BracketMin int = 4
)

type Participant struct {
	UUID string
	Name string
	Type string
	Win  int
	Lose int
}

type Match struct {
	UUID         string
	ParticipantA *Participant
	ParticipantB *Participant
	Left         *Match
	Right        *Match
	BestOf       int
	IsDone       bool
	IsFull       bool
	IsFinal      bool
	StartedAt    *time.Time
	UpdatedAt    *time.Time
	ConcludedAt  *time.Time
}

type Bracket struct {
	UUID        string
	IsDone      bool
	Matches     []*Match
	Champion    *Participant
	CreatedAt   time.Time
	StartedAt   time.Time
	UpdatedAt   time.Time
	ConcludedAt time.Time
}
