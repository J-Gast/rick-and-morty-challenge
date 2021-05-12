package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type Cleaner struct {
	cleaners map[string]func()
}

func (cleaner *Cleaner) Init() {
	cleaner.cleaners = make(map[string]func())
	cleaner.cleaners["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	cleaner.cleaners["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	cleaner.cleaners["darwin"] = cleaner.cleaners["linux"]
}

func (cleaner Cleaner) CallClear() {
	fmt.Println(runtime.GOOS)
	value, ok := cleaner.cleaners[runtime.GOOS]
	if ok {
		value()
	} else {
		fmt.Println("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
