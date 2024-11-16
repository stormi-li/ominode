package main

import "github.com/stormi-li/ominode"

func main() {
	ominode.StartRedis(6379, "node")
}
