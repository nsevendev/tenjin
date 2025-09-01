import {QwikIntrinsicElements} from "@builder.io/qwik";
import {cva, VariantProps} from "class-variance-authority";

export const containerVariants = cva(
    ``, // base
    {
        variants: {
            display: {
                flex: 'flex',
                'inline-flex': 'inline-flex',
                block: 'block',
                'inline-block': 'inline-block',
            },
            direction: {
                row: 'flex-row',
                'row-reverse': 'flex-row-reverse',
                col: 'flex-col',
                'col-reverse': 'flex-col-reverse',
            },
            justify: {
                start: 'justify-start',
                end: 'justify-end',
                center: 'justify-center',
                between: 'justify-between',
                around: 'justify-around',
                evenly: 'justify-evenly',
            },
            align: {
                start: 'items-start',
                end: 'items-end',
                center: 'items-center',
                baseline: 'items-baseline',
                stretch: 'items-stretch',
            },
            wrap: {
                nowrap: 'flex-nowrap',
                wrap: 'flex-wrap',
                'wrap-reverse': 'flex-wrap-reverse',
            },
            gap: {
                0: 'gap-0',
                1: 'gap-1',
                2: 'gap-2',
                3: 'gap-3',
                4: 'gap-4',
                5: 'gap-5',
                6: 'gap-6',
                8: 'gap-8',
                10: 'gap-10',
                12: 'gap-12',
                16: 'gap-16',
                20: 'gap-20',
            },
            padding: {
                0: 'p-0',
                1: 'p-1',
                2: 'p-2',
                3: 'p-3',
                4: 'p-4',
                5: 'p-5',
                6: 'p-6',
                8: 'p-8',
                10: 'p-10',
                12: 'p-12',
                16: 'p-16',
                20: 'p-20',
            },
            margin: {
                0: 'm-0',
                1: 'm-1',
                2: 'm-2',
                3: 'm-3',
                4: 'm-4',
                5: 'm-5',
                6: 'm-6',
                8: 'm-8',
                10: 'm-10',
                12: 'm-12',
                16: 'm-16',
                20: 'm-20',
                auto: 'm-auto',
            },
            width: {
                auto: 'w-auto',
                full: 'w-full',
                screen: 'w-screen',
                fit: 'w-fit',
                min: 'w-min',
                max: 'w-max',
            },
            height: {
                auto: 'h-auto',
                full: 'h-full',
                screen: 'h-screen',
                fit: 'h-fit',
                min: 'h-min',
                max: 'h-max',
            },
        },
        defaultVariants: {
            display: 'flex',
            direction: 'row',
            justify: 'start',
            align: 'start',
            wrap: 'nowrap',
            gap: 0,
            padding: 0,
            margin: 0,
            width: 'auto',
            height: 'auto',
        },
    }
);

type ContainerVariantProps = VariantProps<typeof containerVariants>;

export type ContainerProps = {
    class?: string;
} & ContainerVariantProps & QwikIntrinsicElements['div'];