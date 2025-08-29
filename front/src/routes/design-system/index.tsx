import {$, component$} from "@builder.io/qwik";
import {DocumentHead} from "@builder.io/qwik-city";
import {CardSubTitle, CardTitle, PageSubTitle, PageTitle, SectionSubTitle, SectionTitle, SmallTitle, Title} from "~/components/core/title/title";
import {Badge} from "~/components/core/badge/badge";
import {ArticleCard, SectionCard} from "~/components/core/card/card";
import {Button, PrimaryButton, PrimarySmallButton, SecondaryButton, SecondarySmallButton} from "~/components/core/button/button";
import {LinkText} from "~/components/core/link-text/link-text";
import {Container, FlexRow, FlexRowCenter, FlexRowBetween, FlexCol, FlexColCenter, CenteredContainer, PaddedContainer} from "~/components/core/container/container";
import {Text, Paragraph, SmallText, Caption, Lead, MutedText, CodeText, SuccessMessage, ErrorMessage, WarningMessage, InfoMessage, DefaultQuote, PrimaryQuote} from "~/components/core/text/text";
import {AcademicIcon} from "~/components/core/icon-svg/academic/academic";
import {BellIcon} from "~/components/core/icon-svg/bell/bell";
import {CalendarIcon} from "~/components/core/icon-svg/calendar/calendar";
import {ChartBarIcon} from "~/components/core/icon-svg/chart-bar/chart-bar";
import {CogIcon} from "~/components/core/icon-svg/cog/cog";
import {DocumentTextIcon} from "~/components/core/icon-svg/document-text/document-text";
import {EnvelopeIcon} from "~/components/core/icon-svg/envelope/envelope";
import {EyeIcon} from "~/components/core/icon-svg/eye/eye";
import {HomeIcon} from "~/components/core/icon-svg/home/home";
import {LockCloseIcon} from "~/components/core/icon-svg/lock/lock-close";
import {LockOpenIcon} from "~/components/core/icon-svg/lock/lock-open";
import {PhoneIcon} from "~/components/core/icon-svg/phone/phone";
import {ShieldCheckIcon} from "~/components/core/icon-svg/shield-check/shield-check";
import {UserIcon} from "~/components/core/icon-svg/user/user";
import {WarningIcon} from "~/components/core/icon-svg/warning/warning";
import {SunIcon} from "~/components/core/icon-svg/sun/sun";
import {MoonIcon} from "~/components/core/icon-svg/moon/moon";

