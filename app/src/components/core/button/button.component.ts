import {Component, computed, input, output} from '@angular/core';
import {ButtonVariantProps, buttonVariants} from './button.variant';
import {cn} from '../../../utils/class-name/cn';

type VV<K extends keyof ButtonVariantProps> = NonNullable<ButtonVariantProps[K]>;

@Component({
  selector: 'c-button',
  template: `
    <button
      [attr.type]="type()"
      [class]="classes()"
      [disabled]="disabled()"
      (click)="onClick($event)"
    >
      <ng-content></ng-content>
    </button>
  `,
})
export class ButtonComponent {
  // variant
  readonly variant   = input<VV<'variant'>>('primary');
  readonly size      = input<VV<'size'>>('base');
  readonly transform = input<VV<'transform'>>('default');

  // props
  readonly type      = input<'button' | 'submit' | 'reset'>('button');
  readonly disabled  = input<boolean>(false);
  readonly className = input<string>('', { alias: 'class' });
  readonly pressed = output<MouseEvent>()

  readonly classes = computed(() =>
    cn(
      buttonVariants({
        variant: this.variant(),
        size: this.size(),
        transform: this.transform(),
      }),
      this.className(),
    )
  );

  onClick(e: MouseEvent) {
    if (this.disabled()) {
      e.preventDefault();
      e.stopPropagation();
      return;
    }
    this.pressed.emit(e);
  }


}
