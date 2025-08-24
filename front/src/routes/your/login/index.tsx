import {component$} from "@builder.io/qwik";
import {DocumentHead, Form} from "@builder.io/qwik-city";
import {SectionCard} from "~/components/core/card/card";
import {Title} from "~/components/core/title/title";
import {Button} from "~/components/core/button/button";
import {useForm} from "~/hooks/form-hook/form-hook";
import {useLoginAction} from "~/routes/your/login/_login-action";
import {Input} from "~/components/core/input/input";
import {LinkText} from "~/components/core/link-text/link-text";

type LoginForm = {
    email: string;
    password: string;
};

export default component$(() => {
    const loginAction = useLoginAction();
    
    const form = useForm<LoginForm>({
        initialValues: {
            email: '',
            password: ''
        },
        
        // Validation côté client (temps réel)
        validationRules: {
            email: {
                required: true,
                email: true
            },
            password: {
                required: true,
                minLength: 1
            }
        }
    });
    
    const emailProps = form.getFieldProps('email');
    const passwordProps = form.getFieldProps('password');
    
    // Gestion des erreurs serveur dans le hook
    if (loginAction.value?.failed) {
        Object.entries(loginAction.value.message).forEach(([field, errors]) => {
            if (Array.isArray(errors) && errors.length > 0) {
                form.setFieldError(field as keyof LoginForm, errors[0]);
            }
        });
    }
    
    return (
        <div class="flex h-full">
            <div class="flex-1 flex flex-col items-end justify-center text-black p-8">
                <a href="/" class="cursor-pointer hover:text-blue-700 hover:underline decoration-[1px] underline-offset-2">
                    <Title as="h1" size="6xl" weight="extrabold" class="uppercase">Tenjin</Title>
                </a>
                <Title as="h2" size="3xl" weight="bold" class="mt-2 mb-6">L’éducation rencontre l’emploi</Title>
                <p class="text-xl font-light">Candidats, entreprises, universités</p>
                <p class="text-xl font-light">connectez-vous</p>
            </div>
            <div class="flex-1 flex items-center justify-start">
                <SectionCard>
                    <Form action={loginAction} class="flex flex-col gap-4">
                        <Input
                            label="Email"
                            type="email"
                            placeholder="votre adresse email"
                            required
                            name="email"
                            class="w-[350px]"
                            {...emailProps}
                        />
                        <Input
                            label="Mot de passe"
                            type="password"
                            name="password"
                            required
                            {...passwordProps}
                        />
                        <Button size="base" type="submit" transform="upper" class="mt-4">
                            Se connecter
                        </Button>
                    </Form>
                    <div class="flex justify-end mt-4 w-full">
                        <LinkText size="sm">Mot de passe oublié ?</LinkText>
                    </div>
                    <div class="h-[1px] bg-gray-300 w-full my-8"></div>
                    <div class="flex items-center justify-center w-full">
                        <Button variant="success" size="sm" type="button" transform="upper">
                            <a href="/register">Créer un compte</a>
                        </Button>
                    </div>
                </SectionCard>
            </div>
        </div>
    );
});

export const head: DocumentHead = {
    title: "Tenjin | Login",
    meta: [
        {
            name: "description",
            content: "Tenjin page de connexion utilisateur",
        },
    ],
};

export {useLoginAction} from "~/routes/your/login/_login-action"
