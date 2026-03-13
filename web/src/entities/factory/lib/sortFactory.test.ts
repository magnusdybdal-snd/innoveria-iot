import { describe, expect, test } from "bun:test";

import type { FactoryApiResponse } from "../model/factorySchema";
import { sortFactories, sortItems } from "./sortFactory";

const makeFactory = (
  overrides: Partial<FactoryApiResponse>,
): FactoryApiResponse => ({
  id: "1",
  companyId: "c1",
  name: "Factory A",
  address: "Street 1",
  createdAt: new Date("2024-01-01"),
  updatedAt: new Date("2024-06-01"),
  ...overrides,
});

const factories: FactoryApiResponse[] = [
  makeFactory({
    id: "1",
    name: "Zebra",
    address: "Oslo",
    createdAt: new Date("2024-03-01"),
    updatedAt: new Date("2024-09-01"),
  }),
  makeFactory({
    id: "2",
    name: "Alpha",
    address: "Bergen",
    createdAt: new Date("2024-01-01"),
    updatedAt: new Date("2024-07-01"),
  }),
  makeFactory({
    id: "3",
    name: "Mango",
    address: "Trondheim",
    createdAt: new Date("2024-06-01"),
    updatedAt: new Date("2024-08-01"),
  }),
];

describe("sortFactories", () => {
  test("returns original array when key is null", () => {
    const result = sortFactories(factories, null, "asc");
    expect(result).toBe(factories);
  });

  test("does not mutate the original array", () => {
    const original = [...factories];
    sortFactories(factories, "Name", "asc");
    expect(factories).toEqual(original);
  });

  test("sorts by Name ascending", () => {
    const result = sortFactories(factories, "Name", "asc");
    expect(result.map((f) => f.name)).toEqual(["Alpha", "Mango", "Zebra"]);
  });

  test("sorts by Name descending", () => {
    const result = sortFactories(factories, "Name", "desc");
    expect(result.map((f) => f.name)).toEqual(["Zebra", "Mango", "Alpha"]);
  });

  test("sorts by Address ascending", () => {
    const result = sortFactories(factories, "Address", "asc");
    expect(result.map((f) => f.address)).toEqual([
      "Bergen",
      "Oslo",
      "Trondheim",
    ]);
  });

  test("sorts by Created at ascending", () => {
    const result = sortFactories(factories, "Created at", "asc");
    expect(result.map((f) => f.id)).toEqual(["2", "1", "3"]);
  });

  test("sorts by Updated at descending", () => {
    const result = sortFactories(factories, "Updated at", "desc");
    expect(result.map((f) => f.id)).toEqual(["1", "3", "2"]);
  });
});

describe("sortItems", () => {
  test("returns original array when key is null", () => {
    const result = sortItems(factories, null, "asc", {});
    expect(result).toBe(factories);
  });

  test("returns original array when key has no comparator", () => {
    const result = sortItems(factories, "Unknown", "asc", {});
    expect(result).toBe(factories);
  });

  test("sorts using a custom comparator", () => {
    const result = sortItems(factories, "Name", "asc", {
      Name: (a, b) => a.name.localeCompare(b.name),
    });
    expect(result.map((f) => f.name)).toEqual(["Alpha", "Mango", "Zebra"]);
  });
});
