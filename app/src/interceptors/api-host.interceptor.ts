import {HttpContextToken, HttpInterceptorFn} from '@angular/common/http';
import {environment} from '../environment/environment';

/** Contexte pour ignorer l'intercepteur, false par defaut pour executer l'intercepteur */
export const SKIP_API_HOST = new HttpContextToken<boolean>(() => false);

/** Intercepteur pour ajouter l'URL de base de l'API aux requêtes HTTP */
export const apiHostInterceptor: HttpInterceptorFn = (req, next) => {
  // si oublie du "/" au debut de l'url on le rajoute
  const path = req.url.startsWith('/') ? req.url : `/${req.url}`;
  const url  = `${environment.API_BASE}${environment.API_PREFIX}${path}`;

  return next(
    req.clone({
      url,
      withCredentials: environment.WITH_CREDENTIALS,
    })
  );
};
