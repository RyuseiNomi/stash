package handler

import (
	cm "github.com/RyuseiNomi/stash/src/commands"
)

// StashWorker コマンドを実行するのに必要な情報を保持する構造体
type stashWorker struct {
	args string
}

// NewStashWorker 引数などの情報を初期化した構造体を返却する
func NewStashWorker(a string) *stashWorker {
	return &stashWorker{
		args: a,
	}
}

// Handle the aggregation point of all of process
func (sw *stashWorker) Handle() {
	// ユーザに「save」「pop」「list」「show」のいずれかを選ばせる
	if sw.args == "save" {
		if err := cm.Save(); err != nil {
			panic(err)
		}
	}
	if sw.args == "list" {
		if err := cm.List(); err != nil {
			panic(err)
		}
	}
	if sw.args == "pop" {
		if err := cm.Pop(); err != nil {
			panic(err)
		}
	}
}
