import { Configuration, UserSvcApi } from "@1backend/client";
import { dynamicTest } from "./data_example.js";
import { promptTest } from "./prompt_example.js";

async function start() {
  let userService = new UserSvcApi();
  let loginResponse = await userService.login({
    body: {
      slug: "1backend",
      password: "changeme",
    },
  });

  const token = loginResponse.token?.token;

  userService = new UserSvcApi(
    new Configuration({
      apiKey: token!,
    })
  );

  const readTokenResponse = await userService.readSelf();
  if (readTokenResponse.user?.slug !== "1backend") {
    process.exit(1);
  }

  dynamicTest(token!);
  promptTest(token!);
}

start();
