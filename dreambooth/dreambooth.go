package dreambooth

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
	ModelId           string      `json:"model_id"`
	Prompt            string      `json:"prompt"`
	NegativePrompt    string      `json:"negative_prompt,omitempty"`
	NumInferenceSteps string      `json:"num_inference_steps,omitempty"`
	SafetyChecker     string      `json:"safety_checker,omitempty"`
	EnhancePrompt     string      `json:"enhance_prompt,omitempty"`
	Seed              interface{} `json:"seed,omitempty"`
	GuidanceScale     float64     `json:"guidance_scale,omitempty"`
	MultiLingual      string      `json:"multi_lingual,omitempty"`
	Panorama          string      `json:"panorama,omitempty"`
	SelfAttention     string      `json:"self_attention,omitempty"`
	Upscale           string      `json:"upscale,omitempty"`
	EmbeddingsModel   interface{} `json:"embeddings_model,omitempty"`
	LoraModel         interface{} `json:"lora_model,omitempty"`
	Tomesd            string      `json:"tomesd,omitempty"`
	ClipSkip          string      `json:"clip_skip,omitempty"`
	UseKarrasSigmas   string      `json:"use_karras_sigmas,omitempty"`
	Vae               interface{} `json:"vae,omitempty"`
	LoraStrength      interface{} `json:"lora_strength,omitempty"`
	Scheduler         string      `json:"scheduler,omitempty"`
	Webhook           interface{} `json:"webhook,omitempty"`
	TrackId           interface{} `json:"track_id,omitempty"`
}

type TextToImageResponse struct {
	Status         string   `json:"status"`
	GenerationTime float64  `json:"generationTime"`
	Id             int      `json:"id"`
	Output         []string `json:"output"`
	Meta           struct {
		Prompt         string      `json:"prompt"`
		ModelId        string      `json:"model_id"`
		NegativePrompt string      `json:"negative_prompt"`
		Scheduler      string      `json:"scheduler"`
		Safetychecker  string      `json:"safetychecker"`
		W              int         `json:"W"`
		H              int         `json:"H"`
		GuidanceScale  float64     `json:"guidance_scale"`
		Seed           int64       `json:"seed"`
		Steps          int         `json:"steps"`
		NSamples       int         `json:"n_samples"`
		FullUrl        string      `json:"full_url"`
		Upscale        string      `json:"upscale"`
		MultiLingual   string      `json:"multi_lingual"`
		Panorama       string      `json:"panorama"`
		SelfAttention  string      `json:"self_attention"`
		Embeddings     interface{} `json:"embeddings"`
		Lora           interface{} `json:"lora"`
		Outdir         string      `json:"outdir"`
		FilePrefix     string      `json:"file_prefix"`
	} `json:"meta"`
}

func (self *Client) TextToImage(req TextToImageRequest, res *TextToImageResponse) error {
	stablediffusionapi.SetApiKeyIfNeeded(&req.RequestBase, self.APIKey)
	stablediffusionapi.SetDefaultRequiredImageFieldsIfNeeded(&req.RequiredImageFields)
	return stablediffusionapi.DoPost(self.Client, "/api/v4/dreambooth", req, res)
}

type ImageToImageRequest struct {
	stablediffusionapi.RequestBase
	stablediffusionapi.RequiredImageFields
	ModelId           string      `json:"model_id"`
	Prompt            string      `json:"prompt"`
	NegativePrompt    interface{} `json:"negative_prompt"`
	InitImage         string      `json:"init_image"`
	NumInferenceSteps string      `json:"num_inference_steps"`
	SafetyChecker     string      `json:"safety_checker"`
	EnhancePrompt     string      `json:"enhance_prompt"`
	GuidanceScale     float64     `json:"guidance_scale"`
	Strength          float64     `json:"strength"`
	Scheduler         string      `json:"scheduler"`
	Seed              interface{} `json:"seed"`
	LoraModel         interface{} `json:"lora_model"`
	Tomesd            string      `json:"tomesd"`
	UseKarrasSigmas   string      `json:"use_karras_sigmas"`
	Vae               interface{} `json:"vae"`
	LoraStrength      interface{} `json:"lora_strength"`
	EmbeddingsModel   interface{} `json:"embeddings_model"`
	Webhook           interface{} `json:"webhook"`
	TrackId           interface{} `json:"track_id"`
}

