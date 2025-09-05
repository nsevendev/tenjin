import { cva, type VariantProps } from 'class-variance-authority';

export const textVariants = cva(
  '',
  {
    variants: {
      variant: {
        paragraph: 'text-text dark:text-text leading-6 md:leading-7',
        small: 'text-sm text-text-muted',
        caption: 'text-[11px] md:text-xs text-text-muted',
        lead: 'text-base md:text-lg text-text-muted leading-7 md:leading-8',
        muted: 'text-text-muted',
        code: 'font-mono text-xs md:text-sm bg-bg-muted dark:bg-bg px-1 py-0.5 rounded text-text dark:text-text',
      },
      size: {
        xs: 'text-xs',
        sm: 'text-sm',
        base: 'text-sm md:text-base',
        lg: 'text-base md:text-lg',
        xl: 'text-lg md:text-xl',
      },
      weight: {
        light: 'font-light',
        normal: 'font-normal',
        medium: 'font-medium',
        semibold: 'font-semibold',
        bold: 'font-bold',
      },
      align: {
        left: 'text-left',
        center: 'text-center',
        right: 'text-right',
        justify: 'text-justify',
      },
      color: {
        default: 'text-text dark:text-text',
        muted: 'text-text-muted',
        primary: 'text-primary',
        secondary: 'text-text-muted',
        success: 'text-success',
        error: 'text-error',
        warning: 'text-warning',
      },
    },
    defaultVariants: {
      variant: 'paragraph',
      size: 'base',
      weight: 'normal',
      align: 'left',
      color: 'default',
    },
  }
);

export const messageVariants = cva(
  '',
  {
    variants: {
      variant: {
        success: 'text-success',
        error: 'text-error',
        warning: 'text-warning',
        info: 'text-info',
      },
      size: {
        sm: 'text-sm',
        base: 'text-base',
        lg: 'text-lg',
      },
    },
    defaultVariants: {
      variant: 'info',
      size: 'base',
    },
  }
);

export const blockquoteVariants = cva(
  'border-l-2 pl-3 md:border-l-4 md:pl-4 italic',
  {
    variants: {
      variant: {
        default: 'border-border text-text',
        primary: 'border-primary text-primary',
        success: 'border-success text-success',
        warning: 'border-warning text-warning',
        error: 'border-error text-error',
      },
    },
    defaultVariants: {
      variant: 'default',
    },
  }
);

export type TextVariantProps = VariantProps<typeof textVariants>;
export type MessageVariantProps = VariantProps<typeof messageVariants>;
export type BlockquoteVariantProps = VariantProps<typeof blockquoteVariants>;
