import { cva, VariantProps } from 'class-variance-authority';

export const cardVariants = cva(
  [
    'rounded-xl shadow-sm',
    'bg-surface text-text',
    'dark:bg-surface-dark dark:text-text-dark'
  ],
  {
    variants: {
      size: {
        small: 'w-fit p-4',
        medium: 'w-fit p-6',
        large: 'w-fit p-8',
        expanded: 'w-fit p-10',
        extraLarge: 'w-fit p-12',
        long: 'w-fit px-20 py-12',
        tall: 'w-fit px-12 py-20',
        fullSmall: 'w-full p-4',
        fullMedium: 'w-full p-6',
        fullLarge: 'w-full p-8',
        fullExpanded: 'w-full p-10',
        fullExtraLarge: 'w-full p-12',
        fullLong: 'w-full px-20 py-12',
        fullTall: 'w-full px-12 py-20',
      },
    },
    defaultVariants: {
      size: 'medium',
    },
    compoundVariants: []
  }
);

export type CardVariantProps = VariantProps<typeof cardVariants>;
export type CardTag = 'div' | 'article' | 'section';
