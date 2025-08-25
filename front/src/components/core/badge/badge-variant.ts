import {cva, VariantProps} from "class-variance-authority";
import {QwikIntrinsicElements} from "@builder.io/qwik";

export const badgeVariants = cva(
    'inline-flex items-center justify-center rounded-full font-medium w-fit', // base
    {
        variants: {
            color: {
                blue: 'bg-blue-pastel-300 text-blue-900',
                red: 'bg-red-pastel-300 text-red-900',
                green: 'bg-green-pastel-300 text-green-900',
                yellow: 'bg-yellow-pastel-300 text-yellow-900',
                purple: 'bg-purple-pastel-300 text-purple-900',
                orange: 'bg-orange-pastel-300 text-orange-900',
                gray: 'bg-gray-pastel-300 text-gray-900',
                indigo: 'bg-indigo-300 text-indigo-900',
                pink: 'bg-pink-pastel-300 text-pink-900',
            },
            size: {
                sm: 'px-3 py-0.5 text-xs',
                base: 'px-4 py-1 text-sm',
            },
        },
        defaultVariants: {
            color: 'gray',
            size: 'base',
        },
        // optionnel: variants composés
        compoundVariants: [],
    }
);

type BadgeVariantProps = VariantProps<typeof badgeVariants>;

export type BadgeProps =
    {
        class?: string;
    } & BadgeVariantProps
    & QwikIntrinsicElements['div'];
