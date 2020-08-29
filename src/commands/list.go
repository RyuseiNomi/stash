package commands

import (
	"fmt"
	"os/exec"
)

// List stashされている内容をリストとして表示する
func List() error {
	fmt.Println("現在Stashされている内容は以下のとおりです。")
	fmt.Println("-------------------------------------")
	out, err := exec.Command("git", "stash", "list").Output()
	if err != nil {
		return err
	}
	fmt.Println(string(out))
	return nil
}
