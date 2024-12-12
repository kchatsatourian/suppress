package suppress

import (
	"log/slog"
	"sync"
	"time"

	"github.com/kchatsatourian/suppress/internal/feeds"
	"github.com/kchatsatourian/suppress/internal/state"
	"github.com/kchatsatourian/suppress/internal/subscriptions"
)

func Execute() {
	slog.Debug("Executing...")
	state.ExecutedAt = time.Now()
	state.Read()
	subscriptions.Read()
	var group sync.WaitGroup
	group.Add(len(subscriptions.Subscriptions))
	for endpoint, subscription := range subscriptions.Subscriptions {
		go feeds.Fetch(&group, endpoint, subscription.Channels)
	}
	group.Wait()
	state.Write()
}
