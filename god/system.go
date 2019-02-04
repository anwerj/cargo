package god

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

// System in the container of all system dependencies
type System struct{}

// NewSystem returns new system
func NewSystem() System {
	return System{}
}

// PathExists check if path exists
func (s *System) PathExists(p string) bool {
	if _, err := os.Stat(p); err == nil {
		return true
	}
	return false
}

// AbsPath returns absolute path
func (s *System) AbsPath(p string) string {
	fp, err := filepath.Abs(p)
	if err == nil {
		return fp
	}
	return ""
}

// Exec will excute a command silently
func (s *System) Exec(c string, a ...string) string {
	out, err := exec.Command(c, a...).Output()
	if err != nil {
		return fmt.Sprint(err)
	}
	return fmt.Sprintf("%s", out)
}

// ReadDir lists the directory in non-recursive fashion
func (s *System) ReadDir(p string, a ...interface{}) (dirs []string) {

	filter := func(s string) bool { return true }
	if len(a) > 0 {
		if f, ok := a[0].(func(s string) bool); ok {
			filter = f
		}
	}

	l, err := ioutil.ReadDir(p)
	if err == nil {
		for _, f := range l {
			//fmt.Println(f.Mode(), f.Name())
			if f.IsDir() && f.Name()[0:1] != "." && filter(f.Name()) {
				dirs = append(dirs, f.Name())
			}
		}
	}
	return
}

// ReadFiles lists files in recursive manner
func (s *System) ReadFiles(p string, a ...interface{}) (files []string) {

	filter := func(s string) bool { return true }
	if len(a) > 0 {
		if f, ok := a[0].(func(s string) bool); ok {
			filter = f
		}
	}

	err := filepath.Walk(p, func(path string, info os.FileInfo, err error) error {

		if !info.IsDir() && info.Name()[0:1] != "." && filter(info.Name()) {
			files = append(files, path)
		}
		return err
	})

	if err == nil {

	}

	return
}
