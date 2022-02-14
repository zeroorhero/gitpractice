package test

import "net/rpc"

func test02() {
	return rpc.HandleHTTP()
}
