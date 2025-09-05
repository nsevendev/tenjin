import { cva, type VariantProps } from 'class-variance-authority';

export const selectVariants = cva(
  [
    'w-full bg-transparent outline-none appearance-none',
    'hover:placeholder-gray-500',
    'disabled:placeholder-gray-300 disabled:hover:placeholder-gray-300',
    'disabled:border-gray-300 disabled:hover:border-gray-300 disabled:cursor-not-allowed',
    'text-ellipsis',
  ],
  {
    variants: {
      size: {
        sm:   'px-2.5 py-2 text-sm',
        base: 'px-3 py-3 text-base',
        lg:   'px-4 py-3 text-lg',
      },
      variant: {
        default: 'border border-1 border-gray-400 hover:border-gray-500 focus:border-primary dark:focus:border-primary-light',
        error:   'border border-1 border-error',
        success: 'border border-1 border-success',
        warning: 'border border-1 border-warning',
      },
      radius: {
        none:  'rounded-none',
        sm:    'rounded-sm',
        md:    'rounded-md',
        lg:    'rounded-xl',
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

export type SelectVariantProps = VariantProps<typeof selectVariants>;
export type SelectOptionValue = string | number;
export type SelectOption = {
  label: string;
  value: SelectOptionValue;
  disabled?: boolean;
};

