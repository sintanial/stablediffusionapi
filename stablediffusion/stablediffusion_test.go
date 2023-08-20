package stablediffusion

import (
	"os"
	"strconv"
	"testing"
)

var apiKey = os.Getenv("API_KEY")
var client = NewClient(apiKey)

func TestClient_TextToImage(t *testing.T) {
	var res TextToImageResponse
	if err := client.TextToImage(TextToImageRequest{
		Prompt: "Generate a photo of a cat",
	}, &res); err != nil {
		t.Fatal(err)
	}

	if res.Status == "error" {
		t.Fatal(res)
	}
}

func TestClient_Fetch(t *testing.T) {
	var txt2img TextToImageResponse
	if err := client.TextToImage(TextToImageRequest{
		Prompt: "Generate a photo of a cat",
	}, &txt2img); err != nil {
		t.Fatal(err)
	}
	if txt2img.Status == "error" {
		t.Fatal(txt2img)
	}

	var res FetchResponse
	if err := client.Fetch(FetchRequest{RequestId: strconv.Itoa(txt2img.Id)}, &res); err != nil {
		t.Fatal(err)
	}
	if res.Status == "error" {
		t.Fatal(res)
	}
}

func TestClient_ImageToImage(t *testing.T) {
	var res ImageToImageResponse
	if err := client.ImageToImage(ImageToImageRequest{
		Prompt:    "Generate a photo of a cat",
		InitImage: "https://cdn2.stablediffusionapi.com/generations/7530f54c-b00a-4765-9bdf-c4315f58f76a-0.png",
		Width:     512,
		Height:    512,
		Samples:   4,
	}, &res); err != nil {
		t.Fatal(err)
	}

	if res.Status == "error" {
		t.Fatal(res)
	}
}

func TestClient_Inpaint(t *testing.T) {
	// todo: implement
}

func TestClient_SuperResolution(t *testing.T) {
	// todo: implement
}

func TestClient_SystemLoad(t *testing.T) {
	// todo: implement
}
