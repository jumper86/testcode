package node

import (
	"math/rand"
	"testcode/behavior_tree/def"
	"time"
)

type BtNode struct {
	id           int64       //node id，在删除时需要使用
	types        def.BtnType //节点类型
	name         string      //名字
	activated    bool        //是否激活
	interval     int64       //运行cd, 单位纳秒
	lastTimeTick int64       //上次运行时间

	evaluate func() bool   //个性验证
	status   def.BtnStatus //节点运行状态
}

//////////////////////////////////////////////
//基本逻辑
//interval 传0，则没有动作冷却时间
func NewBtNode(name string, interval int64) BtNode {
	var btn BtNode
	btn.id = rand.Int63n(10000)
	btn.types = def.BaseNode
	btn.name = name
	btn.activated = true
	btn.interval = interval
	btn.lastTimeTick = 0
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

func (this *BtNode) GetActive() bool {
	return this.activated
}

func (this *BtNode) UpdateLastTimeTick() {
	t := time.Now().UnixNano() / 1000000
	this.lastTimeTick = t
}

func (this *BtNode) GetLastTimeTick() int64 {
	return this.lastTimeTick
}

func (this *BtNode) CheckTimer() bool {
	now := time.Now().UnixNano() / 1000000
	if now-this.lastTimeTick > this.interval {
		//fmt.Printf("time: %v --- btni.evaluate sucess: %v\n", time.Now().UnixNano()/1000000, this.GetTypes())
		//fmt.Printf("now: %d, last: %d, interval: %d\n", now, this.lastTimeTick, this.interval)

		//this.UpdateLastTimeTick()

		return true
	}
	return false
}

func (this *BtNode) Tick() def.BtnResult {
	return def.Successed
}

func (this *BtNode) Reset() {
}

func (this *BtNode) Evaluate() bool {
	return true
}
