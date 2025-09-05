import { cva, type VariantProps } from 'class-variance-authority';

export const inputCheckboxVariants = cva(
  [
    'cursor-pointer align-middle',
    'focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-offset-2',
    'disabled:cursor-not-allowed disabled:opacity-50',
  ],
  {
    variants: {
      size: {
        sm:   'h-4 w-4',
        base: 'h-5 w-5',
        lg:   'h-6 w-6',
      },
      variant: {
        default: 'accent-primary focus-visible:ring-primary',
        error:   'accent-error focus-visible:ring-error',
        success: 'accent-success focus-visible:ring-success',
        warning: 'accent-warning focus-visible:ring-warning',
      },
      radius: {
        none:  'rounded-none',
        sm:    'rounded-sm',
        md:    'rounded-md',
        lg:    'rounded-lg',
        full:  'rounded-full',
      },
    },
    defaultVariants: {
      size: 'base',
      variant: 'default',
      radius: 'md',
    },
    compoundVariants: [],
  }
);

export type InputCheckboxVariantProps = VariantProps<typeof inputCheckboxVariants>;

