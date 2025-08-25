import {component$} from "@builder.io/qwik";
import {colorSvg, IconProps, iconSizes, strokeWith} from "~/components/core/icon/icon-type";
import {cn} from "~/utils/classe-name/cn";

export const LockCloseIcon = component$<IconProps>(({
    class: className,
    size = iconSizes.lg,
    color = colorSvg.currentColor,
    strokeWidth = strokeWith.normal,
    ...props
}) => {
    return (
        <svg
            {...props}
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width={strokeWidth}
            stroke={color}
            width={size}
            height={size}
            class={cn(className)}
        >
            <path stroke-linecap="round" stroke-linejoin="round"
                  d="M16.5 10.5V6.75a4.5 4.5 0 1 0-9 0v3.75m-.75 11.25h10.5a2.25 2.25 0 0 0 2.25-2.25v-6.75a2.25 2.25 0 0 0-2.25-2.25H6.75a2.25 2.25 0 0 0-2.25 2.25v6.75a2.25 2.25 0 0 0 2.25 2.25Z"/>
        </svg>
    )
})