export default component$(() => {
    const handler = $(() => {
        console.log("J'ai été cliqué !");
    })
    
    return (
        <div class="py-8 px-16 ">
            {/* THEME TOGGLE */}
            <FlexRowBetween class="mb-10 flex justify-between items-center">
                <LinkText href="/" >Retour accueil</LinkText>
            </FlexRowBetween>
            
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
            
            <div class="h-8"></div>
            
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
            
            <div class="h-8"></div>
            
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
            
            <div class="h-8"></div>
            
            {/* CONTAINERS */}
            <SectionCard>
                <SectionTitle>Containers</SectionTitle>
                <SmallText color="primary">Les bordures sont a titre indicatif et ne font pas partie du composant cela aide à la visualisation des containers</SmallText>
                <div class="space-y-4">
                    <div class="space-y-2">
                        <Text class="my-4">Container de base</Text>
                        <Container class="border border-gray-300 p-2">
                            <Text color="default">Contenu dans un container basique</Text>
                        </Container>
                    </div>
                    
                    <div class="space-y-2">
                        <SmallText>FlexRow - Alignement horizontal</SmallText>
                        <FlexRow class="border border-gray-300 p-2 gap-2">
                            <div class="bg-blue-100 p-2 rounded">Item 1</div>
                            <div class="bg-blue-100 p-2 rounded">Item 2</div>
                            <div class="bg-blue-100 p-2 rounded">Item 3</div>
                        </FlexRow>
                    </div>
                    
                    <div class="space-y-2">
                        <SmallText>FlexRowCenter - Centre horizontal et vertical</SmallText>
                        <FlexRowCenter class="border border-gray-300 p-4 h-20 gap-2">
                            <div class="bg-green-100 p-2 rounded">Centré 1</div>
                            <div class="bg-green-100 p-2 rounded">Centré 2</div>
                        </FlexRowCenter>
                    </div>
                    
                    <div class="space-y-2">
                        <SmallText>FlexRowBetween - Espacement justifié</SmallText>
                        <FlexRowBetween class="border border-gray-300 p-2">
                            <div class="bg-purple-100 p-2 rounded">Gauche</div>
                            <div class="bg-purple-100 p-2 rounded">Droite</div>
                        </FlexRowBetween>
                    </div>
                    
                    <div class="space-y-2">
                        <SmallText>FlexCol - Alignement vertical</SmallText>
                        <FlexCol class="border border-gray-300 p-2 gap-2 w-32">
                            <div class="bg-yellow-100 p-2 rounded">Item 1</div>
                            <div class="bg-yellow-100 p-2 rounded">Item 2</div>
                            <div class="bg-yellow-100 p-2 rounded">Item 3</div>
                        </FlexCol>
                    </div>
                    
                    <div class="space-y-2">
                        <SmallText>PaddedContainer - Avec padding</SmallText>
                        <PaddedContainer class="border border-gray-300 bg-gray-50">
                            <Text color="default">Ce contenu a un padding automatique</Text>
                        </PaddedContainer>
                    </div>
                    
                    <div class="space-y-2">
                        <SmallText>CenteredContainer - Centrage complet</SmallText>
                        <CenteredContainer class="border border-gray-300 h-24">
                            <div class="bg-red-100 p-2 rounded">Parfaitement centré</div>
                        </CenteredContainer>
                    </div>
                </div>
            </SectionCard>
            
            <div class="h-8"></div>
            
            {/* TEXT COMPONENTS */}
            <SectionCard>
                <SectionTitle>Composants de Texte</SectionTitle>
                <div class="space-y-4">
                    <div class="space-y-2">
                        <SmallText>Variations de texte</SmallText>
                        <div class="space-y-2">
                            <Paragraph>Ceci est un paragraphe standard</Paragraph>
                            <Lead>Ceci est un texte d'introduction (lead)</Lead>
                            <SmallText>Ceci est un petit texte</SmallText>
                            <Caption>Ceci est une légende</Caption>
                            <MutedText>Ceci est un texte atténué</MutedText>
                            <CodeText>const code = "exemple";</CodeText>
                        </div>
                    </div>
                    
                    <div class="space-y-2">
                        <SmallText>Alignements de texte</SmallText>
                        <div class="space-y-2">
                            <Text align="left">Texte aligné à gauche</Text>
                            <Text align="center">Texte centré</Text>
                            <Text align="right">Texte aligné à droite</Text>
                        </div>
                    </div>
                    
                    <div class="space-y-2">
                        <SmallText>Poids de police</SmallText>
                        <div class="space-y-2">
                            <Text weight="light" color="default">Texte léger</Text>
                            <Text weight="normal" color="default">Texte normal</Text>
                            <Text weight="medium" color="default">Texte medium</Text>
                            <Text weight="semibold" color="default">Texte semi-bold</Text>
                            <Text weight="bold" color="default">Texte gras</Text>
                        </div>
                    </div>
                    
                    <div class="space-y-2">
                        <SmallText>Tailles de texte</SmallText>
                        <div class="space-y-2">
                            <Text size="xs" color="default">Texte extra small</Text>
                            <Text size="sm" color="default">Texte small</Text>
                            <Text size="base" color="default">Texte base</Text>
                            <Text size="lg" color="default">Texte large</Text>
                            <Text size="xl" color="default">Texte extra large</Text>
                        </div>
                    </div>
                    
                    <div class="space-y-2">
                        <SmallText>Couleurs de texte</SmallText>
                        <div class="space-y-2">
                            <Text color="default">Texte couleur par défaut</Text>
                            <Text color="muted">Texte atténué</Text>
                            <Text color="primary">Texte primaire</Text>
                            <Text color="secondary">Texte secondaire</Text>
                            <Text color="success">Texte succès</Text>
                            <Text color="error">Texte erreur</Text>
                            <Text color="warning">Texte avertissement</Text>
                        </div>
                        
                        <div class="space-y-2">
                            <SmallText>Test classes fixes</SmallText>
                            <div class="space-y-2">
                                <p class="text-gray-900 dark:text-gray-100">Test direct: Texte qui devrait être noir/blanc</p>
                                <p class="text-blue-600 dark:text-blue-400">Test direct: Texte bleu qui change</p>
                                <p class="bg-gray-100 dark:bg-gray-800 p-2 text-gray-900 dark:text-gray-100">Test background qui change avec texte</p>
                                <Paragraph>Test Paragraph component</Paragraph>
                                <SmallText>Test SmallText component</SmallText>
                                <Text>Test Text component de base</Text>
                                <Text color="primary">Test Text component primary</Text>
                            </div>
                        </div>
                    </div>
                </div>
            </SectionCard>
            
            <div class="h-8"></div>
            
            {/* MESSAGE COMPONENTS */}
            <SectionCard>
                <SectionTitle>Messages</SectionTitle>
                <div class="space-y-4">
                    <SuccessMessage>Message de succès</SuccessMessage>
                    <ErrorMessage>Message d'erreur</ErrorMessage>
                    <WarningMessage>Message d'avertissement</WarningMessage>
                    <InfoMessage>Message d'information</InfoMessage>
                </div>
            </SectionCard>
            
            <div class="h-8"></div>
            
            {/* BLOCKQUOTE COMPONENTS */}
            <SectionCard>
                <SectionTitle>Citations</SectionTitle>
                <div class="space-y-4">
                    <DefaultQuote>Citation par défaut - "La simplicité est la sophistication suprême."</DefaultQuote>
                    <PrimaryQuote>Citation primaire - "Le design n'est pas juste ce à quoi ça ressemble. Le design, c'est comment ça fonctionne."</PrimaryQuote>
                </div>
            </SectionCard>
            
            <div class="h-8"></div>
            
            {/* ICON SVG COMPONENTS */}
            <SectionCard>
                <SectionTitle>Icônes SVG</SectionTitle>
                <div class="space-y-6">
                    <div class="space-y-2">
                        <SmallText>Icônes par défaut (taille 24px)</SmallText>
                        <FlexRowCenter class="gap-4 p-4 border border-gray-200 rounded bg-gray-50">
                            <AcademicIcon />
                            <BellIcon />
                            <CalendarIcon />
                            <ChartBarIcon />
                            <CogIcon />
                            <DocumentTextIcon />
                            <EnvelopeIcon />
                            <EyeIcon />
                            <HomeIcon />
                            <LockCloseIcon />
                            <LockOpenIcon />
                            <PhoneIcon />
                            <ShieldCheckIcon />
                            <UserIcon />
                            <WarningIcon />
                        </FlexRowCenter>
                    </div>
                    
                    <div class="space-y-2">
                        <SmallText>Icônes colorées (couleur primary)</SmallText>
                        <FlexRowCenter class="gap-4 p-4 border border-gray-200 rounded bg-gray-50">
                            <AcademicIcon color="primary" />
                            <BellIcon color="primary" />
                            <CalendarIcon color="primary" />
                            <ChartBarIcon color="primary" />
                            <CogIcon color="primary" />
                            <DocumentTextIcon color="primary" />
                            <EnvelopeIcon color="primary" />
                            <EyeIcon color="primary" />
                            <HomeIcon color="primary" />
                            <LockCloseIcon color="primary" />
                            <LockOpenIcon color="primary" />
                            <PhoneIcon color="primary" />
                            <ShieldCheckIcon color="primary" />
                            <UserIcon color="primary" />
                            <WarningIcon color="primary" />
                        </FlexRowCenter>
                    </div>
                    
                    <div class="space-y-2">
                        <SmallText>Icônes avec différentes couleurs</SmallText>
                        <FlexRowCenter class="gap-4 p-4 border border-gray-200 rounded bg-gray-50">
                            <HomeIcon color="success" />
                            <BellIcon color="warning" />
                            <WarningIcon color="error" />
                            <ShieldCheckIcon color="success" />
                            <CogIcon color="secondary" />
                            <UserIcon color="info" />
                        </FlexRowCenter>
                    </div>
                    
                    <div class="space-y-2">
                        <SmallText>Différentes tailles</SmallText>
                        <FlexRowCenter class="gap-4 p-4 border border-gray-200 rounded bg-gray-50 items-end">
                            <HomeIcon size={16} />
                            <HomeIcon size={20} />
                            <HomeIcon size={24} />
                            <HomeIcon size={32} />
                            <HomeIcon size={40} />
                            <HomeIcon size={48} />
                        </FlexRowCenter>
                    </div>
                    
                    <div class="space-y-2">
                        <SmallText>Icônes interactives</SmallText>
                        <FlexRowCenter class="gap-4 p-4 border border-gray-200 rounded bg-gray-50">
                            <BellIcon class="cursor-pointer hover:scale-110 transition-transform" onClick$={handler} />
                            <CogIcon class="cursor-pointer hover:rotate-45 transition-transform" onClick$={handler} />
                            <EyeIcon class="cursor-pointer hover:opacity-50 transition-opacity" onClick$={handler} />
                        </FlexRowCenter>
                    </div>
                    
                    <div class="space-y-2">
                        <SmallText>Nouvelles icônes : Sun & Moon</SmallText>
                        <FlexRowCenter class="gap-4 p-4 border border-gray-200 rounded bg-gray-50">
                            <SunIcon size={32} />
                            <MoonIcon size={32} />
                        </FlexRowCenter>
                    </div>
                    
                    <div class="space-y-2">
                        <SmallText>Toutes les icônes disponibles avec noms</SmallText>
                        <div class="grid grid-cols-5 gap-4 p-4 border border-gray-200 rounded bg-gray-50">
                            <FlexColCenter class="gap-2 p-2">
                                <AcademicIcon size={32} />
                                <Caption>Academic</Caption>
                            </FlexColCenter>
                            <FlexColCenter class="gap-2 p-2">
                                <BellIcon size={32} />
                                <Caption>Bell</Caption>
                            </FlexColCenter>
                            <FlexColCenter class="gap-2 p-2">
                                <CalendarIcon size={32} />
                                <Caption>Calendar</Caption>
                            </FlexColCenter>
                            <FlexColCenter class="gap-2 p-2">
                                <ChartBarIcon size={32} />
                                <Caption>ChartBar</Caption>
                            </FlexColCenter>
                            <FlexColCenter class="gap-2 p-2">
                                <CogIcon size={32} />
                                <Caption>Cog</Caption>
                            </FlexColCenter>
                            <FlexColCenter class="gap-2 p-2">
                                <DocumentTextIcon size={32} />
                                <Caption>DocumentText</Caption>
                            </FlexColCenter>
                            <FlexColCenter class="gap-2 p-2">
                                <EnvelopeIcon size={32} />
                                <Caption>Envelope</Caption>
                            </FlexColCenter>
                            <FlexColCenter class="gap-2 p-2">
                                <EyeIcon size={32} />
                                <Caption>Eye</Caption>
                            </FlexColCenter>
                            <FlexColCenter class="gap-2 p-2">
                                <HomeIcon size={32} />
                                <Caption>Home</Caption>
                            </FlexColCenter>
                            <FlexColCenter class="gap-2 p-2">
                                <LockCloseIcon size={32} />
                                <Caption>LockClose</Caption>
                            </FlexColCenter>
                            <FlexColCenter class="gap-2 p-2">
                                <LockOpenIcon size={32} />
                                <Caption>LockOpen</Caption>
                            </FlexColCenter>
                            <FlexColCenter class="gap-2 p-2">
                                <PhoneIcon size={32} />
                                <Caption>Phone</Caption>
                            </FlexColCenter>
                            <FlexColCenter class="gap-2 p-2">
                                <ShieldCheckIcon size={32} />
                                <Caption>ShieldCheck</Caption>
                            </FlexColCenter>
                            <FlexColCenter class="gap-2 p-2">
                                <UserIcon size={32} />
                                <Caption>User</Caption>
                            </FlexColCenter>
                            <FlexColCenter class="gap-2 p-2">
                                <WarningIcon size={32} />
                                <Caption>Warning</Caption>
                            </FlexColCenter>
                            <FlexColCenter class="gap-2 p-2">
                                <SunIcon size={32} />
                                <Caption>Sun</Caption>
                            </FlexColCenter>
                            <FlexColCenter class="gap-2 p-2">
                                <MoonIcon size={32} />
                                <Caption>Moon</Caption>
                            </FlexColCenter>
                        </div>
                    </div>
                </div>
            </SectionCard>
            
            <div class="h-8"></div>
            
            {/* THEME TOGGLE COMPONENT */}
            <SectionCard>
                <SectionTitle>Theme Toggle</SectionTitle>
                <div class="space-y-6">
                    <div class="space-y-2">
                        <SmallText>Bouton de basculement de thème avec animation</SmallText>
                        <FlexRowCenter class="gap-4 p-4 border border-gray-200 rounded bg-gray-50">
                        </FlexRowCenter>
                    </div>
                    
                    <div class="space-y-2">
                        <SmallText>Différentes variantes</SmallText>
                        <FlexRowCenter class="gap-4 p-4 border border-gray-200 rounded bg-gray-50">
                        
                        </FlexRowCenter>
                    </div>
                    
                    <div class="space-y-2">
                        <SmallText>Différentes tailles</SmallText>
                        <FlexRowCenter class="gap-4 p-4 border border-gray-200 rounded bg-gray-50 items-end">
                        
                        </FlexRowCenter>
                    </div>
                    
                    <div class="space-y-2">
                        <SmallText>Instructions d'utilisation</SmallText>
                        <div class="p-4 bg-blue-50 border border-blue-200 rounded">
                            <Text>
                                Cliquez sur le bouton pour basculer entre le mode clair (icône soleil) et le mode sombre (icône lune).
                                Le thème est automatiquement sauvegardé dans le localStorage et appliqué à toute l'application.
                            </Text>
                        </div>
                    </div>
                </div>
            </SectionCard>
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
