// Package main
package main

import (
	"github.com/Sclock/go-cqhttp/cmd/gocq"
	"github.com/Sclock/go-cqhttp/global/terminal"

	_ "github.com/Sclock/go-cqhttp/db/leveldb"   // leveldb 数据库支持
	_ "github.com/Sclock/go-cqhttp/modules/silk" // silk编码模块
	// 其他模块
	// _ "github.com/Sclock/go-cqhttp/db/sqlite3"   // sqlite3 数据库支持
	// _ "github.com/Sclock/go-cqhttp/db/mongodb"    // mongodb 数据库支持
	// _ "github.com/Sclock/go-cqhttp/modules/pprof" // pprof 性能分析
)

func main() {
	terminal.SetTitle()
	gocq.InitBase()
	gocq.PrepareData()
	gocq.LoginInteract()
	_ = terminal.DisableQuickEdit()
	_ = terminal.EnableVT100()
	gocq.WaitSignal()
	_ = terminal.RestoreInputMode()
}
