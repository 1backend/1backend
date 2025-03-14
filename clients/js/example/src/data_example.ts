import { Configuration, DataSvcApi } from "@openorch/client";

export async function dynamicTest(apiKey: string) {
  const dataService: DataSvcApi = new DataSvcApi(
    new Configuration({
      apiKey: apiKey,
    })
  );

  await dataService.createObject({
    body: {
      object: {
        table: "uzerz",
        data: {
          fieldA: "valueA",
        },
        readers: ["_self"],
      },
    },
  });

  await dataService.createObject({
    body: {
      object: {
        table: "uzerz",
        data: {
          fieldA: "valueB",
        },
        readers: ["_self"],
      },
    },
  });

  let rsp = await dataService.queryObjects({
    body: {
      table: "uzerz",
      readers: ["_self"],
    },
  });

  if (rsp.objects?.length !== 2) {
    throw "expected result length to be 2";
  }
}
