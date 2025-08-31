import {Component, computed, input, output} from '@angular/core';
import {ButtonComponent} from './button.component';
import {ButtonVariantProps, buttonVariants} from './button.variant';
import {cn} from '../../../utils/class-name/cn';

type VV<K extends keyof ButtonVariantProps> = NonNullable<ButtonVariantProps[K]>;

@Component({
  selector: 'c-button-primary',
  standalone: true,
  imports: [ButtonComponent],
  template: `
    <c-button variant="primary" size="base" [class]="classes()" [disabled]="disabled()" (pressed)="this.pressed.emit($event)">
      <ng-content></ng-content>
    </c-button>
  `,
})
export class PrimaryButtonComponent {
  // variant
  readonly transform = input<VV<'transform'>>('default');

  // props
  readonly disabled  = input<boolean>(false);
  readonly className = input<string>('', { alias: 'class' });
  readonly pressed = output<MouseEvent>()

  readonly classes = computed(() =>
    cn(
      buttonVariants({
        transform: this.transform(),
      }),
      this.className(),
    )
  );
}

@Component({
  selector: 'c-button-primary-small',
  standalone: true,
  imports: [ButtonComponent],
  template: `
    <c-button variant="primary" size="sm" [class]="classes()" [disabled]="disabled()" (pressed)="this.pressed.emit($event)">
      <ng-content></ng-content>
    </c-button>
  `,
})
export class PrimarySmallButtonComponent {
  // variant
  readonly transform = input<VV<'transform'>>('default');

  // props
  readonly disabled  = input<boolean>(false);
  readonly className = input<string>('', { alias: 'class' });
  readonly pressed = output<MouseEvent>()

  readonly classes = computed(() =>
    cn(
      buttonVariants({
        transform: this.transform(),
      }),
      this.className(),
    )
  );
}

@Component({
  selector: 'c-button-secondary',
  standalone: true,
  imports: [ButtonComponent],
  template: `
    <c-button variant="secondary" size="base" [class]="classes()" [disabled]="disabled()" (pressed)="this.pressed.emit($event)">
      <ng-content></ng-content>
    </c-button>
  `,
})
export class SecondaryButtonComponent {
  // variant
  readonly transform = input<VV<'transform'>>('default');

  // props
  readonly disabled  = input<boolean>(false);
  readonly className = input<string>('', { alias: 'class' });
  readonly pressed = output<MouseEvent>()

  readonly classes = computed(() =>
    cn(
      buttonVariants({
        transform: this.transform(),
      }),
      this.className(),
    )
  );
}

@Component({
  selector: 'c-button-secondary-small',
  standalone: true,
  imports: [ButtonComponent],
  template: `
    <c-button variant="secondary" size="sm" [class]="classes()" [disabled]="disabled()" (pressed)="this.pressed.emit($event)">
      <ng-content></ng-content>
    </c-button>
  `,
})
export class SecondarySmallButtonComponent {
  // variant
  readonly transform = input<VV<'transform'>>('default');

  // props
  readonly disabled  = input<boolean>(false);
  readonly className = input<string>('', { alias: 'class' });
  readonly pressed = output<MouseEvent>()

  readonly classes = computed(() =>
    cn(
      buttonVariants({
        transform: this.transform(),
      }),
      this.className(),
    )
  );
}

// TODO : faire pour les autres boutons (link, outline, ghost, etc.) + revisé certains
