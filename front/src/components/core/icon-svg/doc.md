# IconSvgBase - Documentation

## Utilisation

Le composant `IconSvgBase` est un composant de base pour créer des icônes SVG avec des props standardisées.

### Créer une nouvelle icône

```tsx
import {component$} from "@builder.io/qwik";
import {IconProps} from "~/components/core/icon-svg/icon-svg-type";
import {IconSvgBase} from "~/components/core/icon-svg/icon-svg-base";

export const MonIcone = component$<IconProps>((props) => {
    return (
        <IconSvgBase {...props}>
            {/* Contenu SVG uniquement (path, circle, etc.) */}
            <path d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"/>
        </IconSvgBase>
    )
})
```

### Utilisation dans une page

```tsx
import { LockCloseIcon } from "~/components/core/icon-svg/lock/lock-close";

export default component$(() => {
    return (
        <div>
            <button onClick$={() => console.log('Clicked!')}>
                <LockCloseIcon size="lg" color="blue" />
                Verrouiller
            </button>
        </div>
    )
})
```

### Props disponibles

- `size` : Taille de l'icône (xs, sm, base, lg, xl ou nombre)
- `color` : Couleur du stroke (currentColor, black, white, etc.)
- `strokeWidth` : Épaisseur du trait (thin, normal, thick, extraThick ou nombre)
- `fill` : Couleur de remplissage (none par défaut)
- `onClick$` : Gestionnaire de clic
- `class` : Classes CSS supplémentaires

### Important

⚠️ **Ne pas** inclure la balise `<svg>` dans le Slot. Passez seulement le contenu (path, circle, etc.).

✅ **Correct :**
```tsx
<IconSvgBase {...props}>
    <path d="..."/>
</IconSvgBase>
```

❌ **Incorrect :**
```tsx
<IconSvgBase {...props}>
    <svg>
        <path d="..."/>
    </svg>
</IconSvgBase>
```