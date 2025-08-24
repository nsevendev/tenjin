import {component$, Slot} from "@builder.io/qwik";
import {ButtonProps, buttonVariants} from "~/components/core/button/button-variant";
import {cn} from "~/utils/classe-name/cn";

export const Button = component$<ButtonProps>(({
    variant = 'primary',
    size = 'base',
    transform,
    class: className,
    ...props
}) => {
    return (
        <button {...props} class={cn(buttonVariants({variant, size, transform}) ,className)} {...props}>
            <Slot/>
        </button>
    )
})

// ----------------------------- Composants pré-configurés pour des cas d'usage spécifiques ----------------------------- //

export const PrimaryButton = component$<Omit<ButtonProps, 'variant' | 'size'>>((props) => (
    <Button variant="primary" size="base" {...props}>
        <Slot/>
    </Button>
));

export const PrimarySmallButton = component$<Omit<ButtonProps, 'variant' | 'size'>>((props) => (
    <Button variant="primary" size="sm" {...props}>
        <Slot/>
    </Button>
));

export const SecondaryButton = component$<Omit<ButtonProps, 'variant' | 'size'>>((props) => (
    <Button variant="secondary" size="base" {...props}>
        <Slot/>
    </Button>
));

export const SecondarySmallButton = component$<Omit<ButtonProps, 'variant' | 'size'>>((props) => (
    <Button variant="secondary" size="sm" {...props}>
        <Slot/>
    </Button>
));
