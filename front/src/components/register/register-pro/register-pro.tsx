import {component$} from "@builder.io/qwik";
import {Form} from "@builder.io/qwik-city";
import {Input} from "~/components/core/input/input";
import {Button} from "~/components/core/button/button";

export const RegisterPro = component$(() => {
    return (
        <Form class="flex flex-col gap-4">
            <Input
                label="Email"
                type="email"
                placeholder="votre adresse email"
                required
                name="email"
                class="w-[350px]"
            />
            <Input
                label="Nom utilisateur"
                type="text"
                placeholder="votre nom"
                name="lastname"
                class="w-[350px]"
                required
            />
            <Input
                label="Fonction de l'utilisateur au sein de l'entreprise"
                type="text"
                placeholder="votre fonction"
                name="function"
                class="w-[350px]"
                required
            />
            <Input
                label="Mot de passe"
                type="password"
                name="password"
                required
            />
            <Button size="base" type="submit" transform="upper" class="mt-4">
                Créer votre compte
            </Button>
        </Form>
    )
})
