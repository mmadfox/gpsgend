package broker

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/mmadfox/go-gpsgen"
	gpsgendproto "github.com/mmadfox/gpsgend/gen/proto/gpsgend/v1"
	"github.com/mmadfox/gpsgend/internal/types"
)

const Cap = 64

type Broker struct {
	out     chan *gpsgendproto.Event
	hub     *hub
	wg      sync.WaitGroup
	closeCh chan struct{}
	once    sync.Once
}

func New() *Broker {
	b := Broker{
		out:     make(chan *gpsgendproto.Event, Cap),
		hub:     newHub(),
		closeCh: make(chan struct{}),
	}
	return &b
}

func (b *Broker) RegisterClient(id uuid.UUID, c Client) {
	b.hub.Register(id, c)
}

func (b *Broker) Unregister(id uuid.UUID) error {
	return b.hub.Unregister(id)
}

func (b *Broker) PublishTrackerCreated(ctx context.Context, trackerID types.ID) {
	event := makeTrackerCreatedEvent(trackerID)
	b.send(event)
}

func (b *Broker) PublishTrackerRemoved(ctx context.Context, trackerID types.ID) {
	event := makeTrackerRemovedEvent(trackerID)
	b.send(event)
}

func (b *Broker) PublishTrackerUpdated(ctx context.Context, trackerUD types.ID) {
	event := makeTrackerUpdatedEvent(trackerUD)
	b.send(event)
}

func (b *Broker) PublishTrackerStarted(ctx context.Context, trackerID types.ID) {
	event := makeTrackerStartedEvent(trackerID)
	b.send(event)
}

func (b *Broker) PublishTrackerStopped(ctx context.Context, trackerID types.ID) {
	event := makeTrackerStoppedEvent(trackerID)
	b.send(event)
}

func (b *Broker) PublishTrackerRoutesAdded(
	ctx context.Context,
	trackerID types.ID,
	routes []*gpsgen.Route,
) {
	event := makeTrackerRoutesAddedEvent(trackerID, routes)
	b.send(event)
}

func (b *Broker) PublishTrackerRouteRemoved(ctx context.Context, trackerID, routeID types.ID) {
	event := makeTrackerRoutesRemoveEvent(trackerID, routeID)
	b.send(event)
}

func (b *Broker) PublishTrackerRoutesReseted(ctx context.Context, trackerID types.ID) {
	event := makeTrackerRoutesResetedEvent(trackerID)
	b.send(event)
}

func (b *Broker) PublishTrackerNavigatorReseted(ctx context.Context, trackerID types.ID) {
	event := makeTrackerNavigatorResetedEvent(trackerID)
	b.send(event)
}

func (b *Broker) PublishTrackerNavigatorJumped(ctx context.Context, trackerID types.ID) {
	evnet := makeTrackerNavigatorJumpedEvent(trackerID)
	b.send(evnet)
}

func (b *Broker) PublishTrackerSensorAdded(ctx context.Context, trackerID, sensorID types.ID) {
	evnet := makeTrackerSensorAddedEvent(trackerID, sensorID)
	b.send(evnet)
}

func (b *Broker) PublishTrackerSensorRemoved(ctx context.Context, trackerID, sensorID types.ID) {
	event := makeTrackerSensorRemovedEvent(trackerID, sensorID)
	b.send(event)
}

func (b *Broker) PublishTrackerShutdowned(ctx context.Context, trackerID types.ID) {
	event := makeTrackerShutdownedEvent(trackerID)
	b.send(event)
}

func (b *Broker) PublishTrackerResumed(ctx context.Context, trackerID types.ID) {
	event := makeTrackerResumedEvent(trackerID)
	b.send(event)
}

func (b *Broker) Run() {
	b.wg.Add(1)
	go b.doWork()
	b.wg.Wait()
}

func (b *Broker) Close() {
	b.hub.Close()

	b.once.Do(func() {
		close(b.closeCh)
	})
}

func (b *Broker) doWork() {
	defer b.wg.Done()

	for {
		select {
		case <-b.closeCh:
			return
		case event := <-b.out:
			b.hub.Broadcast(event)
		}
	}
}

func (b *Broker) send(event *gpsgendproto.Event) {
	b.out <- event
}
