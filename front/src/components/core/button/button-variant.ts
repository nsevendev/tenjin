import {QwikIntrinsicElements} from "@builder.io/qwik";
import {cva, VariantProps} from "class-variance-authority";

export const buttonVariants = cva(
    `inline-flex items-center justify-center rounded-lg font-semibold
    cursor-pointer focus:outline-none disabled:bg-gray-400 disabled:hover:bg-gray-400
    disabled:cursor-not-allowed transition-all duration-150 disabled:active:scale-100
    active:scale-95 active:shadow-inner`, // base
    {
        variants: {
            variant: {
                primary: 'bg-blue-500 hover:bg-blue-600 focus:ring-blue-500 text-white',
                secondary: 'bg-gray-500 hover:bg-gray-600 focus:ring-gray-500 text-white',
                success: 'bg-green-500 hover:bg-green-600 focus:ring-green-500 text-white',
                error: 'bg-red-500 hover:bg-red-600 focus:ring-red-500 text-white',
                warning: 'bg-yellow-500 hover:bg-yellow-600 focus:ring-yellow-500 text-white',
                purple: 'bg-purple-500 hover:bg-purple-600 focus:ring-purple-500 text-white',
                pink: 'bg-pink-500 hover:bg-pink-600 focus:ring-pink-500 text-white',
                indigo: 'bg-indigo-500 hover:bg-indigo-600 focus:ring-indigo-500 text-white',
                primaryPastel: 'bg-blue-pastel-300 hover:bg-blue-pastel-600 focus:ring-blue-pastel-500 text-black hover:text-white',
                secondaryPastel: 'bg-gray-pastel-300 hover:bg-gray-pastel-600 focus:ring-gray-pastel-500 text-black hover:text-white',
                successPastel: 'bg-green-pastel-300 hover:bg-green-pastel-600 focus:ring-green-pastel-500 text-black hover:text-white',
                errorPastel: 'bg-red-pastel-300 hover:bg-red-pastel-600 focus:ring-red-pastel-500 text-black hover:text-white',
                warningPastel: 'bg-yellow-pastel-300 hover:bg-yellow-pastel-600 focus:ring-yellow-pastel-500 text-black hover:text-white',
                purplePastel: 'bg-purple-pastel-300 hover:bg-purple-pastel-600 focus:ring-purple-pastel-500 text-black hover:text-white',
                pinkPastel: 'bg-pink-pastel-300 hover:bg-pink-pastel-600 focus:ring-pink-pastel-500 text-black hover:text-white',
            },
            transform: {
                default: '',
                upper: 'uppercase',
            },
            size: {
                sm: 'px-4 py-2 text-sm',
                base: 'px-6 py-4 text-base',
            },
        },
        defaultVariants: {
            variant: 'primary',
            transform: 'default',
            size: 'base',
        },
        // optionnel: variants composés
        compoundVariants: [],
    }
);

type ButtonVariantProps = VariantProps<typeof buttonVariants>;

export type ButtonProps =
    {
        class?: string;
    } & ButtonVariantProps
    & QwikIntrinsicElements['button'];
