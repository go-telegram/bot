package bot

import "github.com/go-telegram/bot/internal/machine"

type states struct {
	def machine.State
	cur machine.State
}

func (s states) isEqual() bool {
	return s.def == s.cur
}

func (s states) isCur(state machine.State) bool {
	return s.cur == state
}
