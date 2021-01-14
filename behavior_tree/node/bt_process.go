package node

import (
	"testcode/behavior_tree/def"
)

func Process(btni BtNodeInterf) def.BtnResult {

	//fmt.Printf("btni: %v\n", btni.GetTypes())

	//note: 在逻辑节点中Evaluate就会对子节点进行判断
	//	决定执行一个节点之后，在该子节点执行的时候又会对自己进行Evaluate判断
	//	这就构成了重复，因此这里考虑只在逻辑节点中Evaluate对子节点进行检查
	//	子节点自身执行之前不再进行重复检查
	//	因为子节点若是不满足条件，执行也会失败，因此这里不检查也没有关系
	if btni.GetTypes() < def.ActionNodeStartPoint {
		if !Evaluate(btni) {
			//fmt.Printf("btni.evaluate failed: %v\n", btni.GetTypes())

			return def.Failed
		}
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
		//fmt.Printf("time: %v --- btni.evaluate failed: %v\n", time.Now().UnixNano()/1000000, btni.GetTypes())
		//now := time.Now().UnixNano() / 1000000
		//fmt.Printf("now: %d, last: %d\n", now, btni.GetLastTimeTick())

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
