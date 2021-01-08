package behavior_tree

import "time"

//节点类型
type BtnTypes int

const (
	BaseNode BtnTypes = 0
	Walk     BtnTypes = 1
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

type BtNode struct {
	id                int64    //node id，在删除时需要使用
	types             BtnTypes //节点类型
	name              string   //名字
	activated         bool     //是否激活
	interval          int64    //运行cd, 单位纳秒
	lastTimeEvaluated int64    //上次运行时间

	status BtnResult //节点运行状态
}

//////////////////////////////////////////////
//基本逻辑
//interval 传0，则没有动作冷却时间
func NewBtNode(name string, interval int64) BtNode {
	var btn BtNode
	btn.id = 1
	btn.types = BaseNode
	btn.name = name
	btn.activated = true
	btn.interval = interval
	btn.lastTimeEvaluated = 0
	btn.status = Ready
	return btn
}

//////////////////////////////////////////////
//接口实现

func (this *BtNode) GetId() int64 {
	return this.id
}
func (this *BtNode) GetTypes() BtnTypes {
	return this.types
}

func (this *BtNode) Enable() {
	this.activated = true
}

func (this *BtNode) Disable() {
	this.activated = false
}

func (this *BtNode) CheckTimer() bool {
	if time.Now().UnixNano()-this.lastTimeEvaluated > this.interval {
		return true
	}
	return false
}

func (this *BtNode) Evaluate() bool {
	if this.activated && this.CheckTimer() && this.DoEvaluate() {
		return true
	}
	return false
}

//可能需要实现的部分函数
func (this *BtNode) DoEvaluate() bool {
	return true
}

func (this *BtNode) Tick() BtnResult {
	return Successed
}

func (this *BtNode) Reset() {
}
