import {component$, Slot} from "@builder.io/qwik";
import {BadgeProps, badgeVariants} from "~/components/core/badge/badge-variant";
import {cn} from "~/utils/classe-name/cn";

export const Badge = component$<BadgeProps>(({
    color = 'gray',
    size = 'base',
    class: className,
    ...props
}) => {
    return (
        <div class={cn(badgeVariants({color, size}), className)}{...props}>
            <Slot/>
        </div>
    )
})
