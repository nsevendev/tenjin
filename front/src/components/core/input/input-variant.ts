import {QwikIntrinsicElements, JSXOutput} from "@builder.io/qwik";

type InputVariantType = 'default' | 'outlined' | 'filled' | 'minimal' | 'defaultDisabled'
type InputSizeType = 'sm' | 'base' | 'lg'
type InputStateType = 'default' | 'error' | 'success' | 'disabled' | 'focus'

type InputStyleType = {
    container: {
        base: string;
        fullWidth: string;
    };
    label: {
        base: string;
        required: string;
        error: string;
        disabled: string;
    };
    wrapper: {
        base: string;
        variants: Record<InputVariantType, string>;
        sizes: Record<InputSizeType, string>;
        states: Record<InputStateType, string>;
        withLeftIcon: string;
        withRightIcon: string;
    };
    input: {
        base: string;
        sizes: Record<InputSizeType, string>;
        withLeftIcon: string;
        withRightIcon: string;
    };
    icon: {
        left: string;
        right: string;
    };
    message: {
        base: string;
        helper: string;
        error: string;
    };
}

export const inputStyles: InputStyleType = {
    container: {
        base: 'flex flex-col gap-1.5',
        fullWidth: 'w-full'
    },
    label: {
        base: 'text-sm font-medium text-gray-700 transition-colors duration-200',
        required: 'after:ml-1 after:text-red-500',
        error: 'text-red-600',
        disabled: 'text-gray-400 cursor-not-allowed'
    },
    wrapper: {
        base: 'relative flex items-center transition-all duration-200',
        variants: {
            default: 'border border-gray-300 rounded-lg bg-white hover:border-gray-400 focus-within:border-blue-500 focus-within:hover:border-blue-500',
            defaultDisabled: 'border border-gray-300 rounded-lg bg-white',
            outlined: 'border-2 border-gray-300 rounded-lg bg-transparent hover:border-gray-400 focus-within:border-blue-500 focus-within:hover:border-blue-500',
            filled: 'border-transparent rounded-lg bg-gray-100 focus-within:bg-white focus-within:border-1 focus-within:border-blue-500',
            minimal: 'border-0 border-b-2 border-gray-300 rounded-none bg-transparent hover:border-gray-400 focus-within:border-blue-500 focus-within:hover:border-blue-500'
        },
        sizes: {
            sm: 'h-9',
            base: 'h-11',
            lg: 'h-13'
        },
        states: {
            default: '',
            error: 'hover:placeholder-gray-500 border-red-500 hover:border-red-500 focus-within:border-red-500 focus-within:ring-red-500/20',
            success: 'hover:placeholder-gray-500 border-green-500 hover:border-green-500 focus-within:border-green-500 focus-within:ring-green-500/20',
            disabled: 'bg-gray-50 border-gray-200 cursor-not-allowed opacity-50',
            focus: 'hover:placeholder-gray-500 border-blue-500'
        },
        withLeftIcon: 'pl-10',
        withRightIcon: 'pr-10'
    },
    input: {
        base: 'w-full bg-transparent border-0 outline-none placeholder-gray-300 hover:placeholder-gray-400 disabled:hover:border-gray-300 disabled:cursor-not-allowed disabled:hover:placeholder-gray-300',
        sizes: {
            sm: 'px-3 py-2 text-sm',
            base: 'px-4 py-2.5 text-base',
            lg: 'px-5 py-3 text-lg'
        },
        withLeftIcon: 'pl-2',
        withRightIcon: 'pr-2'
    },
    icon: {
        left: 'absolute left-3 flex items-center justify-center text-gray-400 pointer-events-none',
        right: 'absolute right-3 flex items-center justify-center text-gray-400'
    },
    message: {
        base: 'text-xs mt-1 transition-colors duration-200',
        helper: 'text-gray-500',
        error: 'text-red-600'
    }
};

export type InputPropsType = {
    variant?: InputVariantType;
    size?: InputSizeType;
    state?: InputStateType;
    fullWidth?: boolean;
    label?: string;
    helper?: string;
    error?: string;
    required?: boolean;
    leftIcon?: JSXOutput;
    rightIcon?: JSXOutput;
} & Omit<QwikIntrinsicElements['input'], 'size'>;
