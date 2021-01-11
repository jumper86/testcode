package behavior_tree

import (
	"time"
)

type BtNode struct {
	id                int64    //node id，在删除时需要使用
	types             BtnTypes //节点类型
	name              string   //名字
	activated         bool     //是否激活
	interval          int64    //运行cd, 单位纳秒
	lastTimeEvaluated int64    //上次运行时间

	evaluate func() bool //个性验证
	status   BtnResult   //节点运行状态
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

//note: 保证只在第一次执行组合节点的时候，进行一次准入检查，即调用 Evaluate
func (this *BtNode) Evaluate() bool {
	if this.status != Ready {
		return true
	}

	if !(this.activated && this.CheckTimer()) {
		return false
	}
	if this.evaluate != nil && !this.evaluate() {
		return false
	}

	return true
}

func (this *BtNode) Tick() BtnResult {
	return Successed
}

func (this *BtNode) Process() BtnResult {
	if !this.Evaluate() {
		return Failed
	}
	return this.Tick()
}

func (this *BtNode) Reset() {
}
