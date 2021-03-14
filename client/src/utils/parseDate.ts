import moment from "moment";

/* eslint-disable  @typescript-eslint/no-explicit-any */

const isoDateFormat = /^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(?:\.\d*)?Z$/;

function isIsoDateString(value: any): boolean {
  return value && typeof value === "string" && isoDateFormat.test(value);
}

export default function handleDates(body: any) {
  if (body === null || body === undefined || typeof body !== "object")
    return body;

  for (const key of Object.keys(body)) {
    const value = body[key];
    if (isIsoDateString(value)) body[key] = moment(value).toDate();
    else if (typeof value === "object") handleDates(value);
  }
}

/* eslint-enable  @typescript-eslint/no-explicit-any */
