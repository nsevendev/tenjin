import {$, component$, useSignal} from "@builder.io/qwik";
import {Title} from "~/components/core/title/title";
import {SectionCard} from "~/components/core/card/card";
import {Button} from "~/components/core/button/button";
import {Badge} from "~/components/core/badge/badge";
import {RegisterCandidate} from "~/components/register/register-candidate/register-candidate";
import {RegisterPro} from "~/components/register/register-pro/register-pro";

export default component$(() => {
    const isCandidate = useSignal(true)
    const hoverBadgeCandidate = isCandidate.value ? '' : 'hover:bg-blue-pastel-500 hover:text-blue-900'
    const hoverBadgeCompany = isCandidate.value ? 'hover:bg-blue-pastel-500 hover:text-blue-900' : ''
    
    const changeForm = $(() => {
        isCandidate.value = !isCandidate.value
    })
    
    return (
        <>
            <div class="flex h-full">
                <div class="flex-1 flex flex-col items-end justify-center text-black p-8">
                    <a href="/" class="cursor-pointer hover:text-blue-700 hover:underline decoration-[1px] underline-offset-2">
                        <Title as="h1" size="6xl" weight="extrabold" class="uppercase">Tenjin</Title>
                    </a>
                    {
                        isCandidate.value
                            ? (
                                <>
                                    <Title as="h2" size="3xl" weight="bold" class="mt-2 mb-6">Du savoir à l’emploi</Title>
                                    <p class="text-xl font-light">Études, formations, carrières</p>
                                    <p class="text-xl font-light">votre avenir commence ici</p>
                                </>
                            )
                            : (
                                <>
                                    <Title as="h2" size="3xl" weight="bold" class="mt-2">L’éducation et l’emploi</Title>
                                    <Title as="h2" size="3xl" weight="bold" class="mb-6">main dans la main</Title>
                                    <p class="text-xl font-light">Établissements, recruteurs, formateurs </p>
                                    <p class="text-xl font-light">Construisons des parcours qui comptent</p>
                                </>
                        )
                    }
                </div>
                <div class="flex-1 flex flex-col items-start justify-center">
                    <div class="flex gap-4 w-full">
                        <Badge size="sm" color={isCandidate.value ? 'green' : 'gray'} onClick$={changeForm} class={`mb-4 cursor-pointer ${hoverBadgeCandidate}`}>Candidat</Badge>
                        <Badge size="sm" color={isCandidate.value ? 'gray' : 'yellow'} onClick$={changeForm} class={`mb-4 cursor-pointer ${hoverBadgeCompany}`}>Professionel</Badge>
                    </div>
                    <SectionCard>
                    {
                        isCandidate.value
                            ?  <RegisterCandidate/>
                            : <RegisterPro/>
                    }
                        <div class="h-[1px] bg-gray-300 w-full my-8"></div>
                        <div class="flex items-center justify-center w-full">
                            <Button variant="success" size="sm" type="button" transform="upper">
                                <a href="/your/login">Se connecter</a>
                            </Button>
                        </div>
                    </SectionCard>
                </div>
            </div>
        </>
    )
})
