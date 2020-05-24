package main

import (
	"fmt"
	"io"
)

func new(w io.Writer) kVS {
	return kVS{
		storage: w,
		entries: map[string]string{},
	}
}

type kVS struct {
	storage io.Writer
	entries map[string]string
}

func (k *kVS) add(key, value string) error {
	// storeValue :=
	// written, err :=

	k.storage.Write([]byte(fmt.Sprintf("%s:%s\n", key, value)))
	k.entries[key] = value

	return nil
}
