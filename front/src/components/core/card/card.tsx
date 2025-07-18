import {component$, Slot} from "@builder.io/qwik";
import {CardPropsType, cardStyles} from "~/components/core/card/card-variant";
import {cn} from "~/utils/classe-name/cn";

export const Card = component$<CardPropsType>(({
    variant = 'medium',
    containPosition = 'default',
    tag = 'div',
    class: className,
    ...props
}) => {
    const Tag = tag as any;
    
    return (
        <Tag {...props} class={cn(
                cardStyles.base,
                cardStyles.variants[variant],
                cardStyles.containPosition[containPosition],
                className
            )}
        >
            <Slot/>
        </Tag>
    )
})
