import {QwikIntrinsicElements} from "@builder.io/qwik";

type BadgeVariantType = 'blue' | 'red' | 'green' | 'yellow' | 'purple' | 'orange' | 'gray' | 'indigo' | 'pink'
type BadgeSizeType = 'sm' | 'base'

type BadgeStyleType = {
    base: string;
    variants: Record<BadgeVariantType, string>;
    sizes: Record<BadgeSizeType, string>;
}

export const badgeStyles: BadgeStyleType = {
    base: 'inline-flex items-center justify-center rounded-full font-medium w-fit',
    variants: {
        blue: 'bg-blue-pastel-400 text-blue-900',
        red: 'bg-red-pastel-400 text-red-900',
        green: 'bg-green-pastel-400 text-green-900',
        yellow: 'bg-yellow-pastel-400 text-yellow-900',
        purple: 'bg-purple-pastel-400 text-purple-900',
        orange: 'bg-orange-pastel-400 text-orange-900',
        gray: 'bg-gray-pastel-400 text-gray-900',
        indigo: 'bg-indigo-400 text-indigo-900',
        pink: 'bg-pink-pastel-400 text-pink-900',
    },
    sizes: {
        sm: 'px-3 py-0.5 text-xs',
        base: 'px-4 py-1 text-sm',
    }
};

export type BadgePropsType = {
    variant?: BadgeVariantType;
    size?: BadgeSizeType;
} & QwikIntrinsicElements['div']
