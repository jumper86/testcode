package behavior_tree

type BtNodeInterf interface {
	GetId() int64       //获取节点id
	GetTypes() BtnTypes //获取节点类型
	Enable()
	Disable()
	CheckTimer() bool
	Evaluate() bool
	DoEvaluate() bool //每个节点自身情况判断准入
	Tick() BtnResult  //节点运行逻辑
	Reset()           //清理自身以及子节点数据
}

type BtNodeComposeInterf interface {
	BtNodeInterf
	AddChild(bn BtNodeInterf)
	RemoveChild(bn BtNodeInterf)
}
