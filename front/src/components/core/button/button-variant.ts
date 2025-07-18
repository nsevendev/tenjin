import {QwikIntrinsicElements} from "@builder.io/qwik";

type ButtonSizeType = 'sm' | 'base'
type ButtonVariantType = 'primary' | 'secondary' | 'success' | 'error' | 'warning' | 'purple' | 'pink' | 'indigo'
type ButtonStyleType = {
    base: string;
    variants: Record<ButtonVariantType, string>;
    sizes: Record<ButtonSizeType, string>;
}

export const buttonStyles: ButtonStyleType = {
    base: `inline-flex items-center justify-center rounded-lg font-semibold
    cursor-pointer focus:outline-none disabled:bg-gray-400 disabled:hover:bg-gray-400
    disabled:cursor-not-allowed transition-all duration-150 disabled:active:scale-100
    active:scale-95 active:shadow-inner`,
    variants: {
        primary: 'bg-blue-500 hover:bg-blue-600 focus:ring-blue-500 text-white',
        secondary: 'bg-gray-500 hover:bg-gray-600 focus:ring-gray-500 text-white',
        success: 'bg-green-500 hover:bg-green-600 focus:ring-green-500 text-white',
        error: 'bg-red-500 hover:bg-red-600 focus:ring-red-500 text-white',
        warning: 'bg-yellow-500 hover:bg-yellow-600 focus:ring-yellow-500 text-white',
        purple: 'bg-purple-500 hover:bg-purple-600 focus:ring-purple-500 text-white',
        pink: 'bg-pink-500 hover:bg-pink-600 focus:ring-pink-500 text-white',
        indigo: 'bg-indigo-500 hover:bg-indigo-600 focus:ring-indigo-500 text-white',
    },
    sizes: {
        sm: 'px-4 py-2 text-sm',
        base: 'px-6 py-4 text-base',
    }
};

export type ButtonPropsType = {
    variant?: ButtonVariantType;
    size?: ButtonSizeType;
    uppercase?: boolean;
} & QwikIntrinsicElements['button']
