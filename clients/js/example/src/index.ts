import { Configuration, UserSvcApi } from "@openorch/client";
import { dynamicTest } from "./data_example.js";
import { promptTest } from "./prompt_example.js";

async function start() {
  let userService = new UserSvcApi();
  let loginResponse = await userService.login({
    body: {
      slug: "openorch",
      password: "changeme",
    },
  });

  const token = loginResponse.token?.token;

  userService = new UserSvcApi(
    new Configuration({
      apiKey: token!,
    })
  );

  const readTokenResponse = await userService.readUserByToken();
  if (readTokenResponse.user?.slug !== "openorch") {
    process.exit(1);
  }

  dynamicTest(token!);
  promptTest(token!);
}

start();
