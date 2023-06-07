//go:build linux || darwin

package fs

import (
	"os"
	"syscall"
)

// CloseOnExec makes sure closing the file on process forking.
// 关闭文件
func CloseOnExec(file *os.File) {
	if file != nil {
		syscall.CloseOnExec(int(file.Fd()))
	}
}
