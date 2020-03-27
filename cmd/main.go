package main

import (
	spinner2 "github.com/adriancarayol-eb/hud-doctor/internal/spinner"
	"time"
)

func main() {
	steps := make(map[string]string)
	steps["Github connectivity"] = "🌐 Checking"
	steps["VPN connectivity"] = "🌐 Checking"
	steps["AWS connectivity"] = "🌐 Checking"

	spinner := spinner2.NewMultiStepSpinner("Checking ", " connectivity", 100*time.Millisecond, steps)
	spinner.Start()
	time.Sleep(time.Second * 3)
	steps["VPN connectivity"] = "✅"
	time.Sleep(time.Second * 3)
	steps["Github connectivity"] = "✅"
	time.Sleep(time.Second * 3)
	//spinner.Update(steps)
	time.Sleep(time.Second * 3)
	steps["AWS connectivity"] = "❌: Cannot connect to AWS."
	time.Sleep(time.Second * 3)
	spinner.Stop()
}
