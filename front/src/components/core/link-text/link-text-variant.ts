import {QwikIntrinsicElements} from "@builder.io/qwik";

type LinkTextVariantType = 'blue' | 'gray' | 'red' | 'green'
type LinkTextSizeType = 'sm' | 'base' | 'lg'

type LinkTextStyleType = {
    base: string;
    variants: Record<LinkTextVariantType, string>;
    sizes: Record<LinkTextSizeType, string>;
}

export const linkTextStyles: LinkTextStyleType = {
    base: 'inline cursor-pointer hover:underline',
    variants: {
        blue: 'text-blue-500 hover:text-blue-700',
        gray: 'text-gray-500 hover:text-gray-700',
        red: 'text-red-500 hover:text-red-700',
        green: 'text-green-500 hover:text-green-700',
    },
    sizes: {
        sm: 'text-sm',
        base: 'text-base',
        lg: 'text-lg',
    }
};

export type LinkTextPropsType = {
    variant?: LinkTextVariantType;
    size?: LinkTextSizeType;
    href: string;
    external?: boolean;
} & Omit<QwikIntrinsicElements['a'], 'href'>
