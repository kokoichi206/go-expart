package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func (app *Config) pricesTab() *fyne.Container {
	chart := app.getChart()
	chartContainer := container.NewVBox(chart)
	app.PriceChartContainer = chartContainer

	return chartContainer
}

func (app *Config) getChart() *canvas.Image {
	url := fmt.Sprintf("https://goldprice.org/charts/gold_3d_b_o_%s_x.png", strings.ToLower(currency))

	var img *canvas.Image

	err := app.downloadFile(url, "gold-charts.png")
	if err != nil {
		// cannot get resource from url, use default image
		img = canvas.NewImageFromResource(resourceDefaultPng)
	} else {
		img = canvas.NewImageFromFile("./gold-charts.png")
	}

	img.SetMinSize(fyne.Size{
		Width:  770,
		Height: 400,
	})

	img.FillMode = canvas.ImageFillOriginal

	return img
}

func (app *Config) downloadFile(URL, fileName string) error {
	resp, err := app.HTTPClient.Get(URL)
	if err != nil {
		return fmt.Errorf("failed to get url: %w", err)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("wrong response code: %w", err)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	out, err := os.Create(fmt.Sprintf("./%s", fileName))
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	err = png.Encode(out, img)
	if err != nil {
		return fmt.Errorf("failed to encode image: %w", err)
	}

	return nil
}
