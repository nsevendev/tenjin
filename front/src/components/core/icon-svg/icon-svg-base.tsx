import {component$, Slot} from "@builder.io/qwik";
import {colorSvg, IconProps} from "~/components/core/icon-svg/icon-svg-type";
import {getIconSvgProps} from "~/components/core/icon-svg/icon-svg-hook";
import { cn } from "~/utils/classe-name/cn";

export const IconSvgBase = component$<IconProps & {
    viewBox?: string;
    "aria-label"?: string;
}>(({
    class: className,
    fill = colorSvg.none,
    onClick$,
    viewBox = "0 0 24 24",
    "aria-label": ariaLabel,
    // Extraction des props pour le hook
    size,
    color,
    strokeWidth,
    ...restProps
}) => {
    // Passage des bonnes props au hook
    const { sizeValue, strokeValue, colorValue } = getIconSvgProps({
        size,
        color,
        strokeWidth
    });
    
    return (
        <svg
            {...restProps}
            xmlns="http://www.w3.org/2000/svg"
            fill={fill}
            viewBox={viewBox}
            stroke-width={strokeValue}
            stroke={colorValue}
            width={sizeValue}
            height={sizeValue}
            class={cn(
                "inline-block",
                onClick$ && "cursor-pointer",
                className
            )}
            onClick$={onClick$}
            aria-label={ariaLabel}
            role={ariaLabel ? "img" : undefined}
        >
            <Slot />
        </svg>
    );
});
