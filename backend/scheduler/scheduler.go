// found information how to do this on https://articles.wesionary.team/building-basic-event-scheduler-in-go-134c19f77f84
package scheduler

import (
	"context"
	"database/sql"
	//this is what the tutorial recommended, will have to change with our own db
	//"database/sql"
	"log"
	"time"

	//database
	"github.com/robfig/cron/v3"
)

// Scheduler data structure
type Scheduler struct {
	db          *sql.DB
	listeners   Listeners
	cron        *cron.Cron
	cronEntries map[string]cron.EntryID
}

// Listeners has attached event listeners
type Listeners map[string]ListenFunc

// ListenFunc function that listens to events
type ListenFunc func(string)

// Event structure
type Event struct {
	ID      uint
	Name    string
	Payload string
	Cron    string
}

// NewScheduler creates a new scheduler
func NewScheduler(db *sql.DB, listeners Listeners) Scheduler {

	return Scheduler{
		db:          db,
		listeners:   listeners,
		cron:        cron.New(),
		cronEntries: map[string]cron.EntryID{},
	}

}

func (s Scheduler) AddListener(event string, listenFunc ListenFunc) {
	s.listeners[event] = listenFunc
}

func (s Scheduler) CheckEventsInInterval(ctx context.Context, duration time.Duration) {
	ticker := time.NewTicker(duration)
	go func() {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				log.Println("⏰ Ticks Received...")
				events := s.checkDueEvents()
				for _, e := range events {
					s.callListeners(e)
				}
			}

		}
	}()
}

// checkDueEvents checks and returns due events
func (s Scheduler) checkDueEvents() []Event {
	events := []Event{}
	rows, err := s.db.Query(`SELECT "id", "name", "payload" FROM "public"."jobs" WHERE "runAt" < $1 AND "cron"='-'`, time.Now())
	if err != nil {
		log.Print("💀 error: ", err)
		return nil
	}
	for rows.Next() {
		evt := Event{}
		rows.Scan(&evt.ID, &evt.Name, &evt.Payload)
		events = append(events, evt)
	}
	return events
}

// callListeners calls the event listener of provided event
func (s Scheduler) callListeners(event Event) {
	eventFn, ok := s.listeners[event.Name]
	if ok {
		go eventFn(event.Payload)
		_, err := s.db.Exec(`DELETE FROM "public"."jobs" WHERE "id" = $1`, event.ID)
		if err != nil {
			log.Print("💀 error: ", err)
		}
	} else {
		log.Print("💀 error: couldn't find event listeners attached to ", event.Name)
	}

}

// CheckEventsInInterval checks the event in given interval
func (s Scheduler) CheckEventsInInterval(ctx context.Context, duration time.Duration) {
	ticker := time.NewTicker(duration)
	go func() {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				log.Println("⏰ Ticks Received...")
				events := s.checkDueEvents()
				for _, e := range events {
					s.callListeners(e)
				}
			}

		}
	}()
}
