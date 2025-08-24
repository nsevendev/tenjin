import {component$, Slot} from "@builder.io/qwik";
import {CardProps, CardTag, cardVariants} from "~/components/core/card/card-variant";
import {cn} from "~/utils/classe-name/cn";

export const Card = component$<CardProps<CardTag>>(({
    as = 'div',
    containPosition = 'default',
    size = 'medium',
    class: className,
    ...props
}) => {
    const TagCard = as as any;
    
    return (
        <TagCard {...props} class={cn(cardVariants({containPosition, size}), className)} {...props}>
            <Slot/>
        </TagCard>
    )
})

// ----------------------------- Composants pré-configurés pour des cas d'usage spécifiques ----------------------------- //

export const SectionCard =
    component$<Omit<CardProps<'section'>, 'as' | 'size'>>(({containPosition = 'left', ...props}) => (
    <Card as="section" containPosition={containPosition} size="extraLarge" {...props}>
        <Slot/>
    </Card>
))

export const ArticleCard =
    component$<Omit<CardProps<'article'>, 'as' | 'size'>>(({containPosition = 'center', ...props}) => (
    <Card as="article" containPosition={containPosition} size="expanded" {...props}>
        <Slot/>
    </Card>
));
