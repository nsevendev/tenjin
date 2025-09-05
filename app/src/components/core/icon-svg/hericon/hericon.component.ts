import { Component, computed, input } from '@angular/core';
import { ICONS } from './hericon.registry';
import {ICON_COLORS, ICON_SIZES, ICON_STROKES, IconColorKey, IconName, IconSizeKey, IconStrokeKey, IconVariant} from './hericon.type';

@Component({
  selector: 'c-h-icon',
  template: `
    <svg
      [attr.viewBox]="viewBox()"
      xmlns="http://www.w3.org/2000/svg"
      [attr.fill]="fillAttr()"
      [attr.stroke]="strokeColor()"
      [attr.width]="sizePx()"
      [attr.height]="sizePx()"
      [attr.stroke-width]="strokeWidthPx()"
      [attr.aria-label]="ariaLabel() || null"
      role="img"
      class="inline-block"
    >
      @for (d of paths(); track $index) {
        <path
          [attr.d]="d"
          stroke-linecap="round"
          stroke-linejoin="round">
        </path>
      }
    </svg>
  `
})
export class HericonComponent {
  // API
  readonly name       = input.required<IconName>();             // 'calendar', 'user', ...
  readonly variant    = input<IconVariant>('outline');        // 'outline' | 'solid'
  readonly size       = input<IconSizeKey | number>('lg');    // key ou nombre
  readonly color      = input<IconColorKey>('currentColor');  // token ou hex
  readonly fill       = input<IconColorKey>('none');          // utile pour 'solid'
  readonly stroke     = input<IconColorKey>('currentColor');  // contour
  readonly strokeWidth= input<IconStrokeKey | number>('normal');
  readonly ariaLabel  = input<string>('');

  // Sélection de l’icône
  private readonly icon = computed(() => {
    const key = this.name();
    const def = ICONS[key];

    // si pas trouvé, on retourne un carré vide
    if (!def) {
      if (ngDevMode) console.warn(`[c-h-icon] Icon "${key}" introuvable.`);
      return { variant: 'outline', paths: ['M4 4h16v16H4z'], viewBox: '0 0 24 24', filled: false };
    }

    return def;
  });

  readonly viewBox = computed(() => this.icon().viewBox ?? '0 0 24 24');
  readonly paths = computed(() => this.icon().paths);
  readonly sizePx = computed(() => {
    const s = this.size();
    return typeof s === 'number' ? s : ICON_SIZES[s] ?? ICON_SIZES.lg;
  });
  readonly strokeWidthPx = computed(() => {
    const sw = this.strokeWidth();
    return typeof sw === 'number' ? sw : (ICON_STROKES[sw] ?? ICON_STROKES.normal);
  });
  readonly strokeColor = computed(() => {
    const v = this.stroke();
    // si string = clé, sinon c’est déjà une valeur
    return (v in ICON_COLORS) ? (ICON_COLORS as any)[v] : v;
  });
  readonly fillAttr = computed(() => {
    // solid => si fill non forcé, on met currentColor
    const f = this.fill();
    if (f && f !== 'none') {
      return (f in ICON_COLORS) ? (ICON_COLORS as any)[f] : f;
    }
    return this.icon().variant === 'solid' ? 'currentColor' : 'none';
  });
}
