package timeutils

import (
	"fmt"
	"github.com/ezzuodp/go-lib/pkg/log"
	"sort"
	"strings"
	"sync"
	"time"
)

const noStagesText = "no stages"

type Option func(s *Stopwatch)

type Stopwatch struct {
	sync.Mutex
	name   string
	stages map[string]time.Duration
}

func NewStopwatch(name string, options ...Option) *Stopwatch {
	v := &Stopwatch{
		name:   name,
		stages: map[string]time.Duration{},
	}
	for _, option := range options {
		option(v)
	}
	return v
}

type stageDuration struct {
	name string
	d    time.Duration
}

func (s *Stopwatch) stageDurationsSorted() []stageDuration {
	stageDurations := []stageDuration{}
	for n, d := range s.stages {
		stageDurations = append(stageDurations, stageDuration{
			name: n,
			d:    d,
		})
	}
	sort.Slice(stageDurations, func(i, j int) bool {
		return stageDurations[i].d > stageDurations[j].d
	})
	return stageDurations
}

func (s *Stopwatch) sprintStages() string {
	if len(s.stages) == 0 {
		return noStagesText
	}

	stageDurations := s.stageDurationsSorted()

	var stagesStrings []string
	for _, s := range stageDurations {
		stagesStrings = append(stagesStrings, fmt.Sprintf("%s: %s", s.name, s.d))
	}

	return strings.Join(stagesStrings, "\n")
}

func (s *Stopwatch) Reset() {
	s.stages = make(map[string]time.Duration)
}

func (s *Stopwatch) PrintStages() {
	log.Infof("%s stages:\n[%s]", s.name, s.sprintStages())
}

func (s *Stopwatch) TrackStage(name string, f func()) {
	startedAt := time.Now()
	f()
	d := time.Since(startedAt)

	s.Lock()
	s.stages[name] += d
	s.Unlock()
}
