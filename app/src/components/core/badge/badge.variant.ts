import { cva, type VariantProps } from 'class-variance-authority';

export const badgeVariants = cva(
  'inline-flex items-center justify-center rounded-full font-medium w-fit',
  {
    variants: {
      color: {
        blue:   'bg-blue-300 text-blue-900',
        red:    'bg-red-300 text-red-900',
        green:  'bg-green-300 text-green-900',
        yellow: 'bg-yellow-300 text-yellow-900',
        purple: 'bg-purple-300 text-purple-900',
        orange: 'bg-orange-300 text-orange-900',
        gray:   'bg-gray-300 text-gray-900',
        indigo: 'bg-indigo-300 text-indigo-900',
        pink:   'bg-pink-300 text-pink-900',
      },
      size: {
        sm:   'px-2.5 py-0.5 text-[11px] md:text-xs',
        base: 'px-3 py-1 text-xs md:px-4 md:py-1 md:text-sm',
      },
      interactive: {
        true:  'cursor-pointer',
        false: '',
      },
    },
    defaultVariants: {
      color: 'gray',
      size: 'base',
      interactive: false,
    },
    compoundVariants: [
      { color: 'blue',   interactive: true, class: 'hover:bg-blue-200' },
      { color: 'red',    interactive: true, class: 'hover:bg-red-200' },
      { color: 'green',  interactive: true, class: 'hover:bg-green-200' },
      { color: 'yellow', interactive: true, class: 'hover:bg-yellow-200' },
      { color: 'purple', interactive: true, class: 'hover:bg-purple-200' },
      { color: 'orange', interactive: true, class: 'hover:bg-orange-200' },
      { color: 'gray',   interactive: true, class: 'hover:bg-gray-200' },
      { color: 'indigo', interactive: true, class: 'hover:bg-indigo-200' },
      { color: 'pink',   interactive: true, class: 'hover:bg-pink-200' },
    ],
  }
);

export type BadgeVariantProps = VariantProps<typeof badgeVariants>;
