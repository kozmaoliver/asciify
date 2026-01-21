package main

import (
	"flag"
	"fmt"
	"github.com/kozmaoliver/asciify/internal/debug"
	"github.com/kozmaoliver/asciify/internal/edge"
	"github.com/kozmaoliver/asciify/internal/frame"
	"github.com/kozmaoliver/asciify/internal/imageio"
	"github.com/kozmaoliver/asciify/internal/luminance"
	"github.com/kozmaoliver/asciify/internal/terminal"
	"github.com/kozmaoliver/asciify/internal/theme"
	"os"
)

func main() {
	debugFlag := flag.Bool("debug", false, "Enable debug mode (saves intermediate images and logs)")
	debugDir := flag.String("debug-dir", "debug_output", "Directory for debug output files")
	edgeCutoff := flag.Float64("edge-cutoff", 95.0, "Edge detection threshold")
	bgColorStr := flag.String("bg", "none", "Background color: none, black, or white")
	colorFlag := flag.Bool("color", false, "Enable colored output using original image colors")
	flag.Parse()

	debug.Init(*debugFlag, *debugDir)

	if flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s [flags] <image-path>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nFlags:\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	imagePath := flag.Arg(0)

	// Get terminal size
	size, err := terminal.GetTerminalSize()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Could not detect terminal size, using defaults: %v\n", err)
	}
	debug.Log("Terminal size: %dx%d", size.Width, size.Height)

	// Load image
	img, err := imageio.LoadImage(imagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading image: %v\n", err)
		os.Exit(1)
	}
	bounds := img.Bounds()
	debug.Log("Loaded image: %dx%d", bounds.Dx(), bounds.Dy())
	debug.SaveImage(img, "01_original")

	// Resize for terminal
	resized := imageio.ResizeForTerminal(img, size.Width, size.Height)
	resizedBounds := resized.Bounds()
	debug.Log("Resized image: %dx%d", resizedBounds.Dx(), resizedBounds.Dy())
	debug.SaveImage(resized, "02_resized")

	// Create frame
	resizedBounds = resized.Bounds()
	frameWidth := resizedBounds.Dx()
	frameHeight := resizedBounds.Dy()
	f := frame.New(frameWidth, frameHeight)

	// Enable color storage if color output is requested
	if *colorFlag {
		f.EnableColors()
	}

	// Step 1: Generate ASCII image based on luminance
	debug.Log("Step 1: Generating ASCII from luminance")
	defaultTheme := theme.NewDefaultTheme()
	chars := defaultTheme.Characters()
	debug.Log("Theme characters: %d levels", len(chars))
	
	for y := 0; y < frameHeight; y++ {
		for x := 0; x < frameWidth; x++ {
			c := resized.At(resizedBounds.Min.X+x, resizedBounds.Min.Y+y)
			
			if *colorFlag {
				f.SetColor(x, y, c)
			}
			
			lum := luminance.Luminance(c)
			
			index := int(lum * float64(len(chars)))
			if index >= len(chars) {
				index = len(chars) - 1
			}
			if index < 0 {
				index = 0
			}
			
			f.Set(x, y, chars[index])
		}
	}
	debug.Log("Generated luminance-based ASCII frame: %dx%d", frameWidth, frameHeight)

	frameCells := make([][]rune, frameHeight)
	for y := 0; y < frameHeight; y++ {
		frameCells[y] = make([]rune, frameWidth)
		for x := 0; x < frameWidth; x++ {
			frameCells[y][x] = f.Get(x, y)
		}
	}
	debug.SaveFrameAsImage(frameCells, "03_luminance_ascii")

	// Step 2: Apply Difference of Gaussians to enhance edges for detection
	debug.Log("Step 2: Applying Difference of Gaussians (sigma1=0.5, sigma2=1.5)")
	dogImage := imageio.DifferenceOfGaussians(resized, 0.5, 1.5)
	debug.SaveImage(dogImage, "04_dog_filtered")

	// Step 3: Detect edges on DoG-filtered image
	debug.Log("Step 3: Detecting edges with Sobel filter")
	edges := edge.Sobel(dogImage)
	
	edgeCount := 0

	// Step 4: Replace edge positions with edge characters
	debug.Log("Step 4: Applying edges with cutoff threshold: %.2f", *edgeCutoff)
	for y := 0; y < frameHeight; y++ {
		for x := 0; x < frameWidth; x++ {
			edgeInfo := edges[y][x]
			if edgeInfo.Strength > *edgeCutoff {
				edgeChar := edge.EdgeChar(edgeInfo.Direction)
				f.Set(x, y, edgeChar)
				edgeCount++
			}
		}
	}
	debug.Log("Applied %d edge characters (%.2f%% of pixels)", edgeCount, float64(edgeCount)*100.0/float64(frameWidth*frameHeight))

	frameCells = make([][]rune, frameHeight)
	for y := 0; y < frameHeight; y++ {
		frameCells[y] = make([]rune, frameWidth)
		for x := 0; x < frameWidth; x++ {
			frameCells[y][x] = f.Get(x, y)
		}
	}
	debug.SaveFrameAsImage(frameCells, "05_final_with_edges")

	// Parse background color
	var bgColor terminal.BackgroundColor
	switch *bgColorStr {
	case "black":
		bgColor = terminal.BgBlack
	case "white":
		bgColor = terminal.BgWhite
	case "none":
		fallthrough
	default:
		bgColor = terminal.BgNone
	}

	// Render to terminal
	debug.Log("Rendering to terminal (bg: %s, color: %v)", *bgColorStr, *colorFlag)
	terminal.RenderFrame(f, bgColor, *colorFlag)
	
	// Save debug logs if debug mode is enabled
	if debug.IsEnabled() {
		debug.WriteLogsToFile("debug.log")
		debug.Log("Debug session complete. Check %s for output files", *debugDir)
	}
}
