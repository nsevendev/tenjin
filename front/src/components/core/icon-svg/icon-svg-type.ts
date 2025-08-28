import type { QRL } from "@builder.io/qwik";

export type IconProps = {
    size?: IconSize | number | string;
    class?: string;
    color?: ColorSvg | string;
    strokeWidth?: StrokeWidths | number | string;
    fill?: ColorSvg | string;
    onClick$?: QRL<() => void>;
}

export const iconSizes = {
    xs: 12,
    sm: 16,
    base: 20,
    lg: 24,
    xl: 32
}

export type IconSize = keyof typeof iconSizes | number | string;

export const strokeWidths = {
    thin: 1,
    normal: 1.5,
    thick: 2,
    extraThick: 3
}

export type StrokeWidths = keyof typeof strokeWidths | number | string;

export const colorSvg = {
    currentColor: 'currentColor',
    none: 'none',
    transparent: 'transparent',
    black: '#000000',
    white: '#FFFFFF',
    gray: '#6B7280',
    red: '#EF4444',
    green: '#10B981',
    blue: '#3B82F6',
    yellow: '#F59E0B',
    purple: '#8B5CF6',
    pink: '#EC4899'
}

export type ColorSvg = keyof typeof colorSvg | string;
