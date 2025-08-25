import { cva, type VariantProps } from 'class-variance-authority';
import {JSXNode, QwikIntrinsicElements} from "@builder.io/qwik";

/** Container */
export const containerVariants = cva('flex flex-col gap-1.5', {
    variants: {
        fullWidth: { true: 'w-full', false: null },
    },
    defaultVariants: { fullWidth: false },
});

/** Label */
export const labelVariants = cva(
    'text-sm font-medium text-gray-700 transition-colors duration-200',
    {
        variants: {
            state: {
                default: null,
                error: 'text-red-600',
                disabled: 'text-gray-400 cursor-not-allowed',
            },
        },
        defaultVariants: { state: 'default' },
    }
);

/** Wrapper autour de l’input (bordures, bg, icons, tailles, états) */
export const wrapperVariants = cva('relative flex items-center transition-all duration-200', {
    variants: {
        variant: {
            default:
                'border border-gray-300 rounded-lg bg-white hover:border-gray-400 focus-within:border-blue-500 focus-within:hover:border-blue-500',
            defaultDisabled: 'border border-gray-300 rounded-lg bg-white',
            outlined:
                'border-2 border-gray-300 rounded-lg bg-transparent hover:border-gray-400 focus-within:border-blue-500 focus-within:hover:border-blue-500',
            filled:
                'border-transparent rounded-lg bg-gray-100 focus-within:bg-white focus-within:border-1 focus-within:border-blue-500',
            minimal:
                'border-0 border-b-2 border-gray-300 rounded-none bg-transparent hover:border-gray-400 focus-within:border-blue-500 focus-within:hover:border-blue-500',
        },
        size: {
            sm: 'h-9',
            base: 'h-11',
            lg: 'h-13',
        },
        state: {
            default: null,
            error:
                'hover:placeholder-gray-500 border-red-500 hover:border-red-500 focus-within:border-red-500 focus-within:ring-red-500/20',
            success:
                'hover:placeholder-gray-500 border-green-500 hover:border-green-500 focus-within:border-green-500 focus-within:ring-green-500/20',
            disabled: 'bg-gray-50 border-gray-200 cursor-not-allowed opacity-50',
            focus: 'hover:placeholder-gray-500 border-blue-500',
        },
        leftIcon: { true: 'pl-10', false: null },
        rightIcon: { true: 'pr-10', false: null },
    },
    compoundVariants: [
        // { state: 'disabled', variant: 'minimal', class: 'opacity-60' },
    ],
    defaultVariants: {
        variant: 'default',
        size: 'base',
        state: 'default',
        leftIcon: false,
        rightIcon: false,
    },
});

/** Input (champ lui-même) */
export const inputFieldVariants = cva(
    'w-full bg-transparent border-0 outline-none placeholder-gray-300 hover:placeholder-gray-400 disabled:hover:border-gray-300 disabled:cursor-not-allowed disabled:hover:placeholder-gray-300',
    {
        variants: {
            size: {
                sm: 'px-3 py-2 text-sm',
                base: 'px-4 py-2.5 text-base',
                lg: 'px-5 py-3 text-lg',
            },
            leftIcon: { true: 'pl-2', false: null },
            rightIcon: { true: 'pr-2', false: null },
        },
        defaultVariants: { size: 'base', leftIcon: false, rightIcon: false },
    }
);

/** Message helper/erreur */
export const messageVariants = cva('text-xs mt-1 transition-colors duration-200', {
    variants: {
        kind: {
            helper: 'text-gray-500',
            error: 'text-red-600',
        },
    },
    defaultVariants: { kind: 'helper' },
});

/** Types utilitaires */
export type ContainerVariantProps = VariantProps<typeof containerVariants>;
export type LabelVariantProps = VariantProps<typeof labelVariants>;
export type WrapperVariantProps = VariantProps<typeof wrapperVariants>;
export type InputFieldVariantProps = VariantProps<typeof inputFieldVariants>;
export type MessageVariantProps = VariantProps<typeof messageVariants>;

type InputState = NonNullable<WrapperVariantProps['state']>;
type InputVariant = NonNullable<WrapperVariantProps['variant']>;
type InputSize = NonNullable<WrapperVariantProps['size']>;

export type InputPropsType = {
    /** UI */
    variant?: InputVariant;        // 'default' | 'outlined' | 'filled' | 'minimal' | 'defaultDisabled'
    size?: InputSize;              // 'sm' | 'base' | 'lg'
    state?: InputState;            // 'default' | 'error' | 'success' | 'disabled' | 'focus'
    fullWidth?: ContainerVariantProps['fullWidth'];
    
    /** Label & messages */
    label?: string;
    helper?: string;
    error?: string;
    required?: boolean;
    
    /** Icons */
    leftIcon?: JSXNode;
    rightIcon?: JSXNode;
    
    /** Classes externes sur le wrapper */
    class?: string;
} & Omit<QwikIntrinsicElements['input'], 'size' | 'class'>;
