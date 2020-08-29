package commands

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func Pop() error {
	if err := List(); err != nil {
		return err
	}
	fmt.Println("取り出す差分の番号を選択")
	fmt.Print("番号:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringNum := scanner.Text()
	stash := "stash@{" + stringNum + "}"
	_, err := exec.Command("git", "stash", "pop", stash).Output()
	if err != nil {
		return err
	}
	fmt.Println("stashが完了しました")
	return nil
}
