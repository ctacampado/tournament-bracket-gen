package main

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const isLogEnabled = true

type GeneratorTestSuite struct {
	suite.Suite
	isLogEnabled bool
	ctx          context.Context
}

func (suite *GeneratorTestSuite) SetupTest() {
	suite.isLogEnabled = isLogEnabled
	suite.ctx = context.Background()
}

func (suite *GeneratorTestSuite) TestToPow2() {
	testCases := []struct {
		name     string
		n        int
		expected int
	}{
		{"not padded 1", 4, 4},
		{"padded 3", 5, 8},
		{"padded 4", 9, 16},
		{"padded 5", 31, 32},
		{"not padded 2", 64, 64},
	}
	for _, tc := range testCases {
		assert.Equal(suite.T(), tc.expected, toPow2(tc.n), tc.name)
	}
}

func (suite *GeneratorTestSuite) TestInitBracket() {
	testCases := []struct {
		name string
		bo   int
		n    int
	}{
		{"BO1 test 1", BO1, 4},
		{"BO3 test 1", BO3, 4},
		{"BO5 test 1", BO5, 4},
		{"BO1 test 2", BO1, 8},
		{"BO3 test 2", BO3, 8},
		{"BO5 test 2", BO5, 8},
		{"BO1 test 3", BO1, 16},
		{"BO3 test 3", BO3, 16},
		{"BO5 test 3", BO5, 16},
	}
	for _, tc := range testCases {
		b := initBracket(tc.n, tc.bo)
		assert.NotEqual(suite.T(), "", b.UUID)
		assert.Equal(suite.T(), tc.n-1, len(b.Matches))
		for _, m := range b.Matches {
			assert.NotNil(
				suite.T(),
				m,
				fmt.Sprintf("%s: len b.Matches %d", tc.name, len(b.Matches)),
			)
		}
	}
}

func (suite *GeneratorTestSuite) TestGenerate() {
	testCases := []struct {
		name string
		ps   []*Participant
		bo   int
		err  error
	}{
		{
			name: "fail 1",
			ps:   []*Participant{},
			bo:   BO1,
			err:  ErrIncorrectParticipantCount,
		},
		{
			name: "success 1",
			ps: []*Participant{
				{
					Name: "abc",
					Win:  0,
					Lose: 0,
				},
				{
					Name: "def",
					Win:  0,
					Lose: 0,
				},
				{
					Name: "ghi",
					Win:  0,
					Lose: 0,
				},
				{
					Name: "jkl",
					Win:  0,
					Lose: 0,
				},
			},
			bo:  0,
			err: nil,
		},
		{
			name: "success 2",
			ps: []*Participant{
				{
					Name: "abc",
					Win:  0,
					Lose: 0,
				},
				{
					Name: "def",
					Win:  0,
					Lose: 0,
				},
				{
					Name: "ghi",
					Win:  0,
					Lose: 0,
				},
				{
					Name: "jkl",
					Win:  0,
					Lose: 0,
				},
				{
					Name: "mno",
					Win:  0,
					Lose: 0,
				},
				{
					Name: "pqr",
					Win:  0,
					Lose: 0,
				},
				{
					Name: "stu",
					Win:  0,
					Lose: 0,
				},
				{
					Name: "vwx",
					Win:  0,
					Lose: 0,
				},
			},
			bo:  BO3,
			err: nil,
		},
	}
	for _, tc := range testCases {
		b, err := Generate(tc.ps, tc.bo)
		assert.Equal(suite.T(), tc.err, err, tc.name)
		if b != nil {
			assert.Equal(suite.T(), len(tc.ps)-1, len(b.Matches), tc.name)
			for i, m := range b.Matches {
				if i == 0 {
					assert.Equal(suite.T(), true, m.IsFinal)
				}
				if m.IsFull {
					assert.NotNil(suite.T(), m.ParticipantA)
					assert.NotNil(suite.T(), m.ParticipantB)
				}
				suite.logMatch(m)
			}
		}
	}
}

func (suite *GeneratorTestSuite) logMatch(m *Match) {
	if suite.isLogEnabled {
		log.Default().Printf("match %+v", m)
		if m.ParticipantA != nil {
			log.Default().Printf("participantA %+v", *m.ParticipantA)
		}
		if m.ParticipantB != nil {
			log.Default().Printf("participantB %+v", *m.ParticipantB)
		}
	}
}

func TestGeneratorTestSuite(t *testing.T) {
	suite.Run(t, new(GeneratorTestSuite))
}
