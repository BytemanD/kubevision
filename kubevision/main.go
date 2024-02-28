package main

import (
	_ "kubevision/internal/packed"

	"kubevision/internal/cmd"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
