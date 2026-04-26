import type { PayloadSchemaApiResponse } from "@entities/payloadSchema";

export type SortDirection = "asc" | "desc";
export type PayloadSchemasSortKey =
  | "Chirpstack profile ID"
  | "Measurement type"
  | "Payload key"
  | "Unit";

// Sorting payloadSchemas based on Key (PayloadSchemasSortKey)
/**
 * Returns a sorted copy of the payload schema array based on the given column and direction.
 * @param payloadSchemas - Array of PayloadSchemaApiResponse objects to sort
 * @param key - Column to sort by, or null to return the array unsorted
 * @param direction - Sort order: "asc" or "desc"
 * @returns A new sorted PayloadSchemaApiResponse array (does not mutate the input)
 */
export function sortPayloadSchema(
  payloadSchemas: PayloadSchemaApiResponse[],
  key: PayloadSchemasSortKey | null,
  direction: SortDirection,
): PayloadSchemaApiResponse[] {
  if (!key) return payloadSchemas;
  return [...payloadSchemas].sort((a, b) => {
    let cmp = 0;
    switch (key) {
      case "Chirpstack profile ID":
        cmp = a.chirpstackProfileId.localeCompare(b.chirpstackProfileId);
        break;
      case "Measurement type":
        cmp = a.measurementType.localeCompare(b.measurementType);
        break;
      case "Payload key":
        cmp = a.payloadKey.localeCompare(b.payloadKey);
        break;
      case "Unit":
        cmp = a.unit.localeCompare(b.unit);
        break;
    }
    return direction === "asc" ? cmp : -cmp;
  });
}
