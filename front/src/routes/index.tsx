import {component$} from "@builder.io/qwik";
import {DocumentHead, Link} from "@builder.io/qwik-city";

export default component$(() => {
  return (
      <div>
          <h1>Bienvenu</h1>
          <div>
              <Link href="/design-system" class={"text-blue-500 hover:text-blue-700 underline cursor-pointer"}>Voir les composants</Link>
          </div>
          <div>
              <Link href="/your/login" class={"text-blue-500 hover:text-blue-700 underline cursor-pointer"}>Connexion</Link>
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
