package behavior_tree

//节点类型
type BtnTypes int

const (
	BaseNode BtnTypes = 1

	ComposeSequenceNode         BtnTypes = 11
	ComposeSelectorNode         BtnTypes = 12
	ComposeParallelAndNode      BtnTypes = 13
	ComposeParallelOrNode       BtnTypes = 14
	ComposeParallelFlexibleNode BtnTypes = 15

	Walk BtnTypes = 100
)

//节点运行状态
type BtnResult int

const (
	//None      BtnResult = 0
	Ready     BtnResult = 1
	Running   BtnResult = 2
	Successed BtnResult = 3
	Failed    BtnResult = 4
)
