import {component$, Slot} from "@builder.io/qwik";
import {TitleProps, TitleStyles, TitleTagType} from "~/components/core/title/title-variant";

export const Title = component$<TitleProps<TitleTagType>>(({
    tag = 'h1',
    variant = 'black',
    size = 'base',
    className = '',
    ...props
}) => {
    const classes = [
        TitleStyles.base,
        TitleStyles.variants[variant],
        TitleStyles.sizes[size],
        className
    ].filter(Boolean).join(' ');
    
    // Création dynamique de l'élément selon le tag
    const TagComponent = tag as any;
    
    return (
        <TagComponent class={classes} {...props}>
            <Slot />
        </TagComponent>
    );
});

// Composants pré-configurés pour des cas d'usage spécifiques
export const PageTitle = component$<Omit<TitleProps<'h1'>, 'tag' | 'size'>>(({
    variant = 'black',
    ...props
}) => (
    <Title tag="h1" size="3xl" variant={variant} {...props}>
        <Slot />
    </Title>
));

export const SectionTitle = component$<Omit<TitleProps<'h2'>, 'tag' | 'size'>>(({
    variant = 'black',
    ...props
}) => (
    <Title tag="h2" size="2xl" variant={variant} {...props}>
        <Slot />
    </Title>
));

export const SubsectionTitle = component$<Omit<TitleProps<'h3'>, 'tag' | 'size'>>(({
    variant = 'black',
    ...props
}) => (
    <Title tag="h3" size="xl" variant={variant} {...props}>
        <Slot />
    </Title>
));

export const CardTitle = component$<Omit<TitleProps<'h4'>, 'tag' | 'size'>>(({
    variant = 'black',
    ...props
}) => (
    <Title tag="h4" size="lg" variant={variant} {...props}>
        <Slot />
    </Title>
));

export const SmallTitle = component$<Omit<TitleProps<'h5'>, 'tag' | 'size'>>(({
    variant = 'gray',
    ...props
}) => (
    <Title tag="h5" size="base" variant={variant} {...props}>
        <Slot />
    </Title>
));

