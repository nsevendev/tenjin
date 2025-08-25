import {component$, Slot} from "@builder.io/qwik";
import {Link as QwikLink} from "@builder.io/qwik-city";
import {LinkProps, linkVariants} from "~/components/core/link-text/link-text-variant";
import {cn} from "~/utils/classe-name/cn";

export const LinkText = component$<LinkProps>(({
    color = 'blue',
    size = 'base',
    href,
    external = false,
    class: className,
    ...props
}) => {
    const linkClasses = cn(linkVariants({color, size}),className);

    // lien externe ou commence par http/https, utiliser un <a> normal
    if (external || href?.startsWith('http://') || href?.startsWith('https://') || href?.startsWith('mailto:') || href?.startsWith('tel:')) {
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

    // utilise le Link de Qwik pour la navigation interne
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
