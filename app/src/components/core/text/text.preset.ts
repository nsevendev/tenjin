import {TextComponent} from './text.component';
import {Component, input} from '@angular/core';

@Component({
  selector: 'c-text-paragraph',
  imports: [TextComponent],
  template: `
    <c-text variant="paragraph" [class]="classes()">
      <ng-content></ng-content>
    </c-text>`,
})
export class ParagraphTextComponent {
  readonly classes = input<string>('', { alias: 'class' });
}

@Component({
  selector: 'c-text-small',
  imports: [TextComponent],
  template: `<c-text variant="small" size="sm" [class]="classes()"><ng-content></ng-content></c-text>`,
})
export class SmallTextComponent {
  readonly classes = input<string>('', { alias: 'class' });
}

@Component({
  selector: 'c-text-caption',
  imports: [TextComponent],
  template: `<c-text variant="caption" size="xs" [class]="classes()"><ng-content></ng-content></c-text>`,
})
export class CaptionTextComponent {
  readonly classes = input<string>('', { alias: 'class' });
}

@Component({
  selector: 'c-text-lead',
  imports: [TextComponent],
  template: `<c-text variant="lead" size="lg" [class]="classes()"><ng-content></ng-content></c-text>`,
})
export class LeadTextComponent {
  readonly classes = input<string>('', { alias: 'class' });
}

@Component({
  selector: 'c-text-muted',
  imports: [TextComponent],
  template: `<c-text variant="muted" [class]="classes()"><ng-content></ng-content></c-text>`,
})
export class MutedTextComponent {
  readonly classes = input<string>('', { alias: 'class' });
}

@Component({
  selector: 'c-code-text',
  imports: [TextComponent],
  template: `<c-text variant="code" [class]="classes()"><ng-content></ng-content></c-text>`,
})
export class CodeTextComponent {
  readonly classes = input<string>('', { alias: 'class' });
}
