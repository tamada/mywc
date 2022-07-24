package main

import (
	"fmt"
	"os"

	"github.com/tamada/wildcat"
)

func goMain(args []string) int {
	readOpts := &wildcat.ReadOptions{}
	runtimeOpts := &wildcat.RuntimeOptions{}
	argf := wildcat.NewArgf(args[1:], readOpts, runtimeOpts)
	wildcat := wildcat.NewWildcat(readOpts, runtimeOpts, wildcat.DefaultGenerator)
	rs, err := wildcat.CountAll(argf)
	if !err.IsEmpty() {
		fmt.Println(err.Error())
		return 1
	}
	return printResult(rs)
}

func printResult(rs *wildcat.ResultSet) int {
	printer := wildcat.NewPrinter(os.Stdout, "default", wildcat.BuildSizer(false))
	rs.Print(printer)
	return 0
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
