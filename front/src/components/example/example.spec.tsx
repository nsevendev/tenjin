import { test, expect, describe } from "vitest";

// Tests de base pour vérifier que Vitest fonctionne
describe('Configuration Vitest', () => {
  test('Vitest est correctement configuré', () => {
    expect(1 + 1).toBe(2);
  });

  test('Les assertions fonctionnent', () => {
    const message = 'Hello Tenjin';
    expect(message).toContain('Tenjin');
    expect(message.length).toBe(12);
  });

  test('Les objets et arrays', () => {
    const user = { name: 'John', age: 30 };
    expect(user).toHaveProperty('name');
    expect(user.name).toBe('John');
    
    const numbers = [1, 2, 3, 4, 5];
    expect(numbers).toHaveLength(5);
    expect(numbers).toContain(3);
  });
});
