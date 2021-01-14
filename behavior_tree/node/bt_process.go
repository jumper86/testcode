package node

import (
	"testcode/behavior_tree/def"
)

func Process(btni BtNodeInterf) def.BtnResult {

	//fmt.Printf("btni: %v\n", btni.GetTypes())

	if !Evaluate(btni) {
		//fmt.Printf("btni.evaluate failed: %v\n", btni.GetTypes())

		return def.Failed
	}
	if btni.GetStatus() != def.Run {
		btni.SetStatus(def.Run)
	}

	//更新最近执行时间
	btni.UpdateLastTimeTick()

	//fmt.Printf("btni.tick: %v\n", btni.GetTypes())
	tmpRst := btni.Tick()
	//fmt.Printf("btni: %v, --------tmpRst : %v\n", btni.GetTypes(), tmpRst)

	if tmpRst != def.Running {
		//fmt.Printf("btni: %v, --------reset : %v\n", btni.GetTypes(), time.Now().UnixNano())

		btni.Reset()
	}
	return tmpRst
}

//note: 保证只在第一次执行组合节点的时候，进行一次准入检查，即调用 Evaluate
//	Evaluate 函数当目的在于防止不必要的 Tick 调用
func Evaluate(btni BtNodeInterf) bool {
	//fmt.Printf("btni.evaluate : %v\n", btni.GetTypes())

	if !(btni.GetActive() && btni.CheckTimer()) {
		//fmt.Printf("--- btni.evaluate failed: %v\n", btni.GetTypes())

		return false
	}
	//fmt.Printf("time: %v\n", time.Now().UnixNano())

	if btni.GetStatus() == def.Run {
		return true
	}

	//fmt.Printf("======== time: %v\n", time.Now().UnixNano())

	if !btni.Evaluate() {
		return false
	}
	//fmt.Printf("btni.evaluate pass: %v\n", btni.GetTypes())

	return true
}
