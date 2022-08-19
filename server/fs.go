//go:build !wasm
// +build !wasm

package server

import (
	"fmt"
	"github.com/hack-pad/hackpadfs/os"
	"io/ioutil"
	goOS "os"
)

func NewFs() *os.FS {
	return os.NewFS()
}

func RemoveAll(name string) error {
	return os.NewFS().RemoveAll(name)
}

func ensureStorageDirectory(storeDir string) error {
	if stat, err := goOS.Stat(storeDir); goOS.IsNotExist(err) {
		if err := goOS.MkdirAll(storeDir, defaultDirPerms); err != nil {
			return fmt.Errorf("could not create storage directory - %v", err)
		}
	} else {
		// Make sure its a directory and that we can write to it.
		if stat == nil || !stat.IsDir() {
			return fmt.Errorf("storage directory is not a directory")
		}
		tmpfile, err := ioutil.TempFile(storeDir, "_test_")
		if err != nil {
			return fmt.Errorf("storage directory is not writable")
		}
		defer func() {
			tmpfile.Close()
			goOS.Remove(tmpfile.Name())
		}()
	}
	return nil
}
