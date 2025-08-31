import {Component, signal} from '@angular/core';
import {LinkTextComponent} from '../../../components/core/link-text/link-text.component';
import {CardComponent} from '../../../components/core/card/card.component';
import {ContainerComponent} from '../../../components/core/container/container.component';
import {TitleComponent} from '../../../components/core/title/title.component';
import {TextBlockquoteComponent, TextComponent, TextMessageComponent} from '../../../components/core/text/text.component';
import {ErrorTextMessageComponent, InfoTextMessageComponent, SuccessTextMessageComponent, WarningTextMessageComponent} from '../../../components/core/text/message.preset';
import {CaptionTextComponent, CodeTextComponent, LeadTextComponent, MutedTextComponent, ParagraphTextComponent} from '../../../components/core/text/text.preset';
import {PrimaryButtonComponent, PrimarySmallButtonComponent, SecondaryButtonComponent, SecondarySmallButtonComponent} from '../../../components/core/button/button.preset';
import {ButtonComponent} from '../../../components/core/button/button.component';
import {BadgeComponent} from '../../../components/core/badge/badge.component';
import {InputComponent} from '../../../components/core/input/input.component';
import {CInputPrefixDirective} from '../../../components/core/input/input-prefix.directive';
import {CInputSuffixDirective} from '../../../components/core/input/input-suffix.directive';

@Component({
  selector: 'page-design-system',
  templateUrl: './design-system.page.html',
  imports: [
    LinkTextComponent,
    CardComponent,
    ContainerComponent,
    TitleComponent,
    TextBlockquoteComponent,
    ErrorTextMessageComponent,
    WarningTextMessageComponent,
    SuccessTextMessageComponent,
    TextMessageComponent,
    CodeTextComponent,
    MutedTextComponent,
    LeadTextComponent,
    CaptionTextComponent,
    ParagraphTextComponent,
    TextComponent,
    SecondarySmallButtonComponent,
    SecondaryButtonComponent,
    PrimarySmallButtonComponent,
    PrimaryButtonComponent,
    ButtonComponent,
    BadgeComponent,
    InputComponent,
    CInputPrefixDirective,
    CInputSuffixDirective,
    InfoTextMessageComponent
  ]
})
export class DesignSystemPage {
  valueInput = signal('');
  username = signal('');
  filter = signal('');
  valueInputWithValue = signal('je suis un input avec une valeur par defaut');

  handlerBadge() {
    console.log('Design SystemPage BADGE clicked');
  }

  handlerButton() {
    console.log('Design SystemPage BUTTON clicked');
  }
}
