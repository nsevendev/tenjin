import {QwikIntrinsicElements} from "@builder.io/qwik";
import {cva, VariantProps} from "class-variance-authority";

export const textVariants = cva(
    ``, // base
    {
        variants: {
            variant: {
                paragraph: 'text-gray-700 dark:text-gray-300',
                small: 'text-sm text-gray-600 dark:text-gray-400',
                caption: 'text-xs text-gray-500 dark:text-gray-500',
                lead: 'text-lg text-gray-600 dark:text-gray-400',
                muted: 'text-gray-500 dark:text-gray-400',
                code: 'font-mono text-sm bg-gray-100 dark:bg-gray-800 px-1 py-0.5 rounded text-gray-800 dark:text-gray-200',
            },
            size: {
                xs: 'text-xs',
                sm: 'text-sm',
                base: 'text-base',
                lg: 'text-lg',
                xl: 'text-xl',
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
                default: 'text-gray-900 dark:text-gray-100',
                muted: 'text-gray-600 dark:text-gray-400',
                primary: 'text-blue-600 dark:text-blue-400',
                secondary: 'text-gray-600 dark:text-gray-400',
                success: 'text-green-600 dark:text-green-400',
                error: 'text-red-600 dark:text-red-400',
                warning: 'text-yellow-600 dark:text-yellow-400',
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
    `rounded-lg p-4 border`, // base
    {
        variants: {
            variant: {
                success: 'bg-green-50 border-green-200 text-green-800 dark:bg-green-900/20 dark:border-green-800 dark:text-green-200',
                error: 'bg-red-50 border-red-200 text-red-800 dark:bg-red-900/20 dark:border-red-800 dark:text-red-200',
                warning: 'bg-yellow-50 border-yellow-200 text-yellow-800 dark:bg-yellow-900/20 dark:border-yellow-800 dark:text-yellow-200',
                info: 'bg-blue-50 border-blue-200 text-blue-800 dark:bg-blue-900/20 dark:border-blue-800 dark:text-blue-200',
            },
            size: {
                sm: 'p-2 text-sm',
                base: 'p-4 text-base',
                lg: 'p-6 text-lg',
            },
        },
        defaultVariants: {
            variant: 'info',
            size: 'base',
        },
    }
);

export const blockquoteVariants = cva(
    `border-l-4 pl-4 italic`, // base
    {
        variants: {
            variant: {
                default: 'border-gray-300 text-gray-700 dark:border-gray-600 dark:text-gray-300',
                primary: 'border-blue-400 text-blue-700 dark:border-blue-600 dark:text-blue-300',
                success: 'border-green-400 text-green-700 dark:border-green-600 dark:text-green-300',
                warning: 'border-yellow-400 text-yellow-700 dark:border-yellow-600 dark:text-yellow-300',
                error: 'border-red-400 text-red-700 dark:border-red-600 dark:text-red-300',
            },
        },
        defaultVariants: {
            variant: 'default',
        },
    }
);

type TextVariantProps = VariantProps<typeof textVariants>;
type MessageVariantProps = VariantProps<typeof messageVariants>;
type BlockquoteVariantProps = VariantProps<typeof blockquoteVariants>;

export type TextProps = {
    class?: string;
} & TextVariantProps & QwikIntrinsicElements['p'];

export type MessageProps = {
    class?: string;
} & MessageVariantProps & QwikIntrinsicElements['div'];

export type BlockquoteProps = {
    class?: string;
} & BlockquoteVariantProps & QwikIntrinsicElements['blockquote'];