type ImageToImageResponse struct {
	Status         string   `json:"status"`
	GenerationTime float64  `json:"generationTime"`
	Id             int      `json:"id"`
	Output         []string `json:"output"`
	Meta           struct {
		Prompt         string  `json:"prompt"`
		ModelId        string  `json:"model_id"`
		Scheduler      string  `json:"scheduler"`
		Safetychecker  string  `json:"safetychecker"`
		NegativePrompt string  `json:"negative_prompt"`
		W              int     `json:"W"`
		H              int     `json:"H"`
		GuidanceScale  float64 `json:"guidance_scale"`
		InitImage      string  `json:"init_image"`
		Steps          int     `json:"steps"`
		NSamples       int     `json:"n_samples"`
		Strength       float64 `json:"strength"`
		MultiLingual   string  `json:"multi_lingual"`
		FullUrl        string  `json:"full_url"`
		Upscale        string  `json:"upscale"`
		Seed           int64   `json:"seed"`
		Outdir         string  `json:"outdir"`
		FilePrefix     string  `json:"file_prefix"`
	} `json:"meta"`
}

func (self *Client) ImageToImage(req ImageToImageRequest, res *ImageToImageResponse) error {
	stablediffusionapi.SetApiKeyIfNeeded(&req.RequestBase, self.APIKey)
	return stablediffusionapi.DoPost(self.Client, "/api/v4/dreambooth/img2img", req, res)
}

type InpaintRequest struct {
	stablediffusionapi.RequestBase
	stablediffusionapi.RequiredImageFields
	ModelId         string      `json:"model_id"`
	Prompt          string      `json:"prompt"`
	NegativePrompt  interface{} `json:"negative_prompt"`
	InitImage       string      `json:"init_image"`
	MaskImage       string      `json:"mask_image"`
	Steps           string      `json:"steps"`
	SafetyChecker   string      `json:"safety_checker"`
	EnhancePrompt   string      `json:"enhance_prompt"`
	GuidanceScale   float64     `json:"guidance_scale"`
	Strength        float64     `json:"strength"`
	Scheduler       string      `json:"scheduler"`
	LoraModel       interface{} `json:"lora_model"`
	Tomesd          string      `json:"tomesd"`
	UseKarrasSigmas string      `json:"use_karras_sigmas"`
	Vae             interface{} `json:"vae"`
	LoraStrength    interface{} `json:"lora_strength"`
	EmbeddingsModel interface{} `json:"embeddings_model"`
	Seed            interface{} `json:"seed"`
	Webhook         interface{} `json:"webhook"`
	TrackId         interface{} `json:"track_id"`
}

type InpaintResponse struct {
	Status         string   `json:"status"`
	GenerationTime float64  `json:"generationTime"`
	Id             int      `json:"id"`
	Output         []string `json:"output"`
	Meta           struct {
		Prompt         string  `json:"prompt"`
		ModelId        string  `json:"model_id"`
		Scheduler      string  `json:"scheduler"`
		Safetychecker  string  `json:"safetychecker"`
		NegativePrompt string  `json:"negative_prompt"`
		W              int     `json:"W"`
		H              int     `json:"H"`
		GuidanceScale  float64 `json:"guidance_scale"`
		InitImage      string  `json:"init_image"`
		MaskImage      string  `json:"mask_image"`
		MultiLingual   string  `json:"multi_lingual"`
		Steps          int     `json:"steps"`
		NSamples       int     `json:"n_samples"`
		FullUrl        string  `json:"full_url"`
		Upscale        string  `json:"upscale"`
		Seed           int     `json:"seed"`
		Outdir         string  `json:"outdir"`
		FilePrefix     string  `json:"file_prefix"`
	} `json:"meta"`
}

func (self *Client) Inpaint(req InpaintRequest, res *InpaintResponse) error {
	stablediffusionapi.SetApiKeyIfNeeded(&req.RequestBase, self.APIKey)
	return stablediffusionapi.DoPost(self.Client, "/api/v4/dreambooth/inpaint", req, res)
}

type FetchRequest struct {
	stablediffusionapi.RequestBase
	RequestId string `json:"request_id"`
}

type FetchResponse struct {
	Status string   `json:"status"`
	Id     int      `json:"id"`
	Output []string `json:"output"`
}

func (self *Client) Fetch(req FetchRequest, res *FetchResponse) error {
	stablediffusionapi.SetApiKeyIfNeeded(&req.RequestBase, self.APIKey)
	return stablediffusionapi.DoPost(self.Client, "/api/v4/dreambooth/fetch", req, res)
}

type ModelReloadRequest struct {
	stablediffusionapi.RequestBase
	ModelId string `json:"model_id"`
}

type ModelReloadResponse struct {
}

func (self *Client) ModelReload(req ModelReloadRequest, res *ModelReloadResponse) error {
	stablediffusionapi.SetApiKeyIfNeeded(&req.RequestBase, self.APIKey)
	return stablediffusionapi.DoPost(self.Client, "/api/v4/dreambooth/model_reload", req, res)
}
