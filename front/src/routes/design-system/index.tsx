import {$, component$} from "@builder.io/qwik";
import {DocumentHead} from "@builder.io/qwik-city";
import {CardSubTitle, CardTitle, PageSubTitle, PageTitle, SectionSubTitle, SectionTitle, SmallTitle, Title} from "~/components/core/title/title";
import {Badge} from "~/components/core/badge/badge";
import {ArticleCard, SectionCard} from "~/components/core/card/card";
import {Button, PrimaryButton, PrimarySmallButton, SecondaryButton, SecondarySmallButton} from "~/components/core/button/button";
import {LinkText} from "~/components/core/link-text/link-text";

export default component$(() => {
    const handler = $(() => {
        console.log("J'ai été cliqué !");
    })
    
    return (
        <div className="py-8 px-16">
            {/* LINK TEXT */}
            <div class="mb-10">
                <LinkText href="/" >Retour accueil</LinkText>
            </div>
            
            {/* TITIRE */}
            <SectionCard>
                <Title>Je suis un titre de base</Title>
                <PageTitle>Je suis un titre de page</PageTitle>
                <PageSubTitle>Je suis un sous titre de page</PageSubTitle>
                <SectionTitle>Je suis un titre de section</SectionTitle>
                <SectionSubTitle>Je suis un sous titre de section</SectionSubTitle>
                <CardTitle>Je suis un titre d'une card</CardTitle>
                <CardSubTitle>Je suis un sous titre d'une card</CardSubTitle>
                <SmallTitle>Je suis un petit titre</SmallTitle>
            </SectionCard>
            
            <div className="h-8"></div>
            
            {/* BADGE */}
            <ArticleCard class="gap-4">
                <Badge class="cursor-pointer" onClick$={handler}>Basic</Badge>
                <Badge color="blue" size="base">Blue Large</Badge>
                <Badge color="red" size="sm">Red Small</Badge>
                <Badge color="green">Green</Badge>
                <Badge color="yellow">Yellow</Badge>
                <Badge color="purple">Purple</Badge>
                <Badge color="orange">Orange</Badge>
                <Badge color="gray">Gray</Badge>
                <Badge color="indigo">Indigo</Badge>
                <Badge color="pink">Pink</Badge>
            </ArticleCard>
            
            <div className="h-8"></div>
            
            {/* BUTTON */}
            <ArticleCard class="gap-4" containPosition="centerColumn">
                <PrimaryButton>Primary button</PrimaryButton>
                <SecondaryButton>Secondary button</SecondaryButton>
                <PrimarySmallButton>Primary small button</PrimarySmallButton>
                <SecondarySmallButton>Secondary small button</SecondarySmallButton>
                <Button variant="warning" >Button warring</Button>
                <Button variant="success" size="sm" >Button small success</Button>
                <Button variant="error" transform="upper" >Button uppercase error</Button>
                <Button variant="indigo" size="sm" transform="upper" >Button small uppercase indigo</Button>
                <Button variant="pinkPastel" >Button pink pastel</Button>
                <Button variant="successPastel" size="sm" >Button small success pastel</Button>
                <Button variant="errorPastel" transform="upper" >Button uppercase error pastel</Button>
                <Button variant="primaryPastel" size="sm" transform="upper" >Button small uppercase primary pastel</Button>
            </ArticleCard>
        </div>
    );
});

export const head: DocumentHead = {
    title: "Composants du Design System",
    meta: [
        {
            name: "description",
            content: "Page de test des composants du design system",
        },
    ],
};
