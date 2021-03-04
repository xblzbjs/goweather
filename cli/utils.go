package cli

import (
	"flag"
	"os"
)

func ExitInvalidArguments() {
	// 若不规范的命令行参数，退出
	println("\nUsage: goweather [ -period=current|hourly|daily ] [ -units=C(摄氏度)|F(华氏度) ] <地点>...\n")
	flag.Usage()
	println()
	os.Exit(2)
}