import {BlockquoteVariantProps, blockquoteVariants, MessageVariantProps, messageVariants, TextVariantProps, textVariants} from './text.variant';
import {Component, computed, input} from '@angular/core';
import {cn} from '../../../utils/class-name/cn';

type VVT<K extends keyof TextVariantProps> = NonNullable<TextVariantProps[K]>;
type VVTM<K extends keyof MessageVariantProps> = NonNullable<MessageVariantProps[K]>;
type VVTBQ<K extends keyof BlockquoteVariantProps> = NonNullable<BlockquoteVariantProps[K]>;

@Component({
  selector: 'c-text',
  template: `
    <p [class]="classes()"><ng-content></ng-content></p>
  `,
})
export class TextComponent {
  readonly variant = input<VVT<'variant'>>('paragraph');
  readonly size    = input<VVT<'size'>>('base');
  readonly weight  = input<VVT<'weight'>>('normal');
  readonly align   = input<VVT<'align'>>('left');
  readonly color   = input<VVT<'color'>>('default');
  readonly className = input<string>('', { alias: 'class' });

  readonly classes = computed(() =>
    cn(textVariants({
      variant: this.variant(),
      size: this.size(),
      weight: this.weight(),
      align: this.align(),
      color: this.color(),
    }), this.className())
  );
}

@Component({
  selector: 'c-text-message',
  template: `
    <p [class]="classes()"><ng-content></ng-content></p>
  `,
})
export class TextMessageComponent {
  readonly variant = input<VVTM<'variant'>>('info');
  readonly size    = input<VVTM<'size'>>('base');
  readonly className = input<string>('', { alias: 'class' });

  readonly classes = computed(() =>
    cn(messageVariants({
      variant: this.variant(),
      size: this.size(),
    }), this.className())
  );
}

@Component({
  selector: 'c-text-blockquote',
  template: `
    <blockquote [class]="classes()"><ng-content/></blockquote>
  `,
})
export class TextBlockquoteComponent {
  readonly variant = input<VVTBQ<'variant'>>('default');
  readonly className = input<string>('', { alias: 'class' });

  readonly classes = computed(() =>
    cn(blockquoteVariants({ variant: this.variant() }), this.className())
  );
}
