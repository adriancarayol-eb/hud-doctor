package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	spinner2 "github.com/adriancarayol-eb/hud-doctor/pkg/spinner"
)

func step1(steps map[string]string) {
	steps["VPN connectivity"] = "✅"
}

func step2(steps map[string]string) {
	steps["Github connectivity"] = "✅"
}

func step3(steps map[string]string) {
	steps["AWS connectivity"] = "❌: Cannot connect to AWS."
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Print(text)
}

func main() {
	steps := make(map[string]string)
	steps["VPN connectivity"] = "🌐 Checking"
	steps["AWS connectivity"] = "🌐 Checking"

	spinner := spinner2.NewMultiStepSpinner("Checking ", " connectivity", 100*time.Millisecond)
	spinner.Update("Github connectivity", "🌐 Checking")
	spinner.Update("VPN connectivity", "🌐 Checking")
	spinner.Update("AWS connectivity", "🌐 Checking")
	spinner.Start()

	time.Sleep(time.Second * 3)
	spinner.Update("Github connectivity", "OK")
	time.Sleep(time.Second * 3)
	spinner.Update("VPN connectivity", "OK")
	time.Sleep(time.Second * 3)
	spinner.Update("AWS connectivity", "FAILED.")
	time.Sleep(time.Second * 3)
	spinner.Stop()
}
