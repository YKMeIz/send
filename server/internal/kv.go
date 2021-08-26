package internal

import "github.com/YKMeIz/KV"

var kvs *KV.NameSpace

func init() {
	kvs = KV.NewNameSpace(KV.Config{
		DiskLess:     true,
		DatabasePath: "",
	})
}
