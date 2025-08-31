import {Component, computed, input, output, signal} from '@angular/core';
import {BadgeVariantProps, badgeVariants} from './badge.variant';
import {cn} from '../../../utils/class-name/cn';

type VV<K extends keyof BadgeVariantProps> = NonNullable<BadgeVariantProps[K]>;

@Component({
  selector: 'c-badge',
  template: `
    <div [class]="classes()" (click)="onClick($event)">
      <ng-content></ng-content>
    </div>
  `,
})
export class BadgeComponent {
  // variants
  readonly color = input<VV<'color'>>('gray');
  readonly size = input<VV<'size'>>('base');

  // props
  readonly className = input<string>('', { alias: 'class' });
  readonly pressed = output<MouseEvent>();
  readonly hasHandler = input<boolean>(false);

  readonly classes = computed(() =>
    cn(badgeVariants({
      color: this.color(),
      size: this.size(),
      interactive: this._hasHandler()
    }), this.className())
  );

  private readonly _hasHandler = computed(() => this.hasHandler());

  onClick(e: MouseEvent) {
    this.pressed.emit(e);
  }
}
