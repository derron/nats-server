//go:build wasm
// +build wasm

package server

import (
	"context"
	"fmt"
	"github.com/hack-pad/go-indexeddb/idb"
	"github.com/hack-pad/hackpadfs/indexeddb"
)

var globalFs *indexeddb.FS

func NewFs() *indexeddb.FS {
	if globalFs != nil {
		return globalFs
	}
	fs, err := indexeddb.NewFS(context.Background(), "jetstream", indexeddb.Options{TransactionDurability: idb.DurabilityDefault})
	if err != nil {
		fmt.Printf("NewFs returns fs: %+v, err: %+v\n", fs, err)
	}
	globalFs = fs
	return fs
}

func ensureStorageDirectory(storeDir string) error {
	// wasm use in memory storage skip
	return nil
}

func RemoveAll(name string) error {
	fs := NewFs()
	return fs.Clear(context.Background())
}
