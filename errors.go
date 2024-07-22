package main

import "errors"

var (
	ErrIncorrectParticipantCount error = errors.New(
		"minimum number of participants is 4",
	)
	ErrEmptyBracket   error = errors.New("bracket cannot be empty")
	ErrEmptyBracketID error = errors.New(
		"bracket id cannot be empty",
	)
	ErrParticipantsNotSameType error = errors.New(
		"participants must have the same type",
	)
)
