import { describe, expect, test } from "bun:test";

import { hasPayloadData } from "@entities/sensor/lib/hasPayloadData";
import type { SensorReadingApiResponse } from "@entities/sensor/model/sensorSchema";

const makeReading = (
  payload: SensorReadingApiResponse["payload"],
): SensorReadingApiResponse => ({
  deviceEui: "abc123",
  timestamp: "2026-01-01T00:00:00Z",
  companyId: "company-1",
  payload,
});

describe("hasPayloadData", () => {
  test("returns false when reading is null", () => {
    expect(hasPayloadData(null)).toBe(false);
  });

  test("returns false when payload is null", () => {
    expect(hasPayloadData(makeReading(null))).toBe(false);
  });

  test("returns false when payload is an empty object", () => {
    expect(hasPayloadData(makeReading({}))).toBe(false);
  });

  test("returns true when payload has at least one key", () => {
    expect(hasPayloadData(makeReading({ temperature: 23.5 }))).toBe(true);
  });

  test("returns true when payload has multiple keys", () => {
    expect(
      hasPayloadData(makeReading({ temperature: 23.5, humidity: 60 })),
    ).toBe(true);
  });
});
