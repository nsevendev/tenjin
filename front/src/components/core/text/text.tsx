import {component$, Slot} from "@builder.io/qwik";
import {TextProps, MessageProps, BlockquoteProps, textVariants, messageVariants, blockquoteVariants} from "~/components/core/text/text-variant";
import {cn} from "~/utils/classe-name/cn";


export const Text = component$<TextProps>(({
    variant = 'paragraph',
    size = 'base',
    weight = 'normal',
    align = 'left',
    color = 'default',
    class: className,
    ...props
}) => {
    return (
        <p {...props} class={cn(textVariants({variant, size, weight, align, color}), className)}>
            <Slot/>
        </p>
    )
});

export const Message = component$<MessageProps>(({
    variant = 'info',
    size = 'base',
    class: className,
    ...props
}) => {
    return (
        <div {...props} class={cn(messageVariants({variant, size}), className)}>
            <Slot/>
        </div>
    )
});

export const Blockquote = component$<BlockquoteProps>(({
    variant = 'default',
    class: className,
    ...props
}) => {
    return (
        <blockquote {...props} class={cn(blockquoteVariants({variant}), className)}>
            <Slot/>
        </blockquote>
    )
});

// ----------------------------- Composants pré-configurés pour des cas d'usage spécifiques ----------------------------- //

// Composants de texte
export const Paragraph = component$<Omit<TextProps, 'variant'>>((props) => (
    <Text variant="paragraph" {...props}>
        <Slot/>
    </Text>
));

export const SmallText = component$<Omit<TextProps, 'variant' | 'size'>>((props) => (
    <Text variant="small" size="sm" {...props}>
        <Slot/>
    </Text>
));

export const Caption = component$<Omit<TextProps, 'variant' | 'size'>>((props) => (
    <Text variant="caption" size="xs" {...props}>
        <Slot/>
    </Text>
));

export const Lead = component$<Omit<TextProps, 'variant' | 'size'>>((props) => (
    <Text variant="lead" size="lg" {...props}>
        <Slot/>
    </Text>
));

export const MutedText = component$<Omit<TextProps, 'variant'>>((props) => (
    <Text variant="muted" {...props}>
        <Slot/>
    </Text>
));

export const CodeText = component$<Omit<TextProps, 'variant'>>((props) => (
    <Text variant="code" {...props}>
        <Slot/>
    </Text>
));

// Composants de message
export const SuccessMessage = component$<Omit<MessageProps, 'variant'>>((props) => (
    <Message variant="success" {...props}>
        <Slot/>
    </Message>
));

export const ErrorMessage = component$<Omit<MessageProps, 'variant'>>((props) => (
    <Message variant="error" {...props}>
        <Slot/>
    </Message>
));

export const WarningMessage = component$<Omit<MessageProps, 'variant'>>((props) => (
    <Message variant="warning" {...props}>
        <Slot/>
    </Message>
));

export const InfoMessage = component$<Omit<MessageProps, 'variant'>>((props) => (
    <Message variant="info" {...props}>
        <Slot/>
    </Message>
));

// Composants de citation
export const DefaultQuote = component$<Omit<BlockquoteProps, 'variant'>>((props) => (
    <Blockquote variant="default" {...props}>
        <Slot/>
    </Blockquote>
));

export const PrimaryQuote = component$<Omit<BlockquoteProps, 'variant'>>((props) => (
    <Blockquote variant="primary" {...props}>
        <Slot/>
    </Blockquote>
));

export const SuccessQuote = component$<Omit<BlockquoteProps, 'variant'>>((props) => (
    <Blockquote variant="success" {...props}>
        <Slot/>
    </Blockquote>
));

export const WarningQuote = component$<Omit<BlockquoteProps, 'variant'>>((props) => (
    <Blockquote variant="warning" {...props}>
        <Slot/>
    </Blockquote>
));

export const ErrorQuote = component$<Omit<BlockquoteProps, 'variant'>>((props) => (
    <Blockquote variant="error" {...props}>
        <Slot/>
    </Blockquote>
));
