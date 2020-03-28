package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	spinner2 "github.com/adriancarayol-eb/hud-doctor/pkg/spinner"
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
	spinner.Stop()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Print(text)
	time.Sleep(time.Second * 3)
	spinner.Start()
	steps["AWS connectivity"] = "❌: Cannot connect to AWS."
	time.Sleep(time.Second * 3)
	spinner.Stop()
}
