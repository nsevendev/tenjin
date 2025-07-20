import {ClassList, Signal} from "@builder.io/qwik";

export type VariantConfig<T extends string> = Record<string, Record<T, string>>;

// Type générique pour les props
type ComponentProps<T extends string> = {
    variant?: T;
    class?: string | ClassList | Signal<ClassList> | undefined;
};

// generation de classes pour les composants
export function generateClasses<T extends string>(
    baseClasses: string,
    variantConfig: VariantConfig<T>,
    defaultVariant: T,
    props: ComponentProps<T>
): string {
    const {
        variant = defaultVariant,
        class: className = ''
    } = props;
    
    // Récupère toutes les classes des variants
    const variantClasses = Object.values(variantConfig)
        .map(variants => variants[variant])
        .filter(Boolean);
    
    return [
        baseClasses,
        ...variantClasses,
        className
    ].filter(Boolean).join(' ');
}
