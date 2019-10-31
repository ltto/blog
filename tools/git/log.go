package git

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
	"path"
	"strings"
)

type GitLog struct {
	File    os.FileInfo
	Commit  string
	Author  string
	Weekday string
	Month   string
	Day     string
	Time    string
	Year    string
	Remarks string
}

func Log(dir, file string) ([]GitLog, error) {
	cmd := exec.Command("git", "log", file)
	cmd.Dir = dir
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(output)
	scanner := bufio.NewScanner(buffer)
	var list []GitLog
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(text, "commit") && len(text) > 6 {
			glog := GitLog{}
			fileInfo, err := os.Stat(path.Join(dir, file))
			if err != nil {
				return nil, err
			}
			glog.File = fileInfo
			glog.Commit = strings.TrimSpace(strings.TrimLeft(text, "commit"))
			scanner.Scan()
			text = strings.TrimSpace(scanner.Text())
			if strings.HasPrefix(text, "Merge:") {
				scanner.Scan()
				text = strings.TrimSpace(scanner.Text())
			}
			glog.Author = strings.TrimSpace(strings.TrimLeft(text, "Author:"))
			scanner.Scan()
			text = strings.TrimSpace(scanner.Text())
			date := strings.TrimSpace(strings.TrimLeft(text, "Date:"))
			split := strings.Split(date, " ")
			glog.Weekday = split[0]
			glog.Month = split[1]
			glog.Day = split[2]
			glog.Time = split[3]
			glog.Year = split[4]
			scanner.Scan()
			glog.Remarks = strings.TrimSpace(scanner.Text())
			if strings.TrimSpace(glog.Remarks) == "" {
				scanner.Scan()
				glog.Remarks = strings.TrimSpace(scanner.Text())
			}
			list = append(list, glog)
		}
	}
	return list, nil
}
