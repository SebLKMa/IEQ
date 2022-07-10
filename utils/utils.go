package utils

import (
	"flag"
	"os"
)

func IsFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func GetFlagName() (bool, string) {
	found := false
	name := ""
	flag.Visit(func(f *flag.Flag) {
		name = f.Name
		found = true
	})
	return found, name
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
