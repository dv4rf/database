package compute

import (
	"errors"
	"strings"
)

const (
	foundLetterEvent = iota
	foundWhiteSpaceEvent
	// must be last
	eventsNumber
)

const (
	initialState = iota
	wordState
	whiteSpaceState
	invalidState
	// must be last
	statesNumber
)

type transition struct {
	jump   func(byte) int
	action func()
}

type stateMachine struct {
	transitions [statesNumber][eventsNumber]transition
	state       int

	tokens []string
	sb     strings.Builder
}

func newStateMachine() *stateMachine {
	machine := &stateMachine{
		state: initialState,
	}

	machine.transitions = [statesNumber][eventsNumber]transition{
		initialState: {
			foundLetterEvent:     transition{jump: machine.appendLetterJump},
			foundWhiteSpaceEvent: transition{jump: machine.skipWhiteSpaceJump},
		},
		wordState: {
			foundLetterEvent:     transition{jump: machine.appendLetterJump},
			foundWhiteSpaceEvent: transition{jump: machine.skipWhiteSpaceJump, action: machine.addTokenAction},
		},
		whiteSpaceState: {
			foundLetterEvent:     transition{jump: machine.appendLetterJump},
			foundWhiteSpaceEvent: transition{jump: machine.skipWhiteSpaceJump},
		},
		invalidState: {},
	}

	return machine
}

func (sm *stateMachine) parse(input string) ([]string, error) {
	for i := 0; i < len(input); i++ {
		symbol := input[i]
		if isWhiteSpace(symbol) {
			sm.processEvent(foundWhiteSpaceEvent, symbol)
		} else if isLetter(symbol) {
			sm.processEvent(foundLetterEvent, symbol)
		} else {
			return nil, errors.New("invalid symbol")
		}
	}

	sm.processEvent(foundWhiteSpaceEvent, ' ')
	return sm.tokens, nil
}

func (sm *stateMachine) processEvent(event int, symbol byte) {
	transition := sm.transitions[sm.state][event]
	sm.state = transition.jump(symbol)
	if transition.action != nil {
		transition.action()
	}
}

func (sm *stateMachine) appendLetterJump(letter byte) int {
	sm.sb.WriteByte(letter)
	return wordState
}

func (sm *stateMachine) skipWhiteSpaceJump(byte) int {
	return whiteSpaceState
}

func (sm *stateMachine) addTokenAction() {
	sm.tokens = append(sm.tokens, sm.sb.String())
	sm.sb.Reset()
}

func isWhiteSpace(symbol byte) bool {
	return symbol == '\t' || symbol == '\n' || symbol == ' '
}

func isLetter(symbol byte) bool {
	return (symbol >= 'a' && symbol <= 'z') ||
		(symbol >= 'A' && symbol <= 'Z') ||
		(symbol >= '0' && symbol <= '9') ||
		(symbol == '*') || (symbol == '/') || (symbol == '_')
}
