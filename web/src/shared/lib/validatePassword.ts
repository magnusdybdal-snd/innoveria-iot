export type PasswordErrors = {
  length: boolean;
  lowercase: boolean;
  uppercase: boolean;
  number: boolean;
  symbol: boolean;
};

/** All-false PasswordErrors value representing a password with no validation violations. Use as initial state or reset value. */
export const emptyPasswordErrors: PasswordErrors = {
  length: false,
  lowercase: false,
  uppercase: false,
  number: false,
  symbol: false,
};

/**
 * Validates a password against the application's strength requirements.
 * @param value - The password string to validate
 * @returns An object where each key is true if that rule is violated
 */
export const validatePassword = (value: string): PasswordErrors => ({
  length: value.length < 8,
  lowercase: !/[a-z]/.test(value),
  uppercase: !/[A-Z]/.test(value),
  number: !/[0-9]/.test(value),
  symbol: !/[!@#$%^&*()_\-+=]/.test(value),
});
