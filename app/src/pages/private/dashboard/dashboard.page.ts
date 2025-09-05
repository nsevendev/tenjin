import {Component, effect} from '@angular/core';
import {LinkTextComponent} from '../../../components/core/link-text/link-text.component';
import {ContainerComponent} from '../../../components/core/container/container.component';
import {HttpContext, httpResource} from '@angular/common/http';
import {SKIP_API_HOST} from '../../../interceptors/api-host.interceptor';

@Component({
  selector: 'page-dashboard',
  templateUrl: './dashboard.page.html',
  imports: [
    LinkTextComponent,
    ContainerComponent
  ]
})
export class DashboardPage {
  // TODO : a supprimer si d'autre appels API sont fait
  // TODO : supprimer aussi la partie test qui utilise le httpResource
  // par default c'est deja false, mais ceci est pour montrer l'utilisation du contexte
  ctx = new HttpContext().set(SKIP_API_HOST, false);
  httpClient = httpResource(() => ({
    url: '/testconnexionapp',
    context: this.ctx
  }))

  constructor() {
    // TODO : a supprimer aussi
    effect(() => {
      if (this.httpClient.isLoading()) console.log('Chargement en cours...');
      if (this.httpClient.hasValue()) console.log('Réponse API:', this.httpClient.value());
      if (this.httpClient.error()) console.error('Erreur API:', this.httpClient.error());
    });
  }
}
