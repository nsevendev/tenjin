import {Component} from '@angular/core';
import {LinkTextComponent} from '../../../components/core/link-text/link-text.component';
import {ContainerComponent} from '../../../components/core/container/container.component';

@Component({
  selector: 'page-dashboard',
  templateUrl: './dashboard.page.html',
  imports: [
    LinkTextComponent,
    ContainerComponent
  ]
})
export class DashboardPage {}
