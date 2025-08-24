import {component$} from "@builder.io/qwik";
import {Form} from "@builder.io/qwik-city";
import {Input} from "~/components/core/input/input";
import {Button} from "~/components/core/button/button";

export const RegisterCandidate = component$(() => {
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
                required
                name="lastname"
                class="w-[350px]"
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
