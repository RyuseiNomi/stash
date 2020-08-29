package commands

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

// Show stashの差分内容を表示するコマンドを実行する
func Show() error {
	fmt.Println("表示したい差分の番号を選択してください")
	if err := List(); err != nil {
		return err
	}
	fmt.Println("表示したい差分の番号を選択")
	fmt.Print("番号:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringNum := scanner.Text()
	stash := "stash@{" + stringNum + "}"
	out, err := exec.Command("git", "stash", "show", stash).Output()
	if err != nil {
		return err
	}
	fmt.Println("--------------------------------------------------")
	fmt.Println(string(out))
	fmt.Println("--------------------------------------------------")
	return nil
}
