import {QwikIntrinsicElements} from "@builder.io/qwik";
import {cva, VariantProps} from "class-variance-authority";

export const cardVariants = cva(
    'rounded-lg bg-white shadow-sm w-fit', // base
    {
        variants: {
            containPosition: {
                default: '',
                centerColumn: 'flex flex-col items-center justify-center',
                center: 'flex items-center justify-center',
                left: 'flex flex-col items-start justify-center',
            },
            size: {
                small: 'p-4',
                medium: 'p-6',
                large: 'p-8',
                expanded: 'p-10',
                extraLarge: 'p-12',
                long: 'px-20 py-12',
                tall: 'px-12 py-20',
            },
        },
        defaultVariants: {
            containPosition: 'default',
            size: 'medium',
        },
        // optionnel: variants composés
        compoundVariants: [],
    }
);

type CardVariantProps = VariantProps<typeof cardVariants>;
export type CardTag = 'div' | 'article' | 'section';

export type CardProps<TCard extends CardTag> =
    {
        as?: TCard;
        class?: string;
    } & CardVariantProps
    & QwikIntrinsicElements[TCard];
