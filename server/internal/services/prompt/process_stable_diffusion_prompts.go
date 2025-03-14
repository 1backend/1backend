/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package promptservice

import (
	"context"
	"encoding/base64"
	"log/slog"
	"os"

	"github.com/pkg/errors"

	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/clients/stable_diffusion"
	"github.com/openorch/openorch/sdk/go/logger"

	prompttypes "github.com/openorch/openorch/server/internal/services/prompt/types"
)

func (p *PromptService) processStableDiffusion(
	address string,
	currentPrompt *prompttypes.Prompt,
	model *openapi.ModelSvcGetModelResponse,
) error {
	sd := stable_diffusion.Client{
		Address: address,
	}

	// @todo
	// support the high level `Parameters` field as well,
	// not just the `EngineParameters`

	req := &stable_diffusion.Txt2ImgRequest{}
	if currentPrompt.EngineParameters != nil &&
		currentPrompt.EngineParameters.StableDiffusion != nil &&
		currentPrompt.EngineParameters.StableDiffusion.Txt2Img != nil {
		req = currentPrompt.EngineParameters.StableDiffusion.Txt2Img
	}
	req.Prompt = currentPrompt.Prompt

	if req.Steps == nil {
		req.Steps = openapi.PtrInt(20)
	}
	if req.Width == nil {
		req.Width = openapi.PtrInt(100)
	}
	if req.Height == nil {
		req.Height = openapi.PtrInt(100)
	}
	if req.GuidanceScale == nil {
		req.GuidanceScale = openapi.PtrFloat64(7.5)
	}
	if req.HRScale == nil {
		req.HRScale = openapi.PtrFloat64(2)
	}
	if req.SamplerIndex == nil {
		req.SamplerIndex = openapi.PtrString("Euler")
	}
	if req.NumIterations == nil {
		req.NumIterations = openapi.PtrInt(1)
	}
	if req.RestoreFaces == nil {
		req.RestoreFaces = openapi.PtrBool(true)
	}
	if req.Tiling == nil {
		req.Tiling = openapi.PtrBool(true)
	}

	rsp, err := sd.Txt2Img(*req)
	if err != nil {
		return err
	}
	if len(rsp.Images) == 0 {
		return errors.New("no image in response")
	}

	decodedImage, err := base64.StdEncoding.DecodeString(rsp.Images[0])
	if err != nil {
		return err
	}

	tempFile, err := os.CreateTemp("", "upload-*.png")
	if err != nil {
		return err
	}
	defer tempFile.Close()

	_, err = tempFile.Write(decodedImage)
	if err != nil {
		return err
	}

	// Open the file for uploading
	imageFile, err := os.Open(tempFile.Name())
	if err != nil {
		return err
	}
	defer imageFile.Close()

	uploadRsp, _, err := p.clientFactory.Client(sdk.WithToken(p.token)).
		FileSvcAPI.UploadFile(context.Background()).
		File(imageFile).
		Execute()
	if err != nil {
		return errors.Wrap(err, "error uploading image")
	}

	fileIds := []string{*uploadRsp.Upload.FileId}

	_, _, err = p.clientFactory.Client(sdk.WithToken(p.token)).
		ChatSvcAPI.AddMessage(context.Background(), currentPrompt.ThreadId).
		Body(
			openapi.ChatSvcAddMessageRequest{
				Message: &openapi.ChatSvcMessage{
					Id:       sdk.Id("msg"),
					ThreadId: currentPrompt.ThreadId,
					Text:     openapi.PtrString("Sure, here is your image"),
					FileIds:  fileIds,
					Meta: map[string]interface{}{
						"modelId":    model.Model.Id,
						"platformId": model.Platform.Id,
					},
				},
			},
		).
		Execute()

	if err != nil {
		logger.Error("Error when saving chat message after image generation",
			slog.String("error", err.Error()))
		return err
	}

	return nil
}
