package node

import (
	"testcode/behavior_tree/def"
	"time"
)

func Process(btni BtNodeInterf) def.BtnResult {
	if !btni.Evaluate() {
		return def.Failed
	}
	if btni.GetStatus() != def.Run {
		btni.SetStatus(def.Run)
	}

	tmpRst := btni.Tick()
	if tmpRst != def.Running {
		btni.Reset()
	}
	return tmpRst
}

type BtNode struct {
	id                int64       //node id，在删除时需要使用
	types             def.BtnType //节点类型
	name              string      //名字
	activated         bool        //是否激活
	interval          int64       //运行cd, 单位纳秒
	lastTimeEvaluated int64       //上次运行时间

	evaluate func() bool   //个性验证
	status   def.BtnStatus //节点运行状态
}

//////////////////////////////////////////////
//基本逻辑
//interval 传0，则没有动作冷却时间
func NewBtNode(name string, interval int64) BtNode {
	var btn BtNode
	btn.id = 1
	btn.types = def.BaseNode
	btn.name = name
	btn.activated = true
	btn.interval = interval
	btn.lastTimeEvaluated = 0
	btn.status = def.Ready
	return btn
}

//////////////////////////////////////////////
//接口实现

func (this *BtNode) GetId() int64 {
	return this.id
}
func (this *BtNode) GetTypes() def.BtnType {
	return this.types
}

func (this *BtNode) SetTypes(t def.BtnType) {
	this.types = t
	return
}

func (this *BtNode) SetEvaluate(f func() bool) {
	this.evaluate = f
	return
}

func (this *BtNode) SetStatus(s def.BtnStatus) {
	this.status = s
}

func (this *BtNode) GetStatus() def.BtnStatus {
	return this.status
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
//	Evaluate 函数当目的在于防止不必要的 Tick 调用
func (this *BtNode) Evaluate() bool {
	if this.status == def.Run {
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

func (this *BtNode) Tick() def.BtnResult {
	return def.Successed
}

func (this *BtNode) Reset() {
}
