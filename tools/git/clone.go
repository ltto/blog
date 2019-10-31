package git

import (
	"errors"
	"os"
	"os/exec"
	"path"
)

func Clone(pathStr, url string) error {
	_ = os.Mkdir(pathStr, 0777)
	init := exec.Command("git", "init")
	init.Dir = pathStr
	err := init.Run()
	if err != nil {
		return err
	}
	add := exec.Command("git", "remote", "add", "origin", url)
	add.Dir = pathStr
	_ = add.Run()
	_, err = os.Open(path.Join(pathStr, ".git"))
	if err != nil {
		return errors.New("git init 失败")
	}
	return nil
}
