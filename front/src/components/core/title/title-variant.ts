import { QwikIntrinsicElements } from "@builder.io/qwik";
import {cva, VariantProps} from "class-variance-authority";

export const titleVariants = cva(
    '', // base
    {
        variants: {
            weight: {
                light:     'font-light',
                normal:    'font-normal',
                semibold:  'font-semibold',
                bold:      'font-bold',
                extrabold: 'font-extrabold',
            },
            size: {
                xs:  'text-xs',
                sm:  'text-sm',
                base:'text-base',
                lg:  'text-lg',
                xl:  'text-xl',
                '2xl':'text-2xl',
                '3xl':'text-3xl',
                '4xl':'text-4xl',
                '5xl':'text-5xl',
                '6xl':'text-6xl',
            },
        },
        defaultVariants: {
            weight: 'normal',
            size: 'base',
        },
        // optionnel: variants composés
        compoundVariants: [],
    }
);

type TitleVariantProps = VariantProps<typeof titleVariants>;
export type TitleTag = 'h1' | 'h2' | 'h3' | 'h4' | 'h5' | 'h6';

export type TitleProps<TTag extends TitleTag> =
    {
        as?: TTag;
        class?: string;
    } & TitleVariantProps
    & QwikIntrinsicElements[TTag];
