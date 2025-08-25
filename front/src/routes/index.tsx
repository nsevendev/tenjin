import {component$} from "@builder.io/qwik";
import {DocumentHead} from "@builder.io/qwik-city";
import {LinkText} from "~/components/core/link-text/link-text";
import {PageTitle} from "~/components/core/title/title";

export default component$(() => {
  return (
      <div>
          <PageTitle>Bienvenue</PageTitle>
          <div>
              <LinkText href="/design-system">Voir les composants</LinkText>
          </div>
          <div>
              <LinkText href="/your/login">Connexion</LinkText>
          </div>
          <div>
              <LinkText href="/register">Créer son compte</LinkText>
          </div>
      </div>
  );
});

export const head: DocumentHead = {
    title: "Tenjin",
    meta: [
        {
            name: "description",
            content: "Tenjin",
        },
    ],
};
