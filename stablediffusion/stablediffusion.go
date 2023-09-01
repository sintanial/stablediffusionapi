package stablediffusion

import (
	"github.com/sintanial/stablediffusionapi"
	"net/http"
)

type Client struct {
	APIKey string
	Client *http.Client
}

func NewClient(APIKey string) *Client {
	return &Client{
		APIKey: APIKey,
		Client: http.DefaultClient,
	}
}

type TextToImageRequest struct {
	stablediffusionapi.RequestBase
	stablediffusionapi.RequiredImageFields
	//Text prompt with description of the things you want in the image to be generated
	Prompt string `json:"prompt"`
	//Items you don't want in the image.
	NegativePrompt string `json:"negative_prompt,omitempty"`
	//Number of denoising steps. Available values: 21, 31, 41, 51.
	NumInferenceSteps string `json:"num_inference_steps,omitempty"`
	//A checker for NSFW images. If such an image is detected, it will be replaced by a blank image.
	SafetyChecker string `json:"safety_checker,omitempty"`
	//Enhance prompts for better results; default: yes, options: yes/no.
	EnhancePrompt string `json:"enhance_prompt,omitempty"`
	//Seed is used to reproduce results, same seed will give you same image in return again. Pass null for a random number.
	Seed interface{} `json:"seed,omitempty"`
	//Scale for classifier-free guidance (minimum: 1; maximum: 20).
	GuidanceScale float64 `json:"guidance_scale,omitempty"`
	//Allow multi lingual prompt to generate images. Use "no" for the default English.
	MultiLingual string `json:"multi_lingual,omitempty"`
	//Set this parameter to "yes" to generate a panorama image.
	Panorama string `json:"panorama,omitempty"`
	//If you want a high quality image, set this parameter to "yes". In this case the image generation will take more time.
	SelfAttention string `json:"self_attention,omitempty"`
	//Set this parameter to "yes" if you want to upscale the given image resolution two times (2x). If the requested resolution is 512 x 512 px, the generated image will be 1024 x 1024 px.
	Upscale string `json:"upscale,omitempty"`
	//This is used to pass an embeddings model (embeddings_model_id).
	EmbeddingsModel string `json:"embeddings_model,omitempty"`
	//Set an URL to get a POST API call once the image generation is complete.
	Webhook string `json:"webhook,omitempty"`
	//This ID is returned in the response to the webhook API call. This will be used to identify the webhook request.
	TrackId string `json:"track_id,omitempty"`
}

type TextToImageResponse struct {
	stablediffusionapi.ResponseBase
	Id             int       `json:"id,omitempty"`
	GenerationTime float64   `json:"generationTime,omitempty"`
	Output         *[]string `json:"output,omitempty"`
	Meta           *struct {
		H                      int     `json:"H"`
		W                      int     `json:"W"`
		EnableAttentionSlicing string  `json:"enable_attention_slicing"`
		FilePrefix             string  `json:"file_prefix"`
		GuidanceScale          float64 `json:"guidance_scale"`
		Model                  string  `json:"model"`
		NSamples               int     `json:"n_samples"`
		NegativePrompt         string  `json:"negative_prompt"`
		Outdir                 string  `json:"outdir"`
		Prompt                 string  `json:"prompt"`
		Revision               string  `json:"revision"`
		Safetychecker          string  `json:"safetychecker"`
		Seed                   int64   `json:"seed"`
		Steps                  int     `json:"steps"`
		Vae                    string  `json:"vae"`
	} `json:"meta,omitempty"`
}

func (self *Client) TextToImage(req TextToImageRequest, res *TextToImageResponse) error {
	stablediffusionapi.SetApiKeyIfNeeded(&req.RequestBase, self.APIKey)
	stablediffusionapi.SetDefaultRequiredImageFieldsIfNeeded(&req.RequiredImageFields)
	return stablediffusionapi.DoPost(self.Client, "/api/v3/text2img", req, res)
}

type ImageToImageRequest struct {
	stablediffusionapi.RequestBase
	Prompt string `json:"prompt"`
	stablediffusionapi.RequiredImageFields
	InitImage         string  `json:"init_image,omitempty"`
	Width             int     `json:"width,omitempty"`
	Height            int     `json:"height,omitempty"`
	Samples           int     `json:"samples,omitempty"`
	NumInferenceSteps string  `json:"num_inference_steps,omitempty"`
	SafetyChecker     string  `json:"safety_checker,omitempty"`
	EnhancePrompt     string  `json:"enhance_prompt,omitempty"`
	GuidanceScale     float64 `json:"guidance_scale,omitempty"`
	Strength          float64 `json:"strength,omitempty"`
	Seed              string  `json:"seed,omitempty"`
	Webhook           string  `json:"webhook,omitempty"`
	TrackId           string  `json:"track_id,omitempty"`
}

