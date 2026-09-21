package main

import (
	"github.com/kchatsatourian/suppress/internal/scheduler"
	"github.com/kchatsatourian/suppress/internal/state"
	"github.com/kchatsatourian/suppress/internal/telegram"
)

func main() {
	telegram.Initialize()
	state.Initialize()
	scheduler.Initialize()
}
