package event

import "context"

type EventBus interface {
	Publish(ctx context.Context, event any) error
}
