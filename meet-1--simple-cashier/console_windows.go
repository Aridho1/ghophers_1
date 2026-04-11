//go:build windows

package main

import (
	"os"

	"golang.org/x/sys/windows"
)

func init() {
	// Aktifkan ANSI/VT di cmd.exe & PowerShell klasik (sama ide dengan yang dilakukan CPython di Windows).
	for _, fd := range []uintptr{os.Stdout.Fd(), os.Stderr.Fd()} {
		h := windows.Handle(fd)
		var mode uint32
		if err := windows.GetConsoleMode(h, &mode); err != nil {
			continue
		}
		_ = windows.SetConsoleMode(h, mode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
	}
}
