import {component$, Slot} from "@builder.io/qwik";
import {BadgePropsType, badgeStyles} from "~/components/core/badge/badge-variant";
import {cn} from "~/utils/classe-name/cn";

export const Badge = component$<BadgePropsType>(({
    variant = 'blue',
    size = 'base',
    class: className,
    ...props
}) => {
    return (
        <div {...props} class={cn(
                badgeStyles.base,
                badgeStyles.variants[variant],
                badgeStyles.sizes[size],
                className
            )}
        >
            <Slot/>
        </div>
    )
})