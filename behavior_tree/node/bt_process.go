package node

import "testcode/behavior_tree/def"

func Process(btni BtNodeInterf) def.BtnResult {
	if !Evaluate(btni) {
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

//note: 保证只在第一次执行组合节点的时候，进行一次准入检查，即调用 Evaluate
//	Evaluate 函数当目的在于防止不必要的 Tick 调用
func Evaluate(btni BtNodeInterf) bool {
	if !(btni.GetActive() && btni.CheckTimer()) {
		return false
	}

	if btni.GetStatus() == def.Run {
		return true
	}

	if !btni.Evaluate() {
		return false
	}

	return true
}
