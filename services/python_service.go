package services

import (
	"fmt"
	"os/exec"
)

func RunPythonAnalysis() {
	cmd := exec.Command("python3", "python_analysis/analysis.py")
	cmd.Stdout = nil
	cmd.Stderr = nil
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error ejecutando Python:", err)
		return
	}
	fmt.Println("Python analysis ejecutado correctamente")
}