/**
 * Returns true if the value is a valid email address.
 * @param value - The string to validate
 * @returns True if the value matches a basic email format, false otherwise
 */
export const isValidEmail = (value: string): boolean =>
  /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(value);
