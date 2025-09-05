import { cva, type VariantProps } from 'class-variance-authority';

export const buttonVariants = cva(
  [
    'inline-flex items-center justify-center rounded-lg font-semibold cursor-pointer',
    'focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-offset-2',
    'disabled:opacity-50 disabled:cursor-not-allowed',
    'active:scale-95 active:shadow-inner',
  ],
  {
    variants: {
      variant: {
        // solides
        primary:   'bg-primary   text-white hover:bg-primary-dark   focus-visible:ring-primary',
        secondary: 'bg-secondary text-white hover:opacity-90        focus-visible:ring-secondary',
        success:   'bg-success   text-white hover:opacity-90        focus-visible:ring-success',
        error:     'bg-error     text-white hover:opacity-90        focus-visible:ring-error',
        warning:   'bg-warning   text-white hover:opacity-90        focus-visible:ring-warning',
        transparentDefault: 'bg-transparent hover:underline',
        transparentDefaultPrimary: 'bg-transparent hover:text-primary hover:underline        focus-visible:ring-primary',
        transparentPrimary: 'bg-transparent text-primary hover:underline        focus-visible:ring-primary',

        // avec opacity
        primaryOpacity:   'bg-primary/15   text-primary   hover:bg-primary/25   focus-visible:ring-primary',
        secondaryOpacity: 'bg-secondary/15 text-secondary hover:bg-secondary/25 focus-visible:ring-secondary',
        successOpacity:   'bg-success/15   text-success   hover:bg-success/25   focus-visible:ring-success',
        errorOpacity:     'bg-error/15     text-error     hover:bg-error/25     focus-visible:ring-error',
        warningOpacity:   'bg-warning/15   text-warning   hover:bg-warning/25   focus-visible:ring-warning',
      },
      textSize: {
        sm:   'text-xs md:text-sm',
        base: 'text-sm md:text-base',
      },
      size: {
        sm:   'px-3 py-2 md:px-4',
        base: 'px-4 py-3 md:px-6',
      },
      transform: {
        default: '',
        upper:   'uppercase',
      },
    },
    defaultVariants: {
      variant: 'primary',
      size: 'base',
      textSize: 'base',
      transform: 'default',
    },
    compoundVariants: [],
  }
);

export type ButtonVariantProps = VariantProps<typeof buttonVariants>;
