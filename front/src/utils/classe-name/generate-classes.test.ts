import { describe, test, expect } from 'vitest'
import { generateClasses, VariantConfig } from './generate-classes'

describe('generateClasses', () => {
  const baseClasses = 'base-class another-base'
  
  const mockVariantConfig: VariantConfig<'sm' | 'base' | 'lg'> = {
    size: {
      sm: 'text-sm p-2',
      base: 'text-base p-4',
      lg: 'text-lg p-6'
    },
    color: {
      sm: 'bg-blue-100',
      base: 'bg-blue-500',
      lg: 'bg-blue-700'
    }
  }
  
  const defaultVariant = 'base'

  test('should generate classes with default variant when no variant is provided', () => {
    const result = generateClasses(baseClasses, mockVariantConfig, defaultVariant, {})
    
    expect(result).toBe('base-class another-base text-base p-4 bg-blue-500')
  })

  test('should generate classes with specified variant', () => {
    const result = generateClasses(baseClasses, mockVariantConfig, defaultVariant, {
      variant: 'sm'
    })
    
    expect(result).toBe('base-class another-base text-sm p-2 bg-blue-100')
  })

  test('should include custom class when provided', () => {
    const result = generateClasses(baseClasses, mockVariantConfig, defaultVariant, {
      variant: 'lg',
      class: 'custom-class another-custom'
    })
    
    expect(result).toBe('base-class another-base text-lg p-6 bg-blue-700 custom-class another-custom')
  })

  test('should handle empty custom class', () => {
    const result = generateClasses(baseClasses, mockVariantConfig, defaultVariant, {
      variant: 'sm',
      class: ''
    })
    
    expect(result).toBe('base-class another-base text-sm p-2 bg-blue-100')
  })

  test('should handle empty base classes', () => {
    const result = generateClasses('', mockVariantConfig, defaultVariant, {
      variant: 'base',
      class: 'custom-class'
    })
    
    expect(result).toBe('text-base p-4 bg-blue-500 custom-class')
  })

  test('should handle empty variant config', () => {
    const emptyConfig: VariantConfig<'base'> = {}
    
    const result = generateClasses(baseClasses, emptyConfig, 'base', {
      variant: 'base'
    })
    
    expect(result).toBe('base-class another-base')
  })

  test('should handle variant config with empty variant values', () => {
    const configWithEmptyValues: VariantConfig<'sm' | 'base'> = {
      size: {
        sm: '',
        base: 'text-base'
      },
      color: {
        sm: 'bg-red-100',
        base: ''
      }
    }
    
    const result = generateClasses(baseClasses, configWithEmptyValues, 'base', {
      variant: 'sm'
    })
    
    expect(result).toBe('base-class another-base bg-red-100')
  })
})
