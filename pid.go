// Package pid helps to contorl PID file
package pid

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

// File controls pid file
type File struct {
	filepath string
}

// NewFile returns new pointer to PID struct
func NewFile(fp string) *File {
	fp, err := filepath.Abs(fp)
	if err != nil {
		panic(err)
	}
	return &File{
		filepath: fp,
	}
}

// Create creates new pid file
func (p *File) Create() error {
	proc, err := p.Process()
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	if proc != nil {
		return fmt.Errorf("process %d already running", proc.Pid)
	}
	if err := os.MkdirAll(filepath.Dir(p.filepath), 0755); err != nil {
		return err
	}
	return ioutil.WriteFile(p.filepath, []byte(fmt.Sprintf("%d\n", os.Getpid())), 0644)
}

// Remove removes PID file
func (p *File) Remove() error {
	return os.RemoveAll(p.filepath)
}

// Contents returns value stored in pid file
// If file does not exist it returns file does not exist error
func (p *File) Contents() (int, error) {
	contents, err := ioutil.ReadFile(p.filepath)
	if err != nil {
		return 0, err
	}
	pid, err := strconv.Atoi(strings.TrimSpace(string(contents)))
	if err != nil {
		return 0, err
	}
	return pid, nil
}

// Process reads id from pid file and returns process with this pid
// If file does not exist or no process found Process returns nil
func (p *File) Process() (*os.Process, error) {
	pid, err := p.Contents()
	if err != nil {
		return nil, err
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		return nil, err
	}
	err = process.Signal(syscall.Signal(0))
	if err != nil && err.Error() == "no such process" {
		return nil, nil
	}
	if err != nil && err.Error() == "os: process already finished" {
		return nil, nil
	}
	return process, nil
}
