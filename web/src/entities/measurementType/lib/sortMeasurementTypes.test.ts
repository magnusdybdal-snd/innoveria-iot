import type { MeasurementTypeApiResponse } from "@entities/measurementType";
import { describe, expect, test } from "bun:test";

import { sortMeasurementTypes } from "./sortMeasurementTypes.ts";

const makeMeasurementType = (
  overrides: Partial<MeasurementTypeApiResponse>,
): MeasurementTypeApiResponse => ({
  defaultUnit: "°C",
  deprecated: false,
  description: "Ambient or surface temperature",
  displayName: "Temperature",
  slug: "temperature",
  ...overrides,
});

const measurementTypes: MeasurementTypeApiResponse[] = [
  makeMeasurementType({ defaultUnit: "1", description: "Zebra" }),
  makeMeasurementType({ defaultUnit: "2", description: "Alpha" }),
  makeMeasurementType({ defaultUnit: "3", description: "Mango" }),
];

describe("sortMeasurementTypes", () => {
  test("returns original array when key is null", () => {
    const result = sortMeasurementTypes(measurementTypes, null, "asc");
    expect(result).toBe(measurementTypes);
  });

  test("does not mutate the original array", () => {
    const original = [...measurementTypes];
    sortMeasurementTypes(measurementTypes, "Description", "asc");
    expect(measurementTypes).toEqual(original);
  });

  test("sorts by Description ascending", () => {
    const result = sortMeasurementTypes(measurementTypes, "Description", "asc");
    expect(result.map((s) => s.description)).toEqual([
      "Alpha",
      "Mango",
      "Zebra",
    ]);
  });

  test("sorts by Description descending", () => {
    const result = sortMeasurementTypes(
      measurementTypes,
      "Description",
      "desc",
    );
    expect(result.map((s) => s.description)).toEqual([
      "Zebra",
      "Mango",
      "Alpha",
    ]);
  });
});
