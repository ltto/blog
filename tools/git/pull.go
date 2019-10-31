package git

import (
	"blog/system"
	"os"
	"os/exec"
	"path"
)

func Pull(dir string) error {
	_, err := os.Open(path.Join(dir, ".git"))
	if err != nil {
		err = Clone(dir, system.Conf.Blog.Git)
		if err != nil {
			return err
		}
	}
	cmd := exec.Command("git", "pull", "origin", "master")
	cmd.Dir = dir
	return cmd.Run()
}
