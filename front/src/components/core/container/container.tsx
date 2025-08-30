import {component$, Slot} from "@builder.io/qwik";
import {ContainerProps, containerVariants} from "~/components/core/container/container-variant";
import {cn} from "~/utils/classe-name/cn";

export const Container = component$<ContainerProps>(({
    display = 'flex',
    direction = 'row',
    justify = 'start',
    align = 'start',
    wrap = 'nowrap',
    gap = 0,
    padding = 0,
    margin = 0,
    width = 'auto',
    height = 'auto',
    class: className,
    ...props
}) => {
    return (
        <div {...props} class={cn(containerVariants({display, direction, justify, align, wrap, gap, padding, margin, width, height}), className)}>
            <Slot/>
        </div>
    )
});

// ----------------------------- Composants pré-configurés pour des cas d'usage spécifiques ----------------------------- //

// Conteneurs flex row
export const FlexRow = component$<Omit<ContainerProps, 'display' | 'direction'>>((props) => (
    <Container display="flex" direction="row" {...props}>
        <Slot/>
    </Container>
));

export const FlexRowCenter = component$<Omit<ContainerProps, 'display' | 'direction' | 'justify' | 'align'>>((props) => (
    <Container display="flex" direction="row" justify="center" align="center" {...props}>
        <Slot/>
    </Container>
));

export const FlexRowBetween = component$<Omit<ContainerProps, 'display' | 'direction' | 'justify' | 'align'>>((props) => (
    <Container display="flex" direction="row" justify="between" align="center" {...props}>
        <Slot/>
    </Container>
));

export const FlexRowStart = component$<Omit<ContainerProps, 'display' | 'direction' | 'justify' | 'align'>>((props) => (
    <Container display="flex" direction="row" justify="start" align="center" {...props}>
        <Slot/>
    </Container>
));

export const FlexRowEnd = component$<Omit<ContainerProps, 'display' | 'direction' | 'justify' | 'align'>>((props) => (
    <Container display="flex" direction="row" justify="end" align="center" {...props}>
        <Slot/>
    </Container>
));

// Conteneurs flex column
export const FlexCol = component$<Omit<ContainerProps, 'display' | 'direction'>>((props) => (
    <Container display="flex" direction="col" {...props}>
        <Slot/>
    </Container>
));

export const FlexColCenter = component$<Omit<ContainerProps, 'display' | 'direction' | 'justify' | 'align'>>((props) => (
    <Container display="flex" direction="col" justify="center" align="center" {...props}>
        <Slot/>
    </Container>
));

export const FlexColBetween = component$<Omit<ContainerProps, 'display' | 'direction' | 'justify' | 'align'>>((props) => (
    <Container display="flex" direction="col" justify="between" align="center" {...props}>
        <Slot/>
    </Container>
));

export const FlexColStart = component$<Omit<ContainerProps, 'display' | 'direction' | 'justify' | 'align'>>((props) => (
    <Container display="flex" direction="col" justify="start" align="center" {...props}>
        <Slot/>
    </Container>
));

export const FlexColEnd = component$<Omit<ContainerProps, 'display' | 'direction' | 'justify' | 'align'>>((props) => (
    <Container display="flex" direction="col" justify="end" align="center" {...props}>
        <Slot/>
    </Container>
));

// Conteneurs avec espacement
export const FlexRowGap = component$<Omit<ContainerProps, 'display' | 'direction' | 'gap'>>((props) => (
    <Container display="flex" direction="row" gap={4} {...props}>
        <Slot/>
    </Container>
));

export const FlexColGap = component$<Omit<ContainerProps, 'display' | 'direction' | 'gap'>>((props) => (
    <Container display="flex" direction="col" gap={4} {...props}>
        <Slot/>
    </Container>
));

// Conteneurs avec padding
export const PaddedContainer = component$<Omit<ContainerProps, 'padding'>>((props) => (
    <Container padding={4} {...props}>
        <Slot/>
    </Container>
));

export const LargePaddedContainer = component$<Omit<ContainerProps, 'padding'>>((props) => (
    <Container padding={8} {...props}>
        <Slot/>
    </Container>
));

// Conteneurs pleine largeur/hauteur
export const FullWidthContainer = component$<Omit<ContainerProps, 'width'>>((props) => (
    <Container width="full" {...props}>
        <Slot/>
    </Container>
));

export const FullHeightContainer = component$<Omit<ContainerProps, 'height'>>((props) => (
    <Container height="full" {...props}>
        <Slot/>
    </Container>
));

export const FullContainer = component$<Omit<ContainerProps, 'width' | 'height'>>((props) => (
    <Container width="full" height="full" {...props}>
        <Slot/>
    </Container>
));

// Conteneurs centrés
export const CenteredContainer = component$<Omit<ContainerProps, 'display' | 'justify' | 'align'>>((props) => (
    <Container display="flex" justify="center" align="center" {...props}>
        <Slot/>
    </Container>
));

export const CenteredFullContainer = component$<Omit<ContainerProps, 'display' | 'justify' | 'align' | 'width' | 'height'>>((props) => (
    <Container display="flex" justify="center" align="center" width="full" height="full" {...props}>
        <Slot/>
    </Container>
));