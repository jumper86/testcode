package node

import "testcode/behavior_tree/def"

type BtNodeInterf interface {
	GetId() int64          //获取节点id
	GetTypes() def.BtnType //获取节点类型
	SetTypes(t def.BtnType)
	SetStatus(s def.BtnStatus)
	GetStatus() def.BtnStatus

	GetActive() bool
	Enable()
	Disable()
	CheckTimer() bool

	Evaluate() bool
	Tick() def.BtnResult
	Reset() //清理自身以及子节点数据

}

//
type BtNodeLogicInterf interface {
	BtNodeInterf

	GetChildren() []BtNodeInterf
	AddChild(bn BtNodeInterf)
	RemoveChild(bn BtNodeInterf)
}

type BtNodeDecoratorInterf interface {
	BtNodeInterf

	GetChild() BtNodeInterf
	SetChild(bn BtNodeInterf)
	CleanChild()
}
