/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package stable_diffusion

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log/slog"
	"net/http"

	"github.com/pkg/errors"

	"github.com/openorch/openorch/sdk/go/logger"
)

type Client struct {
	Address string
}

func NewClient(address string) *Client {
	return &Client{
		Address: address,
	}
}

// /sdapi/v1/txt2img
type Txt2ImgRequest struct {
	Prompt                  string            `json:"prompt,omitempty"`
	NegativePrompt          *string           `json:"negative_prompt,omitempty"`
	Styles                  *[]string         `json:"styles,omitempty"`
	Seed                    *int              `json:"seed,omitempty"`
	Subseed                 *int              `json:"subseed,omitempty"`
	SubseedStrength         *float64          `json:"subseed_strength,omitempty"`
	SeedResizeFromH         *int              `json:"seed_resize_from_h,omitempty"`
	SeedResizeFromW         *int              `json:"seed_resize_from_w,omitempty"`
	SamplerName             *string           `json:"sampler_name,omitempty"`
	Scheduler               *string           `json:"scheduler,omitempty"`
	BatchSize               *int              `json:"batch_size,omitempty"`
	NumIterations           *int              `json:"n_iter,omitempty"`
	Steps                   *int              `json:"steps,omitempty"`
	GuidanceScale           *float64          `json:"cfg_scale,omitempty"`
	Width                   *int              `json:"width,omitempty"`
	Height                  *int              `json:"height,omitempty"`
	RestoreFaces            *bool             `json:"restore_faces,omitempty"`
	Tiling                  *bool             `json:"tiling,omitempty"`
	DoNotSaveSamples        *bool             `json:"do_not_save_samples,omitempty"`
	DoNotSaveGrid           *bool             `json:"do_not_save_grid,omitempty"`
	Eta                     *float64          `json:"eta,omitempty"`
	DenoisingStrength       *float64          `json:"denoising_strength,omitempty"`
	SMinUncond              *float64          `json:"s_min_uncond,omitempty"`
	SChurn                  *float64          `json:"s_churn,omitempty"`
	STmax                   *float64          `json:"s_tmax,omitempty"`
	STmin                   *float64          `json:"s_tmin,omitempty"`
	SNoise                  *float64          `json:"s_noise,omitempty"`
	OverrideSettings        map[string]string `json:"override_settings,omitempty"`
	OverrideSettingsRestore *bool             `json:"override_settings_restore_afterwards,omitempty"`
	RefinerCheckpoint       *string           `json:"refiner_checkpoint,omitempty"`
	RefinerSwitchAt         *float64          `json:"refiner_switch_at,omitempty"`
	DisableExtraNetworks    *bool             `json:"disable_extra_networks,omitempty"`
	FirstPassImage          *string           `json:"firstpass_image,omitempty"`
	Comments                map[string]string `json:"comments,omitempty"`
	EnableHR                *bool             `json:"enable_hr,omitempty"`
	FirstPhaseWidth         *int              `json:"firstphase_width,omitempty"`
	FirstPhaseHeight        *int              `json:"firstphase_height,omitempty"`
	HRScale                 *float64          `json:"hr_scale,omitempty"`
	HRUpscaler              *string           `json:"hr_upscaler,omitempty"`
	HRSecondPassSteps       *int              `json:"hr_second_pass_steps,omitempty"`
	HRResizeX               *int              `json:"hr_resize_x,omitempty"`
	HRResizeY               *int              `json:"hr_resize_y,omitempty"`
	HRCheckpointName        *string           `json:"hr_checkpoint_name,omitempty"`
	HRSamplerName           *string           `json:"hr_sampler_name,omitempty"`
	HRScheduler             *string           `json:"hr_scheduler,omitempty"`
	HRPrompt                *string           `json:"hr_prompt,omitempty"`
	HRNegativePrompt        *string           `json:"hr_negative_prompt,omitempty"`
	ForceTaskID             *string           `json:"force_task_id,omitempty"`
	SamplerIndex            *string           `json:"sampler_index,omitempty"`
	ScriptName              *string           `json:"script_name,omitempty"`
	ScriptArgs              *[]string         `json:"script_args,omitempty"`
	SendImages              *bool             `json:"send_images,omitempty"`
	SaveImages              *bool             `json:"save_images,omitempty"`
	AlwaysOnScripts         map[string]string `json:"alwayson_scripts,omitempty"`
	InfoText                *string           `json:"infotext,omitempty"`
}

type Txt2ImgResponse struct {
	Images     []string               `json:"images"`
	Parameters map[string]interface{} `json:"parameters"`
	Info       string                 `json:"info"`
}

/*
FileURL returns the local stable diffusion URL of a fileName acquired through the stable diffusion API.

eg. http://127.0.0.1:8001/file=/tmp/tmppjk80zb6/tmpv4bk66e9.png
*/
func FileURL(addr string, fileName string) string {
	return fmt.Sprintf("%v:/file=%v", addr, fileName)
}

// GetImageAsBase64 fetches the image from the given URL and returns it as a base64 encoded string.
func GetImageAsBase64(imageURL string) (string, error) {
	imageData, err := GetImage(imageURL)
	if err != nil {
		return "", errors.Wrap(err, "error getting image")
	}
	base64Image := base64.StdEncoding.EncodeToString(imageData)

	return base64Image, nil
}

// GetImage fetches the image from the given URL.
func GetImage(imageURL string) ([]byte, error) {
	resp, err := http.Get(imageURL)
	if err != nil {
		return nil, errors.New("failed to fetch image: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch image: status code " + resp.Status)
	}

	return io.ReadAll(resp.Body)
}

func (c *Client) Txt2Img(req Txt2ImgRequest) (*Txt2ImgResponse, error) {
	url := c.Address + "/sdapi/v1/txt2img"

	jsonBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Accept", "*/*")
	httpReq.Header.Set("Connection", "keep-alive")
	httpReq.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("API request failed with status code: " + resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var predictResp Txt2ImgResponse
	err = json.Unmarshal(body, &predictResp)
	if err != nil {
		logger.Error("Mismatch between types and StableDiffusion response",
			slog.String("responseJSON", string(body)),
		)
		return nil, err
	}

	return &predictResp, nil
}
