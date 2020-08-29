package commands

import (
	"fmt"
	"bufio"
	"os"
	"os/exec"
	"regexp"
)

// Save gitの差分をstashする
func Save() error {
	fmt.Println("以下の差分をStashします。")	
	fmt.Println("------------------------------")	
	out, err := exec.Command("git", "status").Output()
	if err != nil {
		return err
	}
	r := regexp.MustCompile(`modified:  .*`)
	files := r.FindAllStringSubmatch(string(out), -1)
	for i := range files {
		fmt.Printf("\x1b[31m%s\x1b[0m \n", files[i])
	}
	fmt.Println("\n-----------------------")
	fmt.Println("stashに対する説明文を入力してください。")
	fmt.Print("説明:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	out, err = exec.Command("git", "stash", "save", text).Output()
	if err != nil {
		panic(err)
	}
	fmt.Print("stashが完了しました。")
	return nil
}