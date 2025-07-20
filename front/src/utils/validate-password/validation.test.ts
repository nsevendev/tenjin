import { test, expect, describe } from "vitest";
import { validatePassword } from "./validation";

describe('Validation Utils', () => {
  describe('validatePassword', () => {
    test('valide un mot de passe fort', () => {
      const result = validatePassword('Password123');
      expect(result.isValid).toBe(true);
      expect(result.errors).toHaveLength(0);
    });

    test('rejette un mot de passe trop court', () => {
      const result = validatePassword('Pass1');
      expect(result.isValid).toBe(false);
      expect(result.errors).toContain('Mot de passe trop court (minimum 8 caractères)');
    });

    test('rejette un mot de passe sans majuscule', () => {
      const result = validatePassword('password123');
      expect(result.isValid).toBe(false);
      expect(result.errors).toContain('Doit contenir au moins une majuscule');
    });
    
    test('rejette un mot de passe sans minuscule', () => {
      const result = validatePassword('PASSWORD123');
      expect(result.isValid).toBe(false);
      expect(result.errors).toContain('Doit contenir au moins une minuscule');
    });
    
    test('rejette un mot de passe sans chiffre', () => {
      const result = validatePassword('Password');
      expect(result.isValid).toBe(false);
      expect(result.errors).toContain('Doit contenir au moins un chiffre');
    });
    
    test('rejette un mot de passe vide', () => {
      const result = validatePassword('');
      expect(result.isValid).toBe(false);
      expect(result.errors).toContain('Mot de passe requis');
    });
  });
});
