import { cn } from './cn';

describe('cn', () => {
  it('should merge class names correctly', () => {
    const result = cn('p-4', 'text-center', 'bg-blue-500');
    expect(result).toBe('p-4 text-center bg-blue-500');
  });

  it('should handle conflicting tailwind classes', () => {
    const result = cn('p-4 p-2', 'text-center text-left');
    expect(result).toBe('p-2 text-left');
  });
});