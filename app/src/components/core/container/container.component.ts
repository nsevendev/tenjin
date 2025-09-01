import {ContainerVariantProps, containerVariants} from './container.variant';
import {Component, computed, input} from '@angular/core';
import {cn} from '../../../utils/class-name/cn';

type VV<K extends keyof ContainerVariantProps> = NonNullable<ContainerVariantProps[K]>;

@Component({
  selector: 'c-container',
  template: `
    <div [class]="classes()">
      <ng-content></ng-content>
    </div>
  `,
})
export class ContainerComponent {
  // variants
  readonly display  = input<VV<'display'>>('flex');
  readonly direction= input<VV<'direction'>>('row');
  readonly justify  = input<VV<'justify'>>('start');
  readonly align    = input<VV<'align'>>('start');
  readonly wrap     = input<VV<'wrap'>>('nowrap');
  readonly gap      = input<VV<'gap'>>(0);
  readonly padding  = input<VV<'padding'>>(0);
  readonly margin   = input<VV<'margin'>>(0);
  readonly width    = input<VV<'width'>>('auto');
  readonly height   = input<VV<'height'>>('auto');

  // props
  readonly className = input<string>('', { alias: 'class' });

  readonly classes = computed(() =>
    cn(
      containerVariants({
        display: this.display(),
        direction: this.direction(),
        justify: this.justify(),
        align: this.align(),
        wrap: this.wrap(),
        gap: this.gap(),
        padding: this.padding(),
        margin: this.margin(),
        width: this.width(),
        height: this.height(),
      }),
      this.className()
    )
  );
}
