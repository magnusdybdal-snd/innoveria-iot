import { describe, expect, test } from "bun:test";

import type { MeasureTypeApiResponse } from "../model/measureTypeSchema.ts";
import { sortMeasureTypes } from "./sortMeasureTypes.ts";

const makeMeasureType = (
  overrides: Partial<MeasureTypeApiResponse>,
): MeasureTypeApiResponse => ({
  defaultUnit: "°C",
  deprecated: false,
  description: "Ambient or surface temperature",
  displayName: "Temperature",
  slug: "temperature",
  ...overrides,
});

const measureTypes: MeasureTypeApiResponse[] = [
  makeMeasureType({ defaultUnit: "1", description: "Zebra" }),
  makeMeasureType({ defaultUnit: "2", description: "Alpha" }),
  makeMeasureType({ defaultUnit: "3", description: "Mango" }),
];

describe("sortMeasureTypes", () => {
  test("returns original array when key is null", () => {
    const result = sortMeasureTypes(measureTypes, null, "asc");
    expect(result).toBe(measureTypes);
  });

  test("does not mutate the original array", () => {
    const original = [...measureTypes];
    sortMeasureTypes(measureTypes, "Description", "asc");
    expect(measureTypes).toEqual(original);
  });

  test("sorts by Description ascending", () => {
    const result = sortMeasureTypes(measureTypes, "Description", "asc");
    expect(result.map((s) => s.description)).toEqual([
      "Alpha",
      "Mango",
      "Zebra",
    ]);
  });

  test("sorts by Description descending", () => {
    const result = sortMeasureTypes(measureTypes, "Description", "desc");
    expect(result.map((s) => s.description)).toEqual([
      "Zebra",
      "Mango",
      "Alpha",
    ]);
  });
});
