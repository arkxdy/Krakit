export const HTTP_CLIENT_URL =
  process.env.NEXT_PUBLIC_HTTP_CLIENT_CORS_ORIGIN ?? "http://localhost:3001";

export const krakitIcon = "/mnx_logo.ico";

export const PAGE_SIZE = 6;

export const MARKING_TYPE_LABELS: Record<string, string> = {
  positive_only: "Marks awarded for correct answers only",
  positive_negative:
    "Marks awarded for correct answers, negative for wrong answers",
  numerical: "Numerical answers, no options",
};
