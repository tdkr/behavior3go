package loader

import (
	"sync"

	. "github.com/tdkr/behavior3go/actions"
	. "github.com/tdkr/behavior3go/composites"
	"github.com/tdkr/behavior3go/config"
	. "github.com/tdkr/behavior3go/core"
	. "github.com/tdkr/behavior3go/decorators"
)

var instance *treeManager
var once sync.Once

func TreeManager() *treeManager {
	once.Do(func() {
		instance = &treeManager{
			nodeMap: map[string]nodeCreator{},
			treeMap: make(map[string]*config.BTTreeCfg),
		}
		// Actions
		instance.RegisterNode("Error", func() IBaseNode {
			return &Error{}
		})
		instance.RegisterNode("Failer", func() IBaseNode {
			return &Failer{}
		})
		instance.RegisterNode("Runner", func() IBaseNode {
			return &Runner{}
		})
		instance.RegisterNode("Succeeder", func() IBaseNode {
			return &Succeeder{}
		})
		instance.RegisterNode("Wait", func() IBaseNode {
			return &Wait{}
		})
		//composites
		instance.RegisterNode("MemPriority", func() IBaseNode {
			return &MemPriority{}
		})
		instance.RegisterNode("MemSequence", func() IBaseNode {
			return &MemSequence{}
		})
		instance.RegisterNode("Priority", func() IBaseNode {
			return &Priority{}
		})
		instance.RegisterNode("Sequence", func() IBaseNode {
			return &Sequence{}
		})
		//decorators
		instance.RegisterNode("Inverter", func() IBaseNode {
			return &Inverter{}
		})
		instance.RegisterNode("Limiter", func() IBaseNode {
			return &Limiter{}
		})
		instance.RegisterNode("MaxTime", func() IBaseNode {
			return &MaxTime{}
		})
		instance.RegisterNode("Repeater", func() IBaseNode {
			return &Repeater{}
		})
		instance.RegisterNode("RepeatUntilFailure", func() IBaseNode {
			return &RepeatUntilFailure{}
		})
		instance.RegisterNode("RepeatUntilSuccess", func() IBaseNode {
			return &RepeatUntilSuccess{}
		})
	})
	return instance
}

type nodeCreator = func() IBaseNode

type treeManager struct {
	treeMap map[string]*config.BTTreeCfg
	nodeMap map[string]nodeCreator
}

func (mgr *treeManager) CreateNode(id string) IBaseNode {
	if f, ok := mgr.nodeMap[id]; ok {
		return f()
	}
	return nil
}

func (mgr *treeManager) RegisterNode(id string, creator nodeCreator) {
	mgr.nodeMap[id] = creator
}

func (mgr *treeManager) SetTreeCfg(name string, cfg *config.BTTreeCfg) {
	mgr.treeMap[name] = cfg
}

func (mgr *treeManager) GetTreeConfig(name string) *config.BTTreeCfg {
	return mgr.treeMap[name]
}

func (mgr *treeManager) SetTreeMap(treeMap map[string]*config.BTTreeCfg) {
	mgr.treeMap = treeMap
}
