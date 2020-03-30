package spinner

import (
	"fmt"
	"github.com/adriancarayol-eb/hud-doctor/internal/hud"
	tm "github.com/buger/goterm"
	"sort"
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
	Steps    map[string]string
}

func NewMultiStepSpinner(prefix, suffix string, d time.Duration) *MultiStepSpinner {
	return &MultiStepSpinner{
		mu:       &sync.RWMutex{},
		renderer: hud.New(),
		stopChan: make(chan struct{}, 1),
		active:   false,
		Delay:    d,
		Suffix:   suffix,
		Prefix:   prefix,
		Steps:    make(map[string]string),
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
				keys := make([]string, 0, len(ms.Steps))
				for k := range ms.Steps {
					keys = append(keys, k)
				}
				sort.Strings(keys)
				totals := tm.NewTable(0, 10, 5, ' ', 0)
				fmt.Fprintf(totals, "Task\tStatus\t\n")
				for _, k := range keys {
					fmt.Fprintf(totals, "%s\t%s\n", k, ms.Steps[k])
				}
				ms.renderer.Refresh(totals)
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
		ms.active = false
		tm.Flush()
	}
}

func (ms *MultiStepSpinner) Update(task, status string) {
	ms.mu.Lock()
	ms.mu.Unlock()
	ms.Steps[task] = status
}