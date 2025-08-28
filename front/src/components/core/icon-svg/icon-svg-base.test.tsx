import { test, expect, describe } from "vitest";
import { getIconSvgProps } from "./icon-svg-hook";

describe('IconSvgBase', () => {
  test('les props sont correctement traitées par le hook', () => {
    const result = getIconSvgProps({
      size: 'xl',
      color: 'blue',
      strokeWidth: 'thick'
    });
    
    expect(result.sizeValue).toBe(32);
    expect(result.colorValue).toBe('#3B82F6');
    expect(result.strokeValue).toBe(2);
  });

  test('utilise les valeurs par défaut', () => {
    const result = getIconSvgProps({});
    
    expect(result.sizeValue).toBe(24);
    expect(result.colorValue).toBe('currentColor');
    expect(result.strokeValue).toBe(1.5);
  });
});