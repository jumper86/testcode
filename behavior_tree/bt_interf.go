package behavior_tree

type BtNodeInterf interface {
	GetId() int64       //获取节点id
	GetTypes() BtnTypes //获取节点类型
	Enable()
	Disable()

	CheckTimer() bool
	Evaluate() bool
	Tick() BtnResult
	Process() BtnResult //节点运行逻辑
	Reset()             //清理自身以及子节点数据

}

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
