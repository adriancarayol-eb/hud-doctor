package main

import (
	"sync"
	"time"

	spinner2 "github.com/adriancarayol-eb/hud-doctor/pkg/spinner"
)

func step1(wg *sync.WaitGroup, sp *spinner2.MultiStepSpinner) {
	defer wg.Done()
	time.Sleep(6 * time.Second)
	sp.Update("VPN connectivity", "✅ VPN")
}

func step2(wg *sync.WaitGroup, sp *spinner2.MultiStepSpinner) {
	defer wg.Done()
	time.Sleep(4 * time.Second)
	sp.Update("Github connectivity", "✅ Github")
}

func step3(wg *sync.WaitGroup, sp *spinner2.MultiStepSpinner) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	sp.Update("AWS connectivity", "❌: Cannot connect to AWS.")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	spinner := spinner2.NewMultiStepSpinner("yak doctor - Checking connectivity:", 100*time.Millisecond)
	spinner.Update("Github connectivity", "🌐 Checking")
	spinner.Update("VPN connectivity", "🌐 Checking")
	spinner.Update("AWS connectivity", "🌐 Checking")
	spinner.Start()

	go step1(&wg, spinner)
	go step2(&wg, spinner)
	go step3(&wg, spinner)
	wg.Wait()
	spinner.Stop()
}
