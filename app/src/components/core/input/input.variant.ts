import { cva, type VariantProps } from 'class-variance-authority';

export const inputVariants = cva(
  [
    'w-full bg-transparent outline-none placeholder-gray-400 hover:placeholder-gray-500',
    'disabled:placeholder-gray-300 disabled:hover:placeholder-gray-300',
    'disabled:border-gray-300 disabled:hover:border-gray-300 disabled:cursor-not-allowed',
  ],
  {
    variants: {
      size: {
        sm:   'px-2.5 py-2 text-sm',     // ~40px haut
        base: 'px-3 py-3 text-base',     // ~48px haut
        lg:   'px-4 py-3 text-lg',       // ~48-52px haut
      },
      variant: {
        default: 'border border-1 border-gray-400 hover:border-gray-500 focus:border-primary dark:focus:border-primary-light',
        error:   'border border-1 border-error',
        success: 'border border-1 border-success',
        warning: 'border border-1 border-warning',
      },
      align: {
        left:   '', // texte aligné à gauche
        center: '', // texte centré
        right:  '', // texte aligné à droite
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
      align: 'left',
      radius: 'md',
    },
    compoundVariants: [],
  }
);

export type InputVariantProps = VariantProps<typeof inputVariants>;
export type InputType = 'text' | 'email' | 'password' | 'number' | 'tel' | 'url';
