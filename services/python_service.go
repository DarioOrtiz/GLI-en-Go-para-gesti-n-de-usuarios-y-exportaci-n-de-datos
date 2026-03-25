package services

import (
    "fmt"
    "os/exec"
)

func RunPythonAnalysis() {
    cmd := exec.Command("python3", "python/analysis.py")
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("Error ejecutando Python:", err)
        return
    }
    fmt.Println(string(output))
}