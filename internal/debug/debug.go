package debug

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"time"
)

type Debugger struct {
	enabled bool
	outputDir string
	stepCount int
	logs []string
}

var globalDebugger *Debugger

func Init(enabled bool, outputDir string) {
	globalDebugger = &Debugger{
		enabled:   enabled,
		outputDir: outputDir,
		stepCount: 0,
		logs:      make([]string, 0),
	}
	
	if enabled {
		if outputDir != "" {
			os.MkdirAll(outputDir, 0755)
		}
		Log("Debug mode enabled, output directory: %s", outputDir)
	}
}

func IsEnabled() bool {
	return globalDebugger != nil && globalDebugger.enabled
}

func Log(format string, args ...interface{}) {
	if !IsEnabled() {
		return
	}
	msg := fmt.Sprintf(format, args...)
	timestamp := time.Now().Format("15:04:05.000")
	logMsg := fmt.Sprintf("[%s] %s", timestamp, msg)
	globalDebugger.logs = append(globalDebugger.logs, logMsg)
	fmt.Fprintf(os.Stderr, "%s\n", logMsg)
}

func SaveImage(img image.Image, name string) {
	if !IsEnabled() {
		return
	}
	
	globalDebugger.stepCount++
	filename := fmt.Sprintf("step_%02d_%s.png", globalDebugger.stepCount, name)
	
	var fullPath string
	if globalDebugger.outputDir != "" {
		fullPath = filepath.Join(globalDebugger.outputDir, filename)
	} else {
		fullPath = filename
	}
	
	file, err := os.Create(fullPath)
	if err != nil {
		Log("ERROR: Failed to create debug image file %s: %v", fullPath, err)
		return
	}
	defer file.Close()
	
	err = png.Encode(file, img)
	if err != nil {
		Log("ERROR: Failed to encode debug image %s: %v", fullPath, err)
		return
	}
	
	Log("Saved debug image: %s", fullPath)
}

func SaveFrameAsImage(frame [][]rune, name string) {
	if !IsEnabled() {
		return
	}
	
	height := len(frame)
	if height == 0 {
		return
	}
	width := len(frame[0])
	
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			ch := frame[y][x]
			var gray uint8
			if ch == 0 {
				gray = 0
			} else {
				gray = uint8(int(ch) % 256)
			}
			img.SetRGBA(x, y, color.RGBA{R: gray, G: gray, B: gray, A: 255})
		}
	}
	
	globalDebugger.stepCount++
	filename := fmt.Sprintf("step_%02d_%s.png", globalDebugger.stepCount, name)
	
	var fullPath string
	if globalDebugger.outputDir != "" {
		fullPath = filepath.Join(globalDebugger.outputDir, filename)
	} else {
		fullPath = filename
	}
	
	file, err := os.Create(fullPath)
	if err != nil {
		Log("ERROR: Failed to create debug frame file %s: %v", fullPath, err)
		return
	}
	defer file.Close()
	
	err = png.Encode(file, img)
	if err != nil {
		Log("ERROR: Failed to encode debug frame %s: %v", fullPath, err)
		return
	}
	
	Log("Saved debug frame: %s", fullPath)
}

func GetLogs() []string {
	if !IsEnabled() {
		return nil
	}
	return globalDebugger.logs
}

func WriteLogsToFile(filename string) {
	if !IsEnabled() {
		return
	}
	
	var fullPath string
	if globalDebugger.outputDir != "" {
		fullPath = filepath.Join(globalDebugger.outputDir, filename)
	} else {
		fullPath = filename
	}
	
	file, err := os.Create(fullPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create log file: %v\n", err)
		return
	}
	defer file.Close()
	
	for _, log := range globalDebugger.logs {
		fmt.Fprintf(file, "%s\n", log)
	}
	
	Log("Saved debug logs to: %s", fullPath)
}
