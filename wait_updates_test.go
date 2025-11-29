package bot

import (
	"context"
	"sync"
	"testing"

	"github.com/go-telegram/bot/models"
)

func TestWaitUpdates_ProcessAllAfterContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var processed []int64
	b := &Bot{
		processReceivedUpdatesBeforeShutdown: true,
		notAsyncHandlers:                     true,
		updates:                              make(chan *models.Update, 2),
		defaultHandlerFunc: func(ctx context.Context, b *Bot, update *models.Update) {
			processed = append(processed, update.ID)
		},
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go b.waitUpdates(ctx, &wg)

	b.updates <- &models.Update{ID: 1}
	cancel() // simulate shutdown while updates are still coming in
	b.updates <- &models.Update{ID: 2}
	close(b.updates)

	wg.Wait()

	if len(processed) != 2 {
		t.Fatalf("expected 2 updates processed, got %d", len(processed))
	}
	if processed[0] != 1 || processed[1] != 2 {
		t.Fatalf("unexpected updates order: %v", processed)
	}
}

func TestWaitUpdates_StopsWhenChannelClosed(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var processed int
	b := &Bot{
		processReceivedUpdatesBeforeShutdown: true,
		notAsyncHandlers:                     true,
		updates:                              make(chan *models.Update),
		defaultHandlerFunc: func(ctx context.Context, b *Bot, update *models.Update) {
			processed++
		},
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go b.waitUpdates(ctx, &wg)

	close(b.updates)
	wg.Wait()

	if processed != 0 {
		t.Fatalf("expected 0 updates processed after channel close, got %d", processed)
	}
}
