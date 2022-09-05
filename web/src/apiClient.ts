type CutRequestBody = { url: string };
type CutResponse = { url: string };

export const getShortenedUrl = async (url: string): Promise<string> => {
  const request = await fetch('/cut', {
    method: 'POST',
    body: JSON.stringify({ url } as CutRequestBody),
    headers: { 'Content-Type': 'application/json' },
  });

  const json = (await request.json()) as CutResponse;

  return json.url;
};
