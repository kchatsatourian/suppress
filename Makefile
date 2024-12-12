ifneq (,$(wildcard .env))
    include .env
endif

execute:
	@GITHUB_GIST=$(GITHUB_GIST) TELEGRAM_BOT_TOKEN=$(TELEGRAM_BOT_TOKEN) go run main.go

initialize:
	@go mod init github.com/kchatsatourian/suppress

schedule:
	@SCHEDULE=$(SCHEDULE) GITHUB_GIST=$(GITHUB_GIST) TELEGRAM_BOT_TOKEN=$(TELEGRAM_BOT_TOKEN) go run main.go
