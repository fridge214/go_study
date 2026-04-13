package grammar

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"path/filepath"
	"runtime"
)

// SinDemo 绘制带坐标轴与数值刻度的 sin 曲线，输出到同级同名 PNG。
//
// 思路：
// 1. 建立画布与绘图区（留边距）。
// 2. 用复数点 z = x + i*sin(x) 表示曲线上每个采样点。
// 3. 把 x∈[0,2π]、y∈[-1,1] 映射到像素坐标。
// 4. 先画坐标轴和刻度，再画曲线，最后保存 PNG。
func SinDemo() {
	const (
		imgW   = 1000
		imgH   = 600
		margin = 80
	)

	// 定义常用颜色。
	white := color.RGBA{255, 255, 255, 255}
	black := color.RGBA{20, 20, 20, 255}
	blue := color.RGBA{40, 90, 220, 255}
	gray := color.RGBA{170, 170, 170, 255}

	// 新建画布并填充白底。
	canvas := image.NewRGBA(image.Rect(0, 0, imgW, imgH))
	fill(canvas, white)

	// 计算绘图区边界。
	plotLeft := margin
	plotRight := imgW - margin
	plotTop := margin
	plotBottom := imgH - margin
	plotW := plotRight - plotLeft
	plotH := plotBottom - plotTop

	// 画 x 轴（y=0）与 y 轴（x=0）。
	xAxisY := plotTop + plotH/2
	yAxisX := plotLeft
	drawHLine(canvas, plotLeft, plotRight, xAxisY, gray)
	drawVLine(canvas, yAxisX, plotTop, plotBottom, gray)

	// x 轴刻度：0, π/2, π, 3π/2, 2π（用小数显示，便于绘制）。
	xTicks := []struct {
		v     float64
		label string
	}{
		{0, "0.00"},
		{math.Pi / 2, "1.57"},
		{math.Pi, "3.14"},
		{3 * math.Pi / 2, "4.71"},
		{2 * math.Pi, "6.28"},
	}
	for _, t := range xTicks {
		px := plotLeft + int(math.Round((t.v/(2*math.Pi))*float64(plotW)))
		drawVLine(canvas, px, xAxisY-6, xAxisY+6, gray)
		drawString(canvas, px-14, xAxisY+14, t.label, black)
	}

	// y 轴刻度：1.0, 0.5, 0.0, -0.5, -1.0。
	yTicks := []struct {
		v     float64
		label string
	}{
		{1.0, "1.0"},
		{0.5, "0.5"},
		{0.0, "0.0"},
		{-0.5, "-0.5"},
		{-1.0, "-1.0"},
	}
	for _, t := range yTicks {
		py := plotTop + int(math.Round((1-t.v)*float64(plotH)/2))
		drawHLine(canvas, yAxisX-6, yAxisX+6, py, gray)
		drawString(canvas, yAxisX-46, py-4, t.label, black)
	}

	// 按像素采样并绘制 sin 曲线（复数点 z = x + i*sin(x)）。
	for px := plotLeft; px <= plotRight; px++ {
		x := 2 * math.Pi * float64(px-plotLeft) / float64(plotW)
		z := complex(x, math.Sin(x))
		y := imag(z)
		py := plotTop + int(math.Round((1-y)*float64(plotH)/2))

		// 画 3 像素粗细，便于观察。
		plotPoint(canvas, px, py-1, blue)
		plotPoint(canvas, px, py, blue)
		plotPoint(canvas, px, py+1, blue)
	}

	// 画绘图区外框。
	drawHLine(canvas, plotLeft, plotRight, plotTop, black)
	drawHLine(canvas, plotLeft, plotRight, plotBottom, black)
	drawVLine(canvas, plotLeft, plotTop, plotBottom, black)
	drawVLine(canvas, plotRight, plotTop, plotBottom, black)

	// 输出到源码同级目录，文件名与源码同名。
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("resolve source path failed")
		return
	}
	outPath := filepath.Join(filepath.Dir(thisFile), "01_07_sindemo.png")

	f, err := os.Create(outPath)
	if err != nil {
		fmt.Printf("create image failed: %v\n", err)
		return
	}
	defer f.Close()

	if err := png.Encode(f, canvas); err != nil {
		fmt.Printf("encode image failed: %v\n", err)
		return
	}
	fmt.Printf("sine image saved: %s\n", outPath)
}

// fill 用纯色填充整张画布。
func fill(img *image.RGBA, c color.Color) {
	b := img.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			img.Set(x, y, c)
		}
	}
}

// drawHLine 画水平线。
func drawHLine(img *image.RGBA, x1, x2, y int, c color.Color) {
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	for x := x1; x <= x2; x++ {
		plotPoint(img, x, y, c)
	}
}

// drawVLine 画垂直线。
func drawVLine(img *image.RGBA, x, y1, y2 int, c color.Color) {
	if y1 > y2 {
		y1, y2 = y2, y1
	}
	for y := y1; y <= y2; y++ {
		plotPoint(img, x, y, c)
	}
}

// plotPoint 安全设置单像素点（自动越界保护）。
func plotPoint(img *image.RGBA, x, y int, c color.Color) {
	b := img.Bounds()
	if x < b.Min.X || x >= b.Max.X || y < b.Min.Y || y >= b.Max.Y {
		return
	}
	img.Set(x, y, c)
}

// drawString 用 3x5 像素字形绘制字符串（用于刻度数字）。
func drawString(img *image.RGBA, x, y int, s string, c color.Color) {
	cx := x
	for _, ch := range s {
		drawGlyph(img, cx, y, ch, c)
		cx += 4
	}
}

// drawGlyph 绘制单个字符，支持数字、负号和小数点。
func drawGlyph(img *image.RGBA, x, y int, ch rune, col color.Color) {
	patterns := map[rune][]string{
		'0': {"111", "101", "101", "101", "111"},
		'1': {"010", "110", "010", "010", "111"},
		'2': {"111", "001", "111", "100", "111"},
		'3': {"111", "001", "111", "001", "111"},
		'4': {"101", "101", "111", "001", "001"},
		'5': {"111", "100", "111", "001", "111"},
		'6': {"111", "100", "111", "101", "111"},
		'7': {"111", "001", "001", "001", "001"},
		'8': {"111", "101", "111", "101", "111"},
		'9': {"111", "101", "111", "001", "111"},
		'-': {"000", "000", "111", "000", "000"},
		'.': {"000", "000", "000", "000", "010"},
	}
	p, ok := patterns[ch]
	if !ok {
		return
	}
	for row := 0; row < len(p); row++ {
		for colIdx, bit := range p[row] {
			if bit == '1' {
				plotPoint(img, x+colIdx, y+row, col)
			}
		}
	}
}
