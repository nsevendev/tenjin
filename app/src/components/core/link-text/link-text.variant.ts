import { cva, VariantProps } from 'class-variance-authority';

export const linkTextVariants = cva(
  [
    'inline cursor-pointer',
    'focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-offset-2',
    'hover:underline underline-offset-2 md:underline-offset-4'
  ],
  {
    variants: {
      color: {
        blue: [
          'text-primary-light hover:text-primary-dark',
          'dark:text-primary-light dark:hover:text-primary',
          'focus-visible:ring-primary-dark dark:focus-visible:ring-primary-light'
        ],
        red: [
          'text-error',
          'focus-visible:ring-error dark:focus-visible:ring-error'
        ],
        green: [
          'text-success',
          'focus-visible:ring-success dark:focus-visible:ring-success'
        ],
      },
      size: {
        sm: 'text-sm',
        base: 'text-sm md:text-base',
        lg: 'text-base md:text-lg',
      },
      underline: {
        always: 'underline',
        hover: '', // deja par defaut
        never: 'no-underline hover:no-underline'
      }
    },
    defaultVariants: {
      color: 'blue',
      size: 'base',
      underline: 'hover'
    },
    compoundVariants: [
      {
        size: 'sm',
        class: 'focus-visible:ring-offset-1'
      },
      {
        size: ['base', 'lg'],
        class: 'focus-visible:ring-offset-2'
      }
    ]
  }
);

export type LinkTextVariantProps = VariantProps<typeof linkTextVariants>;
