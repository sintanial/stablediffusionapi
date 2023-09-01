package dreambooth

import (
	"os"
	"testing"
)

var apiKey = os.Getenv("API_KEY")
var client = NewClient(apiKey)

func TestClient_TextToImage(t *testing.T) {
	var res TextToImageResponse
	if err := client.TextToImage(TextToImageRequest{
		Prompt:  "a portrait of Beautiful caucasian girl in a magic forest ,Sony α7 III camera with a 85mm lens at F 1.2 aperture setting to blur the background and isolate the subject. The image should be shot in high resolution and in a 1:1 aspect ratio with photorealism mode on to create an ultra-realistic image that captures the subject’s natural beauty and personality",
		ModelId: "midjourney",
	}, &res); err != nil {
		t.Fatal(err)
	}

	if res.Status == "error" {
		t.Fatal(res)
	}
}
