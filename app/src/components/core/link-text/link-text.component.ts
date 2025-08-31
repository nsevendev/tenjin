import {Component, computed, input} from '@angular/core';
import {RouterLink} from '@angular/router';
import {cn} from '../../../utils/class-name/cn';
import {LinkTextVariantProps, linkTextVariants} from './link-text.variant';
import {NgTemplateOutlet} from '@angular/common';

type VV<K extends keyof LinkTextVariantProps> = NonNullable<LinkTextVariantProps[K]>;

@Component({
  selector: 'c-link-text',
  imports: [
    RouterLink,
    NgTemplateOutlet
  ],
  template: `
    <ng-template #content>
      <ng-content/>
    </ng-template>

    @if (isExternal()) {
      <a [href]="url()" rel="noopener noreferrer" target="_blank" (click)="handleClick($event)" [class]="classes()">
        <ng-container [ngTemplateOutlet]="content"/>
      </a>
    } @else {
      <a [routerLink]="url()" (click)="handleClick($event)" [class]="classes()">
        <ng-container [ngTemplateOutlet]="content"/>
      </a>
    }
  `,

})
export class LinkTextComponent {
  // variants
  readonly color = input<VV<'color'>>('blue');
  readonly size = input<VV<'size'>>('base');
  readonly underline = input<VV<'underline'>>('hover');

  // props
  readonly url = input.required<string>();
  readonly className = input<string>('', { alias: 'class' });
  readonly external = input<boolean>(false);

  readonly classes = computed(() =>
    cn(
      linkTextVariants({
        color: this.color(),
        size: this.size(),
        underline: this.underline()
      }),
      this.className()
    )
  );

  readonly isExternal = computed(() =>{
    return this.external() || this.isExternalUrl();
  })

  // au changement d'url, on vérifie si c'est une url externe
  private readonly isExternalUrl = computed(() => {
    const url = this.url();
    return url?.startsWith('http://') ||
      url?.startsWith('https://') ||
      url?.startsWith('mailto:') ||
      url?.startsWith('tel:');
  });

  handleClick(event: Event) {
    if (!this.url()) {
      event.preventDefault();
      return;
    }
  }
}
