import {ICONS} from './hericon.registry';

export type IconVariant = 'outline' | 'solid';
export type IconSizeKey = 'xs' | 'sm' | 'base' | 'lg' | 'xl';
export type IconStrokeKey = 'thin' | 'normal' | 'thick' | 'extraThick';

export type IconName = keyof typeof ICONS;

export const ICON_SIZES: Record<IconSizeKey, number> = {
  xs: 12, sm: 16, base: 20, lg: 24, xl: 32,
};

export const ICON_STROKES: Record<IconStrokeKey, number> = {
  thin: 1, normal: 1.5, thick: 2, extraThick: 3,
};

export const ICON_COLORS = {
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
  pink: '#EC4899',
  primary: '#0089ff',
  secondary: '#6B7280',
  success: '#22c55e',
  error: '#ef4444',
  warning: '#eab308',
  info: '#3B82F6',
} as const;

export type IconColorKey = keyof typeof ICON_COLORS | string;

export type IconDef = {
  viewBox?: string;              // default "0 0 24 24"
  variant: IconVariant;
  paths: string[];               // <path d="...">
  filled?: boolean;              // utile pour gérer fill par défaut des solid
};

export type IconRegistry = Record<string, IconDef>;
