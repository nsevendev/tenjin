import { test, expect, describe } from "vitest";
import { getIconSvgProps } from "./icon-svg-hook";

describe('getIconSvgProps', () => {
  test('retourne les valeurs par défaut', () => {
    const result = getIconSvgProps({});
    
    expect(result.sizeValue).toBe(24); // iconSizes.lg
    expect(result.strokeValue).toBe(1.5); // strokeWidths.normal
    expect(result.colorValue).toBe('currentColor');
  });

  test('convertit les tailles nommées en nombres', () => {
    const result = getIconSvgProps({ size: 'xs' });
    expect(result.sizeValue).toBe(12);
  });

  test('garde les tailles numériques', () => {
    const result = getIconSvgProps({ size: 48 });
    expect(result.sizeValue).toBe(48);
  });

  test('convertit les strokeWidths nommés', () => {
    const result = getIconSvgProps({ strokeWidth: 'thick' });
    expect(result.strokeValue).toBe(2);
  });

  test('convertit les couleurs nommées', () => {
    const result = getIconSvgProps({ color: 'blue' });
    expect(result.colorValue).toBe('#3B82F6');
  });
});