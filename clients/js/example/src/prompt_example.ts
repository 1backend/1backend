import {
  Configuration,
  PromptSvcApi,
  ModelSvcApi,
  FileSvcApi,
} from "@1backend/client";
import { ResponseError } from "@1backend/client";

const tinyLamaModelId = `huggingface/TheBloke/tinyllama-1.1b-chat-v1.0.Q4_K_S.gguf`;
const tinyLamaAssetURL = `https://huggingface.co/TheBloke/TinyLlama-1.1B-Chat-v1.0-GGUF/resolve/main/tinyllama-1.1b-chat-v1.0.Q4_K_S.gguf?download=true`;

export async function promptTest(apiKey: string) {
  const promptSvc: PromptSvcApi = new PromptSvcApi(
    new Configuration({
      apiKey: apiKey,
    })
  );

  const modelSvc: ModelSvcApi = new ModelSvcApi(
    new Configuration({
      apiKey: apiKey,
    })
  );

  const fileSvc: FileSvcApi = new FileSvcApi(
    new Configuration({
      apiKey: apiKey,
    })
  );

  await fileSvc.downloadFile({
    body: {
      url: tinyLamaAssetURL,
    },
  });

  let exists = false;
  while (!exists) {
    const dlResponse = await fileSvc.getDownload({
      url: tinyLamaAssetURL,
    });

    console.log(dlResponse);
    const exists = dlResponse._exists;

    if (!exists || dlResponse.download?.status !== "completed") {
      await new Promise((resolve) => setTimeout(resolve, 2000));
    } else {
      break;
    }
  }

  console.log(`Making model with ID ${tinyLamaModelId} default`);

  await modelSvc.makeDefault({
    modelId: tinyLamaModelId,
  });

  console.log("Starting the default model");

  await modelSvc.startDefaultModel();

  console.log("Prompting");

  try {
    const promptRsp = await Promise.race([
      promptSvc.prompt({
        body: {
          sync: true,
          prompt: "Is a cat an animal? Just answer with yes or no please.",
        },
      }),
      timeout(60000),
    ]);

    console.log("Prompt response:", promptRsp);

    console.log(promptRsp);
  } catch (error: any) {
    console.error("Prompt failed:", error);

    if (error instanceof ResponseError && error.response) {
      const res = error.response;
      console.error("HTTP Status:", res.status);
      console.error("URL:", res.url);

      try {
        const bodyText = await res.text();
        console.error("HTTP Response Body:", bodyText);

        // Optional: try to parse as JSON if it might be structured
        try {
          const json = JSON.parse(bodyText);
          console.error("Parsed JSON Error:", json);
        } catch {
          // Not JSON, just log as text
        }
      } catch (readErr) {
        console.error("Failed to read error response body:", readErr);
      }
    }

    // Make test fail
    throw error;
  }
}

const timeout = (ms: number) =>
  new Promise((_, reject) =>
    setTimeout(() => reject(new Error("Request timed out")), ms)
  );
