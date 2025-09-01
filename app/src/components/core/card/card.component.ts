import {Component, computed, input} from '@angular/core';
import {NgTemplateOutlet} from '@angular/common';
import {cn} from '../../../utils/class-name/cn';
import {CardTag, CardVariantProps, cardVariants} from './card.variant';

type VV<K extends keyof CardVariantProps> = NonNullable<CardVariantProps[K]>;

@Component({
  selector: 'c-card',
  template: `
    <ng-template #content>
      <ng-content></ng-content>
    </ng-template>

    @switch (tag()) {
      @case ('article') {
        <article [class]="classes()">
          <ng-container [ngTemplateOutlet]="content"/>
        </article>
      }
      @case ('section') {
        <section [class]="classes()">
          <ng-container [ngTemplateOutlet]="content"/>
        </section>
      }
      @default {
        <div [class]="classes()">
          <ng-container [ngTemplateOutlet]="content"/>
        </div>
      }
    }
  `,
  imports: [
    NgTemplateOutlet
  ]
})
export class CardComponent {
  // variants
  readonly size = input<VV<'size'>>('medium');

  // props
  readonly as = input<CardTag>('div');
  readonly className = input<string>('', { alias: 'class' });

  readonly tag = computed(() => this.as() ?? 'div');

  readonly classes = computed(() =>
    cn(
      cardVariants({
        size: this.size(),
      }),
      this.className()
    )
  );
}
