package main

import (
	"bufio"
	"bytes"
	"fmt"
	"image/png"
	"log"
	"os"
	"strings"
	"time"

	"github.com/harrydb/go/img/grayscale"
	"github.com/rudis1261/inkkeys/device"
	"github.com/tarm/serial"
)

const LEDS = 10
const SCREEN_WIDTH = 296
const SCREEN_HEIGHT = 128

type Color struct {
	R int
	G int
	B int
}

func (color Color) HexString() string {
	return fmt.Sprintf("%02x%02x%02x", color.R, color.G, color.B)
}

type HotKey struct {
	stream *serial.Port
}

type LED struct {
	stream *serial.Port
}

func (led LED) SetColors(colors []string) {
	colorCodes := fmt.Sprintf("%s %s\n", device.Command["LED"], strings.Join(colors, " "))
	led.stream.Write([]byte(colorCodes))
	// time.Sleep(time.Microsecond * 1)
}

func (led LED) Rainbow(cycles int) {
	for j:=0; j<256*cycles; j++ {
		colors := make([]string, 0)
		for i := 0; i < LEDS; i++ {
			color := led.Wheel(((i * 256 / LEDS) + j) & 255)
			colors = append(colors, color.HexString())
		}
		led.SetColors(colors)
	}
}

func (led LED) Solid(color Color) {
	colors := make([]string, 0)
	for i := 0; i < LEDS; i++ {
		colors = append(colors, color.HexString())
	}
	led.SetColors(colors)
}

// Input a value 0 to 255 to get a color value.
// The colours are a transition r - g - b - back to r.
func (led LED) Wheel(WheelPos int) Color {
  WheelPos = 255 - WheelPos
	var r, g, b int

	switch(true) {
		case WheelPos < 85:
			r = 255 - (WheelPos * 3) 
			g = 0
			b = WheelPos * 3
			break
		case WheelPos < 170: 
			WheelPos -= 85
			r = 0
			g = WheelPos * 3
			b = 255 - (WheelPos * 3)
			break
		default:
			WheelPos -= 170
			r = WheelPos * 3 
			g = 255 - (WheelPos * 3)
			b = 0
	}
	return Color{R: r % 256, G: g % 256, B: b % 256}
}

type Position struct {
	X int
	Y int
}

type LCDImage struct {
	stream *serial.Port
	Width int
	Height int
	Position Position
	Data []byte
}

func (img LCDImage) SetDimension() {
	img.stream.Write([]byte(
		fmt.Sprintf(
			"%s %d %d %d %d\n", 
			device.Command["DISPLAY"], 
			img.Position.X, 
			img.Position.Y, 
			SCREEN_WIDTH, 
			SCREEN_HEIGHT,
		),
	))
}

func (img LCDImage) EnableScreen() {
	img.stream.Write([]byte(
		fmt.Sprintf(
			"%s %s\n", 
			device.Command["REFRESH"], 
			device.RefreshTypeCode["FULL"],
		),
	))
}

func (img LCDImage) DisableScreen() {
	img.stream.Write([]byte(
		fmt.Sprintf(
			"%s %s\n", 
			device.Command["REFRESH"], 
			device.RefreshTypeCode["OFF"],
		),
	))
}

func (img LCDImage) SendData() {
	img.stream.Write(img.Data)
}

func (img LCDImage) Draw() {
	// def resendImageData(self):
	// for part in self.imageBuffer:
	// 		image = part["image"]
	// 		x = part["x"]
	// 		y = part["y"]
	// 		w, h = image.size
	// 		data = image.convert("1").rotate(180).tobytes()
	// 		self.sendToDevice(CommandCode.DISPLAY.value + " " + str(x) + " " + str(y) + " " + str(w) + " " + str(h))
	// 		self.sendBinaryToDevice(data)
	//    self.imageBuffer = []
	img.SetDimension()
	img.SendData()
	img.EnableScreen()
	img.DisableScreen()
}

func main() {
	config := &serial.Config{
		Name:        "COM20",
		Baud:        115200,
		ReadTimeout: 1,
		Size:        8,
	}

	stream, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal(err)
	}

	// Some LED Magics
	led := LED{stream: stream}

	led.Solid(Color{R: 255, G: 0, B: 0})
	time.Sleep(time.Second)

	led.Solid(Color{R: 0, G: 255, B: 0})
	time.Sleep(time.Second)

	led.Solid(Color{R: 0, G: 0, B: 255})
	time.Sleep(time.Second)

	led.Solid(Color{R: 10, G: 10, B: 10})
	time.Sleep(time.Second)

	filename := "../python-controller/icons/app.png"
	infile, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer infile.Close()

	src, err := png.Decode(infile)
	if err != nil {
		panic(err.Error())
	}
	
	gray := grayscale.Convert(src, grayscale.ToGrayLuminance)
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	png.Encode(w, gray)
	log.Printf("%x", gray.Pix)

	lcd := LCDImage{
		stream: stream, 
		Width: SCREEN_WIDTH,
		Height: SCREEN_HEIGHT,
		Position: Position{ X: 0, Y: 0 }, 
		Data: []byte(gray.Pix), // []byte(fmt.Sprintf("%d", w)),
	}

	lcd.Draw()

	// led.Rainbow(100)
	stream.Write([]byte("I\n"))
	
	buf := make([]byte, 1024)
	for {
		n, err := stream.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		s := string(buf[:n])
		fmt.Println(s)
		newLine := strings.Contains(s, "\n")
		if newLine {
			break
		}
	}

	// time.Sleep(time.Minute)
}
