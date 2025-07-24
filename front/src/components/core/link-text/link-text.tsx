import {component$, Slot} from "@builder.io/qwik";
import {Link as QwikLink} from "@builder.io/qwik-city";
import {LinkTextPropsType, linkTextStyles} from "~/components/core/link-text/link-text-variant";
import {cn} from "~/utils/classe-name/cn";

export const LinkText = component$<LinkTextPropsType>(({
    variant = 'blue',
    size = 'base',
    href,
    external = false,
    class: className,
    ...props
}) => {
    const linkClasses = cn(
        linkTextStyles.base,
        linkTextStyles.variants[variant],
        linkTextStyles.sizes[size],
        className
    );

    // Si c'est un lien externe ou commence par http/https, utiliser un <a> normal
    if (external || href.startsWith('http://') || href.startsWith('https://') || href.startsWith('mailto:') || href.startsWith('tel:')) {
        return (
            <a 
                {...props} 
                href={href} 
                class={linkClasses}
                target={external ? '_blank' : undefined}
                rel={external ? 'noopener noreferrer' : undefined}
            >
                <Slot/>
            </a>
        );
    }

    // Sinon, utiliser le Link de Qwik pour la navigation interne
    return (
        <QwikLink 
            {...props} 
            href={href} 
            class={linkClasses}
        >
            <Slot/>
        </QwikLink>
    );
});