type ImageToImageResponse struct {
	stablediffusionapi.ResponseBase
	Id             int      `json:"id"`
	GenerationTime float64  `json:"generationTime"`
	Output         []string `json:"output"`
	Meta           *struct {
		H                      int     `json:"H"`
		W                      int     `json:"W"`
		EnableAttentionSlicing string  `json:"enable_attention_slicing"`
		FilePrefix             string  `json:"file_prefix"`
		GuidanceScale          float64 `json:"guidance_scale"`
		Model                  string  `json:"model"`
		NSamples               int     `json:"n_samples"`
		NegativePrompt         string  `json:"negative_prompt"`
		Outdir                 string  `json:"outdir"`
		Prompt                 string  `json:"prompt"`
		Revision               string  `json:"revision"`
		SafetyChecker          string  `json:"safety_checker"`
		Seed                   int     `json:"seed"`
		Steps                  int     `json:"steps"`
		Vae                    string  `json:"vae"`
	} `json:"meta"`
}

func (self *Client) ImageToImage(req ImageToImageRequest, res *ImageToImageResponse) error {
	stablediffusionapi.SetApiKeyIfNeeded(&req.RequestBase, self.APIKey)
	stablediffusionapi.SetDefaultRequiredImageFieldsIfNeeded(&req.RequiredImageFields)
	return stablediffusionapi.DoPost(self.Client, "/api/v3/img2img", req, res)
}

type InpaintRequest struct {
	stablediffusionapi.RequestBase
	stablediffusionapi.RequiredImageFields
	Prompt            string      `json:"prompt"`
	NegativePrompt    string      `json:"negative_prompt,omitempty"`
	InitImage         string      `json:"init_image"`
	MaskImage         string      `json:"mask_image,omitempty"`
	NumInferenceSteps string      `json:"num_inference_steps,omitempty"`
	SafetyChecker     string      `json:"safety_checker,omitempty"`
	EnhancePrompt     string      `json:"enhance_prompt,omitempty"`
	GuidanceScale     float64     `json:"guidance_scale,omitempty"`
	Strength          float64     `json:"strength,omitempty"`
	Seed              interface{} `json:"seed,omitempty"`
	Webhook           interface{} `json:"webhook,omitempty"`
	TrackId           interface{} `json:"track_id,omitempty"`
}

type InpaintResponse struct {
	stablediffusionapi.ResponseBase
	stablediffusionapi.RequiredImageFields
	GenerationTime float64  `json:"generationTime"`
	Id             int      `json:"id"`
	Output         []string `json:"output"`
	Meta           struct {
		H              int     `json:"H"`
		W              int     `json:"W"`
		FilePrefix     string  `json:"file_prefix"`
		GuidanceScale  float64 `json:"guidance_scale"`
		InitImage      string  `json:"init_image"`
		MaskImage      string  `json:"mask_image"`
		NSamples       int     `json:"n_samples"`
		NegativePrompt string  `json:"negative_prompt"`
		Outdir         string  `json:"outdir"`
		Prompt         string  `json:"prompt"`
		Safetychecker  string  `json:"safetychecker"`
		Seed           int64   `json:"seed"`
		Steps          int     `json:"steps"`
		Strength       float64 `json:"strength"`
	} `json:"meta"`
}

func (self *Client) Inpaint(req InpaintRequest, res *InpaintResponse) error {
	stablediffusionapi.SetApiKeyIfNeeded(&req.RequestBase, self.APIKey)
	stablediffusionapi.SetDefaultRequiredImageFieldsIfNeeded(&req.RequiredImageFields)
	return stablediffusionapi.DoPost(self.Client, "/api/v3/inpaint", req, res)
}

type FetchRequest struct {
	stablediffusionapi.RequestBase
	RequestId string `json:"request_id,omitempty"`
}

type FetchResponse struct {
	stablediffusionapi.ResponseBase
	Id     int      `json:"id"`
	Output []string `json:"output"`
}

func (self *Client) Fetch(req FetchRequest, res *FetchResponse) error {
	stablediffusionapi.SetApiKeyIfNeeded(&req.RequestBase, self.APIKey)
	return stablediffusionapi.DoPost(self.Client, "/api/v3/fetch/"+req.RequestId, req.RequestBase, res)
}

type SystemLoadRequest struct {
	stablediffusionapi.RequestBase
}

type SystemLoadResponse struct {
	QueueNum  int    `json:"queue_num"`
	QueueTime int    `json:"queue_time"`
	Status    string `json:"status"`
}

func (self *Client) SystemLoad(req SystemLoadRequest, res *SystemLoadResponse) error {
	stablediffusionapi.SetApiKeyIfNeeded(&req.RequestBase, self.APIKey)
	return stablediffusionapi.DoPost(self.Client, "/api/v3/system_load", req, res)
}

type SuperResolutionRequest struct {
	stablediffusionapi.RequestBase
	Url         string      `json:"url"`
	Scale       int         `json:"scale"`
	Webhook     interface{} `json:"webhook"`
	FaceEnhance bool        `json:"face_enhance"`
}

type SuperResolutionResponse struct {
	stablediffusionapi.ResponseBase
	GenerationTime float64 `json:"generationTime"`
	Id             int     `json:"id"`
	Output         string  `json:"output"`
}

func (self *Client) SuperResolution(req SuperResolutionRequest, res *SuperResolutionResponse) error {
	stablediffusionapi.SetApiKeyIfNeeded(&req.RequestBase, self.APIKey)
	return stablediffusionapi.DoPost(self.Client, "/api/v3/system_load", req, res)
}
