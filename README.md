# ASCIIfy
Terminal ASCII Image Renderer

A fast, terminal-based ASCII art generator written in Go that converts images into beautiful ASCII representations using advanced luminance mapping and edge detection.

## Features

- **Smart Terminal Fitting**: Automatically detects terminal dimensions and scales images to fit perfectly
- **Advanced Edge Detection**: Uses Sobel filters to enhance image clarity with directional characters
- **Aspect Ratio Preservation**: Accounts for terminal character proportions to display images correctly
- **Multiple Format Support**: PNG, JPEG, and GIF (single frame)
- **Modular Architecture**: Clean, extensible design ready for future enhancements

## Quick Start

### Prerequisites

- Go >= 1.24

### Installation

### Go

```bash
go install github.com/kozmaoliver/asciify@latest
```

### Manual

```bash
git clone git@github.com:kozmaoliver/asciify.git
cd asciify
go install
```

If you don't want to install it globally, you can just build an executable:
```bash
go build
```

### Usage
Call `asciify` or the executable `./path/to/asciify/asciify` with an image path argument in the terminal.
```bash
# Basic usage
asciify path/to/your/image.png

# With colors
asciify -color colorful_image.jpg
```

## How It Works

The tool follows a sophisticated pipeline to convert images to ASCII:
```
Image → Load → Terminal Sizing → Resize → Luminance → Edge Detection → ASCII Mapping → Render
```

1. **Image Loading**: Supports PNG, JPEG, and GIF formats
2. **Terminal Detection**: Automatically detects your terminal dimensions
3. **Smart Resizing**: Scales image while preserving aspect ratio
4. **Luminance Analysis**: Calculates perceived brightness using L = 0.2126*R + 0.7152*G + 0.0722*B
5. **Edge Detection**: Applies Sobel filters to detect edges and directions
6. **Character Mapping**: Maps brightness to ASCII characters, uses directional chars for edges
7. **Terminal Rendering**: Outputs the final ASCII art to your terminal

> **Tip**: For best results, use images with good contrast and clear subjects. The edge detection works particularly well with architectural photos and portraits!

## Showcase

A generated ASCII image form my profile picture

![image](./docs/resources/demo-1.png)

Same with colors

![image](./docs/resources/demo-1-color.png)

![image](./docs/resources/demo-go-color.png)
