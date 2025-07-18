import {component$, Slot} from "@builder.io/qwik";
import {ButtonPropsType, buttonStyles} from "~/components/core/button/button-variant";
import {cn} from "~/utils/classe-name/cn";

export const Button = component$<ButtonPropsType>(({
    variant = 'primary',
    size = 'base',
    uppercase = false,
    class: className,
    ...props
}) => {
    return (
        <button {...props} class={cn(
                buttonStyles.base,
                buttonStyles.variants[variant],
                buttonStyles.sizes[size],
                uppercase && 'uppercase',
                className
            )}
        >
            <Slot/>
        </button>
    )
})
