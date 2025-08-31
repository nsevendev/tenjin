import {Component, computed, input} from '@angular/core';
import {TitleTag, TitleVariantProps, titleVariants} from './title.variant';
import {NgTemplateOutlet} from '@angular/common';
import {cn} from '../../../utils/class-name/cn';

type VV<K extends keyof TitleVariantProps> = NonNullable<TitleVariantProps[K]>;

@Component({
  selector: 'c-title',
  template: `
    <ng-template #content>
      <ng-content/>
    </ng-template>

    @switch (tag()) {
      @case ('h1') {
        <h1 [class]="classes()">
          <ng-container [ngTemplateOutlet]="content"/>
        </h1>
      }
      @case ('h2') {
        <h2 [class]="classes()">
          <ng-container [ngTemplateOutlet]="content"/>
        </h2>
      }
      @case ('h3') {
        <h3 [class]="classes()">
          <ng-container [ngTemplateOutlet]="content"/>
        </h3>
      }
      @case ('h4') {
        <h4 [class]="classes()">
          <ng-container [ngTemplateOutlet]="content"/>
        </h4>
      }
      @case ('h5') {
        <h5 [class]="classes()">
          <ng-container [ngTemplateOutlet]="content"/>
        </h5>
      }
      @default {
        <h6 [class]="classes()">
          <ng-container [ngTemplateOutlet]="content"/>
        </h6>
      }
    }
  `,
  imports: [
    NgTemplateOutlet
  ]
})
export class TitleComponent {
  // props
  readonly as = input<TitleTag>('h1');
  readonly size = input<VV<'size'>>('base');
  readonly weight = input<VV<'weight'>>('normal');
  readonly className = input<string>('', { alias: 'class' });

  readonly tag = computed<TitleTag>(() => this.as() ?? 'h1');

  readonly classes = computed(() =>
    cn(titleVariants({ size: this.size(), weight: this.weight() }), this.className())
  );
}
