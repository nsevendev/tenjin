import { test, expect, describe } from "vitest";
import { LockCloseIcon } from "./lock-close";

describe('LockCloseIcon', () => {
  test('le composant est défini et exporté', () => {
    expect(LockCloseIcon).toBeDefined();
    expect(typeof LockCloseIcon).toBe('function');
  });
});