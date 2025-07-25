// This Go program generates a Code128 barcode and a QR code,
// saves them as PNG image files, and uses random values to make each barcode unique.
package main

import (
	"image/png"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/qr"
)

func main() {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().Unix())

	// Generate a random number between 0 and 9999
	random := rand.Intn(9999)

	// Generate barcode and QR data strings
	barcodeData := "Gobarcode39randomnumber" + strconv.Itoa(random)
	qrData := "Devzeeh"

	// Encode the barcode data as Code128
	// Code128 is used here instead of Code39 for better compatibility
	code39Barcode, err := code128.Encode(barcodeData)
	if err != nil {
		panic(err) // Stop program if encoding fails
	}

	// Scale the Code128 barcode to width: 400px, height: 100px
	scaledBarcode, err := barcode.Scale(barcode.Barcode(code39Barcode), 400, 100)
	if err != nil {
		panic(err)
	}

	// Encode the QR data
	qrCode, err := qr.Encode(qrData, qr.M, qr.Auto)
	if err != nil {
		panic(err)
	}

	// Scale the QR code to 256x256 pixels
	qrCode, err = barcode.Scale(qrCode, 256, 256)
	if err != nil {
		panic(err)
	}

	// Create the output file for the barcode image
	fileBarcode, err := os.Create("file/code128.png")
	if err != nil {
		panic(err)
	}
	defer fileBarcode.Close() // Ensure file is closed after writing

	// Create the output file for the QR code image
	fileQR, err := os.Create("file/qrcode.png")
	if err != nil {
		panic(err)
	}
	defer fileQR.Close()

	// Write the barcode image to file
	if err = png.Encode(fileBarcode, scaledBarcode); err != nil {
		panic(err)
	}
	if err = png.Encode(fileQR, qrCode); err != nil {
		panic(err)
	}

	// Print confirmation messages
	println("Code39 barcode saved as code128.png")
	println("QR code saved as qrcode.png")
}
