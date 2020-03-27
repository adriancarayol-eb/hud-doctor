package main

import (
	spinner2 "github.com/adriancarayol-eb/hud-doctor/internal/spinner"
	"time"
)

func main() {
	steps := make([]string, 3)
	steps[0] = "🌐 Checking GitHub connectivity..."
	spinner := spinner2.NewMultiStepSpinner("Checking ", " connectivity", 100*time.Millisecond, steps)
	spinner.Start()
	time.Sleep(time.Second * 3)
	steps[1] = "🌐 Checking VPN connectivity..."
	time.Sleep(time.Second * 3)
	steps[2] = "🌐 Checking AWS connectivity..."
	time.Sleep(time.Second * 3)
	//spinner.Update(steps)
	time.Sleep(time.Second * 3)
	steps[2] = "❌  Error checking AWS connectivity..."
	time.Sleep(time.Second * 3)
	spinner.Stop()
}
