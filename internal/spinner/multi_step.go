package spinner

import (
	"fmt"
	"github.com/adriancarayol-eb/hud-doctor/internal/hud"
	"strings"
	"sync"
	"time"
)

type MultiStepSpinner struct {
	mu       *sync.RWMutex
	renderer *hud.Renderer
	active   bool
	stopChan chan struct{}
	Delay    time.Duration
	Suffix   string
	Prefix   string
	Steps    []string
}

func NewMultiStepSpinner(prefix, suffix string, d time.Duration, steps []string) *MultiStepSpinner {
	return &MultiStepSpinner{
		mu:       &sync.RWMutex{},
		renderer: hud.New(),
		stopChan: make(chan struct{}, 1),
		active:   false,
		Delay:    d,
		Suffix:   suffix,
		Prefix:   prefix,
		Steps:    steps,
	}
}

func (ms *MultiStepSpinner) Start() {
	ms.mu.Lock()
	if ms.active {
		ms.mu.Unlock()
		return
	}
	ms.active = true
	ms.mu.Unlock()

	go func() {
		for {
			select {
			case <-ms.stopChan:
				return
			default:
				ms.mu.Lock()
				if !ms.active {
					return
				}
				bodySteps := strings.Join(ms.Steps, "\n")
				bodyPayload := fmt.Sprintf("%s ... %s\n%s", ms.Prefix, ms.Suffix, bodySteps)
				ms.renderer.Refresh(hud.Payload{Body: bodyPayload, TimeStamp: time.Now()})
				ms.mu.Unlock()
				time.Sleep(ms.Delay)
			}
		}
	}()
}

func (ms *MultiStepSpinner) Stop() {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	if ms.active {
		ms.stopChan <- struct{}{}
	}
}

func (ms *MultiStepSpinner) Update(steps []string) {
	ms.mu.Lock()
	ms.Steps = steps
	ms.mu.Unlock()
}
