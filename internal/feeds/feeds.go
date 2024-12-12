package feeds

import (
	"log/slog"
	"sync"

	"github.com/kchatsatourian/suppress/internal/state"
	"github.com/kchatsatourian/suppress/internal/telegram"
	Feed "github.com/mmcdole/gofeed"
)

var (
	parser = Feed.NewParser()
)

func Fetch(group *sync.WaitGroup, subscription string, chats []int64) {
	defer group.Done()
	feed, err := parser.ParseURL(subscription)
	if err != nil {
		slog.Warn("Could not fetch RSS feed.", "subscription", subscription, "error", err)
		return
	}

	if feed.Items == nil {
		return
	}

	for _, item := range feed.Items {
		publishedAt := item.PublishedParsed
		if publishedAt == nil {
			publishedAt = item.UpdatedParsed
		}

		if publishedAt.After(state.UpdatedAt) {
			var group sync.WaitGroup
			group.Add(len(chats))
			for _, chat := range chats {
				go telegram.Send(&group, chat, item.Link)
			}
			group.Wait()
		}
	}
}
