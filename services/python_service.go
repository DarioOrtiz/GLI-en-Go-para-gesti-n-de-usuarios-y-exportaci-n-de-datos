package services

import (
	"fmt"
	"os/exec"
)

func RunPythonAnalysis() {
	 fmt.Println("Ejecutando análisis Python...")
	cmd := exec.Command("python3", "python/process_data.py")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error ejecutando Python:", err)
	} else {
		fmt.Println(string(output))
	}
}