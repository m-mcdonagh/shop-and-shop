//go:build linux
// +build linux

package files

import (
	"os"
)

var DataHome string

func init() {
	var ok bool
	if DataHome, ok = os.LookupEnv("XDG_DATA_HOME"); ok {
		DataHome += "/shop-and-shop/"
	} else if _, err := os.Stat("~/.local/share"); err == nil {
		DataHome = "~/.local/share/shop-and-shop/"
	} else {
		DataHome = "~/.shop-and-shop/"
	}
	err := os.Mkdir(DataHome, 0755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}
}
