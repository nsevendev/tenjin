import {component$, Slot} from "@builder.io/qwik";
import {TitleProps, TitleTag, titleVariants} from "~/components/core/title/title-variant";
import {cn} from "~/utils/classe-name/cn";

export const Title = component$<TitleProps<TitleTag>>(({
    as = 'h1',
    size,
    weight,
    class: className,
    ...props
}) => {
    // Création dynamique de l'élément selon le tag
    const TagComponent = as as any;
    
    return (
        <TagComponent class={cn(titleVariants({size, weight}), className)} {...props}>
            <Slot />
        </TagComponent>
    );
});

// ----------------------------- Composants pré-configurés pour des cas d'usage spécifiques ----------------------------- //

export const PageTitle = component$<Omit<TitleProps<'h1'>, 'as' | 'color' | 'size' | 'weight'>>((props) => (
    <Title as="h1" size="3xl" weight="semibold" {...props}>
        <Slot />
    </Title>
));

export const PageSubTitle = component$<Omit<TitleProps<'h1'>, 'as' | 'color' | 'size' | 'weight'>>((props) => (
    <Title as="h2" size="2xl" weight="normal" {...props}>
        <Slot />
    </Title>
));

export const SectionTitle = component$<Omit<TitleProps<'h2'>, 'as' | 'color' | 'size' | 'weight'>>((props) => (
    <Title as="h2" size="2xl" weight="semibold" {...props}>
        <Slot />
    </Title>
));

export const SectionSubTitle = component$<Omit<TitleProps<'h3'>, 'as' | 'color' | 'size' | 'weight'>>((props) => (
    <Title as="h3" size="xl" weight="normal" {...props}>
        <Slot />
    </Title>
));

export const CardTitle = component$<Omit<TitleProps<'h4'>, 'as' | 'color' | 'size' | 'weight'>>((props) => (
    <Title as="h4" size="lg" weight="semibold" {...props}>
        <Slot />
    </Title>
));

export const CardSubTitle = component$<Omit<TitleProps<'h4'>, 'as' | 'color' | 'size' | 'weight'>>((props) => (
    <Title as="h5" size="base" weight="light" {...props}>
        <Slot />
    </Title>
));

export const SmallTitle = component$<Omit<TitleProps<'h6'>, 'as' | 'color' | 'size' | 'weight'>>((props) => (
    <Title as="h6" size="sm" weight="light" {...props}>
        <Slot />
    </Title>
));

