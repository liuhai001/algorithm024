package main

const MAX_LEVEL = 100

func recursion(level, param int) {
	// recursion terminator 递归终止条件
	if level > MAX_LEVEL {
		// process result
		return
	}

	// process current logic 处理当前逻辑
	process(level, param)

	// drill down 进入下一层
	recursion(level+1, param)

	// reverse the current level status if needed 恢复当前层的状态
}

func process(level, param int) {

}
