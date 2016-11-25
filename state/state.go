package state

import (
	"fmt"

	"github.com/acomagu/u-aizu-bot/types"
)

type State struct {
	didVerified bool
}

var states := make(map[types.UserID]State)

func getStates() {}
