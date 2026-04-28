import { describe, expect, test } from "bun:test";

import type { SensorApiResponse } from "../model/sensorSchema";
import { sortSensors } from "./sortSensors";

const makeSensor = (
  overrides: Partial<SensorApiResponse>,
): SensorApiResponse => ({
  id: "1",
  deviceEui: "abc",
  name: "Sensor A",
  factory: "Factory A",
  status: 0,
  machine: "Machine A",
  lastReading: "1 min",
  deviceProfileId: "profile-a",
  ...overrides,
});

const sensors: SensorApiResponse[] = [
  makeSensor({ id: "1", name: "Zebra", status: 2 }),
  makeSensor({ id: "2", name: "Alpha", status: 0 }),
  makeSensor({ id: "3", name: "Mango", status: 1 }),
];

describe("sortSensors", () => {
  test("returns original array when key is null", () => {
    const result = sortSensors(sensors, null, "asc");
    expect(result).toBe(sensors);
  });

  test("does not mutate the original array", () => {
    const original = [...sensors];
    sortSensors(sensors, "Name", "asc");
    expect(sensors).toEqual(original);
  });

  test("sorts by Name ascending", () => {
    const result = sortSensors(sensors, "Name", "asc");
    expect(result.map((s) => s.name)).toEqual(["Alpha", "Mango", "Zebra"]);
  });

  test("sorts by Name descending", () => {
    const result = sortSensors(sensors, "Name", "desc");
    expect(result.map((s) => s.name)).toEqual(["Zebra", "Mango", "Alpha"]);
  });

  test("sorts by Status ascending", () => {
    const result = sortSensors(sensors, "Status", "asc");
    expect(result.map((s) => s.status)).toEqual([0, 1, 2]);
  });

  test("sorts by Status descending", () => {
    const result = sortSensors(sensors, "Status", "desc");
    expect(result.map((s) => s.status)).toEqual([2, 1, 0]);
  });
});
