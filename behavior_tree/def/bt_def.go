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
	ComposeAnchorLoopNode       BtnType = 16

	//装饰器节点
	DecoratorIfNode                  BtnType = 51
	DecoratorSucceederNode           BtnType = 52
	DecoratorInverterNode            BtnType = 53
	DecoratorRepeatUtilFailedNode    BtnType = 54
	DecoratorRepeatUtilSeccessedNode BtnType = 55
	//行为节点
	Action1 BtnType = 100
	Action2 BtnType = 100
)

//节点运行结果
type BtnResult int

const (
	None      BtnResult = 0
	Running   BtnResult = 2
	Successed BtnResult = 3
	Failed    BtnResult = 4
)

//节点当前状态
type BtnStatus int

const (
	Ready BtnStatus = 1
	Run   BtnStatus = 2
)
