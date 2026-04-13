/**
 * Downsamples an array to at most `maxPoints` elements using uniform index sampling.
 * Returns the original array unchanged if it is within the limit.
 * @param items - The source array to downsample
 * @param maxPoints - Maximum number of elements to retain
 * @returns A new array with at most `maxPoints` evenly-spaced elements
 */
export function downsample<T>(items: T[], maxPoints: number = 100): T[] {
  if (items.length <= maxPoints) return items;
  return Array.from(
    { length: maxPoints },
    (_, i) => items[Math.round((i * (items.length - 1)) / (maxPoints - 1))],
  );
}
