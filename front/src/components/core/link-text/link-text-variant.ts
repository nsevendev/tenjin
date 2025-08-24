import {QwikIntrinsicElements} from "@builder.io/qwik";
import {cva, VariantProps} from "class-variance-authority";

export const linkVariants = cva(
    'inline cursor-pointer hover:underline', // base
    {
        variants: {
            color: {
                blue: 'text-blue-500 hover:text-blue-700',
                gray: 'text-gray-500 hover:text-gray-700',
                red: 'text-red-500 hover:text-red-700',
                green: 'text-green-500 hover:text-green-700',
            },
            size: {
                sm: 'text-sm',
                base: 'text-base',
                lg: 'text-lg',
            },
        },
        defaultVariants: {
            color: 'blue',
            size: 'base',
        },
        // optionnel: variants composés
        compoundVariants: [],
    }
);

type LinkVariantProps = VariantProps<typeof linkVariants>;

export type LinkProps =
    {
        external?: boolean;
        class?: string;
    } & LinkVariantProps
    & QwikIntrinsicElements['a'];
