import {component$} from "@builder.io/qwik";
import {DocumentHead, Form} from "@builder.io/qwik-city";
import {Card} from "~/components/core/card/card";
import {CardTitle} from "~/components/core/title/title";
import {Button} from "~/components/core/button/button";
import {useForm} from "~/hooks/form-hook/form-hook";
import {useLoginAction} from "~/routes/your/login/_login-action";
import {Input} from "~/components/core/input/input";

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
    if (loginAction.value?.fieldErrors) {
        Object.entries(loginAction.value.fieldErrors).forEach(([field, errors]) => {
            if (Array.isArray(errors) && errors.length > 0) {
                form.setFieldError(field as keyof LoginForm, errors[0]);
            }
        });
    }
    
    return (
        <div class="flex h-full">
            <div class="flex-1">
                <Card>
                    <CardTitle>Connexion</CardTitle>
                    <Form action={loginAction} class="flex flex-col gap-4">
                        <Input
                            label="Email"
                            type="email"
                            placeholder="votre@email.com"
                            required
                            name="email"
                            {...emailProps}
                        />
                        <Input
                            label="Mot de passe"
                            type="password"
                            placeholder="••••••••"
                            required
                            name="password"
                            {...passwordProps}
                        />
                        <Button size="sm" type="submit">
                            Se connecter
                        </Button>
                    </Form>
                </Card>
            </div>
            <div class="flex-1"></div>
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
