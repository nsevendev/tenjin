import { cva, type VariantProps } from 'class-variance-authority';

export const titleVariants = cva(
  [
    'text-text dark:text-text',
    'focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary focus-visible:ring-offset-2',
  ],
  {
    variants: {
      weight: {
        light:     'font-light',
        normal:    'font-normal',
        semibold:  'font-semibold',
        bold:      'font-bold',
        extrabold: 'font-extrabold',
      },
      size: {
        xs:   'text-xs',
        sm:   'text-sm',
        base: 'text-base',
        lg:   'text-lg',
        xl:   'text-xl',
        '2xl':'text-2xl',
        '3xl':'text-3xl',
        '4xl':'text-4xl',
        '5xl':'text-5xl',
        '6xl':'text-6xl',
      },
    },
    defaultVariants: {
      weight: 'normal',
      size: 'base',
    },
    compoundVariants: [],
  }
);

export type TitleVariantProps = VariantProps<typeof titleVariants>;
export type TitleTag = 'h1' | 'h2' | 'h3' | 'h4' | 'h5' | 'h6';
