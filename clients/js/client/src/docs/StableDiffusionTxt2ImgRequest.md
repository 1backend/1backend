
# StableDiffusionTxt2ImgRequest


## Properties

Name | Type
------------ | -------------
`alwayson_scripts` | { [key: string]: string; }
`batch_size` | number
`cfg_scale` | number
`comments` | { [key: string]: string; }
`denoising_strength` | number
`disable_extra_networks` | boolean
`do_not_save_grid` | boolean
`do_not_save_samples` | boolean
`enable_hr` | boolean
`eta` | number
`firstpass_image` | string
`firstphase_height` | number
`firstphase_width` | number
`force_task_id` | string
`height` | number
`hr_checkpoint_name` | string
`hr_negative_prompt` | string
`hr_prompt` | string
`hr_resize_x` | number
`hr_resize_y` | number
`hr_sampler_name` | string
`hr_scale` | number
`hr_scheduler` | string
`hr_second_pass_steps` | number
`hr_upscaler` | string
`infotext` | string
`n_iter` | number
`negative_prompt` | string
`override_settings` | { [key: string]: string; }
`override_settings_restore_afterwards` | boolean
`prompt` | string
`refiner_checkpoint` | string
`refiner_switch_at` | number
`restore_faces` | boolean
`s_churn` | number
`s_min_uncond` | number
`s_noise` | number
`s_tmax` | number
`s_tmin` | number
`sampler_index` | string
`sampler_name` | string
`save_images` | boolean
`scheduler` | string
`script_args` | Array&lt;string&gt;
`script_name` | string
`seed` | number
`seed_resize_from_h` | number
`seed_resize_from_w` | number
`send_images` | boolean
`steps` | number
`styles` | Array&lt;string&gt;
`subseed` | number
`subseed_strength` | number
`tiling` | boolean
`width` | number

## Example

```typescript
import type { StableDiffusionTxt2ImgRequest } from ''

// TODO: Update the object below with actual values
const example = {
  "alwayson_scripts": null,
  "batch_size": null,
  "cfg_scale": null,
  "comments": null,
  "denoising_strength": null,
  "disable_extra_networks": null,
  "do_not_save_grid": null,
  "do_not_save_samples": null,
  "enable_hr": null,
  "eta": null,
  "firstpass_image": null,
  "firstphase_height": null,
  "firstphase_width": null,
  "force_task_id": null,
  "height": null,
  "hr_checkpoint_name": null,
  "hr_negative_prompt": null,
  "hr_prompt": null,
  "hr_resize_x": null,
  "hr_resize_y": null,
  "hr_sampler_name": null,
  "hr_scale": null,
  "hr_scheduler": null,
  "hr_second_pass_steps": null,
  "hr_upscaler": null,
  "infotext": null,
  "n_iter": null,
  "negative_prompt": null,
  "override_settings": null,
  "override_settings_restore_afterwards": null,
  "prompt": null,
  "refiner_checkpoint": null,
  "refiner_switch_at": null,
  "restore_faces": null,
  "s_churn": null,
  "s_min_uncond": null,
  "s_noise": null,
  "s_tmax": null,
  "s_tmin": null,
  "sampler_index": null,
  "sampler_name": null,
  "save_images": null,
  "scheduler": null,
  "script_args": null,
  "script_name": null,
  "seed": null,
  "seed_resize_from_h": null,
  "seed_resize_from_w": null,
  "send_images": null,
  "steps": null,
  "styles": null,
  "subseed": null,
  "subseed_strength": null,
  "tiling": null,
  "width": null,
} satisfies StableDiffusionTxt2ImgRequest

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as StableDiffusionTxt2ImgRequest
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


