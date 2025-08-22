import { QwikIntrinsicElements } from "@builder.io/qwik";

export type TitleVariantType = 'blue' | 'gray' | 'red' | 'green' | 'black';
export type TitleSizeType = 'xs' | 'sm' | 'base' | 'lg' | 'xl' | '2xl' | '3xl';
export type TitleTagType = 'h1' | 'h2' | 'h3' | 'h4' | 'h5' | 'h6';

export type TitleStyleType = {
    base: string;
    variants: Record<TitleVariantType, string>;
    sizes: Record<TitleSizeType, string>;
};

export const TitleStyles: TitleStyleType = {
    base: 'font-semibold leading-tight',
    variants: { // TODO : ajout mode dark
        blue: 'text-blue-600',
        gray: 'text-gray-600',
        red: 'text-red-600',
        green: 'text-green-600',
        black: 'text-gray-900',
    },
    sizes: {
        xs: 'text-xs',
        sm: 'text-sm',
        base: 'text-base',
        lg: 'text-lg',
        xl: 'text-xl',
        '2xl': 'text-2xl',
        '3xl': 'text-3xl',
    }
};

// Types conditionnels pour chaque balise
export type TitleProps<T extends TitleTagType = 'h1'> = {
    tag?: T;
    variant?: TitleVariantType;
    size?: TitleSizeType;
    className?: string;
} & (T extends 'h1' ? QwikIntrinsicElements['h1'] :
    T extends 'h2' ? QwikIntrinsicElements['h2'] :
        T extends 'h3' ? QwikIntrinsicElements['h3'] :
            T extends 'h4' ? QwikIntrinsicElements['h4'] :
                T extends 'h5' ? QwikIntrinsicElements['h5'] :
                    QwikIntrinsicElements['h6']);
