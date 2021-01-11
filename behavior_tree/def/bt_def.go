package def

//节点类型
type BtnType int

const (
	BaseNode BtnType = 1

	//逻辑节点
	ComposeSequenceNode         BtnType = 11
	ComposeSelectorNode         BtnType = 12
	ComposeParallelAndNode      BtnType = 13
	ComposeParallelOrNode       BtnType = 14
	ComposeParallelFlexibleNode BtnType = 15

	//装饰器节点
	DecoratorSucceederNode        BtnType = 51
	DecoratorInverterNode         BtnType = 52
	DecoratorRepeatUtilFailedNode BtnType = 53

	//行为节点
	Walk BtnType = 100
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
