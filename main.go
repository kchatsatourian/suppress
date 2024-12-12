package main

import (
	"github.com/kchatsatourian/suppress/internal/scheduler"
	"github.com/kchatsatourian/suppress/internal/telegram"
)

func main() {
	telegram.Initialize()
	scheduler.Initialize()
}
