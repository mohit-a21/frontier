// @ts-ignore
export const fetcher = (args) =>
  fetch(args, {
    headers: {
      "X-Shield-Email": "admin@odpf.io",
    },
  }).then((res) => res.json());

export async function update(
  url: string,
  { arg }: { arg: Record<string, string> }
) {
  await fetch(url, {
    method: "POST",
    headers: {
      "X-Shield-Email": "admin@odpf.io",
    },
    body: JSON.stringify(arg),
  });
}
