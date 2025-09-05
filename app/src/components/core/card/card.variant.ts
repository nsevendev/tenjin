import { cva, VariantProps } from 'class-variance-authority';

export const cardVariants = cva(
  [
    'rounded-lg md:rounded-xl shadow-sm',
    'bg-surface text-text',
    'dark:bg-surface-dark dark:text-text-dark'
  ],
  {
    variants: {
      size: {
        small: 'w-full md:w-fit p-4',
        medium: 'w-full md:w-fit p-4 md:p-6',
        large: 'w-full md:w-fit p-6 md:p-8',
        expanded: 'w-full md:w-fit p-6 md:p-10',
        extraLarge: 'w-full md:w-fit p-8 md:p-12',
        long: 'w-full md:w-fit px-6 py-8 md:px-20 md:py-12',
        tall: 'w-full md:w-fit px-8 py-12 md:px-12 md:py-20',
        fullSmall: 'w-full p-4',
        fullMedium: 'w-full p-4 md:p-6',
        fullLarge: 'w-full p-6 md:p-8',
        fullExpanded: 'w-full p-6 md:p-10',
        fullExtraLarge: 'w-full p-8 md:p-12',
        fullLong: 'w-full px-6 py-8 md:px-20 md:py-12',
        fullTall: 'w-full px-8 py-12 md:px-12 md:py-20',
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
