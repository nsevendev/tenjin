import {TextMessageComponent} from './text.component';
import {Component, input} from '@angular/core';

@Component({
  selector: 'c-text-message-success',
  standalone: true,
  imports: [TextMessageComponent],
  template: `<c-text-message variant="success" [class]="classes()"><ng-content></ng-content></c-text-message>`,
})
export class SuccessTextMessageComponent {
  readonly classes = input<string>('', { alias: 'class' });
}

@Component({
  selector: 'c-text-message-error',
  standalone: true,
  imports: [TextMessageComponent],
  template: `<c-text-message variant="error" [class]="classes()"><ng-content></ng-content></c-text-message>`,
})
export class ErrorTextMessageComponent {
  readonly classes = input<string>('', { alias: 'class' });
}

@Component({
  selector: 'c-text-message-warning',
  standalone: true,
  imports: [TextMessageComponent],
  template: `<c-text-message variant="warning" [class]="classes()"><ng-content></ng-content></c-text-message>`,
})
export class WarningTextMessageComponent {
  readonly classes = input<string>('', { alias: 'class' });
}

@Component({
  selector: 'c-text-message-info',
  standalone: true,
  imports: [TextMessageComponent],
  template: `<c-text-message variant="info" [class]="classes()"><ng-content></ng-content></c-text-message>`,
})
export class InfoTextMessageComponent {
  readonly classes = input<string>('', { alias: 'class' });
}
