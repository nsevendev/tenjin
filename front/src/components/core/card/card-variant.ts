import {QwikIntrinsicElements} from "@builder.io/qwik";

type CardVariantType = 'small' | 'medium' | 'large' | 'expanded' | 'extraLarge' | 'long' | 'tall'
type CardContainPositionType = 'center' | 'default' | 'centerColumn';
type CardTagType = 'div' | 'article' | 'section';
type CardStyleType = {
    base: string;
    variants: Record<CardVariantType, string>;
    containPosition: Record<CardContainPositionType, string>;
}

export const cardStyles: CardStyleType = {
    base: 'rounded-lg bg-white shadow-sm w-fit',
    containPosition: {
        centerColumn: 'flex items-center justify-center flex-col',
        default: '',
        center: 'flex items-center justify-center',
    },
    variants: {
        small: 'p-4',
        medium: 'p-6',
        large: 'p-8',
        expanded: 'p-10',
        extraLarge: 'p-12',
        long: 'px-20 py-12',
        tall: 'px-12 py-20',
    }
};

export type CardPropsType<T extends CardTagType = 'div'> = {
    variant?: CardVariantType;
    containPosition?: CardContainPositionType;
    tag?: CardTagType;
} & QwikIntrinsicElements[T]
