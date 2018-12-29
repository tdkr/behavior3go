package loader

import (
	_ "fmt"
	_ "reflect"

	. "github.com/tdkr/behavior3go/actions"
	. "github.com/tdkr/behavior3go/composites"
	. "github.com/tdkr/behavior3go/config"
	. "github.com/tdkr/behavior3go/core"
	. "github.com/tdkr/behavior3go/decorators"
)

func createBaseStructMaps() *RegisterStructMaps {
	st := NewRegisterStructMaps()
	//actions
	st.Register("Error", &Error{})
	st.Register("Failer", &Failer{})
	st.Register("Runner", &Runner{})
	st.Register("Succeeder", &Succeeder{})
	st.Register("Wait", &Wait{})
	st.Register("Log", &Log{})
	//composites
	st.Register("MemPriority", &MemPriority{})
	st.Register("MemSequence", &MemSequence{})
	st.Register("Priority", &Priority{})
	st.Register("Sequence", &Sequence{})

	//decorators
	st.Register("Inverter", &Inverter{})
	st.Register("Limiter", &Limiter{})
	st.Register("MaxTime", &MaxTime{})
	st.Register("Repeater", &Repeater{})
	st.Register("RepeatUntilFailure", &RepeatUntilFailure{})
	st.Register("RepeatUntilSuccess", &RepeatUntilSuccess{})
	return st
}

func CreateBevTreeFromConfig(config *BTTreeCfg, extMap *RegisterStructMaps) *BehaviorTree {
	tree := NewBeTree()
	tree.Load(config, extMap)
	return tree
}
