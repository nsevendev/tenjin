import {$, component$} from "@builder.io/qwik";
import {DocumentHead, Link} from "@builder.io/qwik-city";
import {LinkText} from "~/components/core/link-text/link-text";
import {Badge} from "~/components/core/badge/badge";
import {Card} from "~/components/core/card/card";
import {Button} from "~/components/core/button/button";
import {Input} from "~/components/core/input/input";
import {LockCloseIcon} from "~/components/core/icon/lock/lock-close";

export default component$(() => {
    const handler = $(() => {
        console.log('Form submitted!');
    })
    
    return (
        <>
            <h1 class="text-red-500 mb-4">Composant</h1>
            <div>
                <LinkText href="/">Retour à l'accueil</LinkText>
            </div>
            
            <div class="flex">
                <div class="ml-8">
                    <div class="p-8 mt-2 bg-blue-50 w-fit">
                        je suis en 50
                    </div>
                    <div class="p-8 mt-2 bg-blue-100 w-fit">
                        je suis en 100
                    </div>
                    <div class="p-8 mt-2 bg-blue-200 w-fit">
                        je suis en 200
                    </div>
                    <div class="p-8 mt-2 bg-blue-300 w-fit">
                        je suis en 300
                    </div>
                    <div class="p-8 mt-2 bg-blue-400 w-fit">
                        je suis en 400
                    </div>
                    <div class="p-8 mt-2 bg-blue-500 w-fit text-white">
                        je suis en 500
                    </div>
                    <div class="p-8 mt-2 bg-blue-600 w-fit text-white">
                        je suis en 600
                    </div>
                    <div class="p-8 mt-2 bg-blue-700 w-fit text-white">
                        je suis en 700
                    </div>
                    <div class="p-8 mt-2 bg-blue-800 w-fit text-white">
                        je suis en 800
                    </div>
                    <div class="p-8 mt-2 bg-blue-900 w-fit text-white">
                        je suis en 900
                    </div>
                </div>
                
                <div class={"ml-8"}>
                    <div class="p-8 mt-2 bg-blue-pastel-50 w-fit">
                        je suis en 50
                    </div>
                    <div class="p-8 mt-2 bg-blue-pastel-100 w-fit">
                        je suis en 100
                    </div>
                    <div class="p-8 mt-2 bg-blue-pastel-200 w-fit">
                        je suis en 200
                    </div>
                    <div class="p-8 mt-2 bg-blue-pastel-300 w-fit">
                        je suis en 300
                    </div>
                    <div class="p-8 mt-2 bg-blue-pastel-400 w-fit">
                        je suis en 400
                    </div>
                    <div class="p-8 mt-2 bg-blue-pastel-500 w-fit text-white">
                        je suis en 500
                    </div>
                    <div class="p-8 mt-2 bg-blue-pastel-600 w-fit text-white">
                        je suis en 600
                    </div>
                    <div class="p-8 mt-2 bg-blue-pastel-700 w-fit text-white">
                        je suis en 700
                    </div>
                    <div class="p-8 mt-2 bg-blue-pastel-800 w-fit text-white">
                        je suis en 800
                    </div>
                    <div class="p-8 mt-2 bg-blue-pastel-900 w-fit text-white">
                        je suis en 900
                    </div>
                </div>
                
                <div class={"ml-8"}>
                    <div class="p-8 mt-2 bg-red-50 w-fit">
                        je suis en 50
                    </div>
                    <div class="p-8 mt-2 bg-red-100 w-fit">
                        je suis en 100
                    </div>
                    <div class="p-8 mt-2 bg-red-200 w-fit">
                        je suis en 200
                    </div>
                    <div class="p-8 mt-2 bg-red-300 w-fit">
                        je suis en 300
                    </div>
                    <div class="p-8 mt-2 bg-red-400 w-fit">
                        je suis en 400
                    </div>
                    <div class="p-8 mt-2 bg-red-500 w-fit text-white">
                        je suis en 500
                    </div>
                    <div class="p-8 mt-2 bg-red-600 w-fit text-white">
                        je suis en 600
                    </div>
                    <div class="p-8 mt-2 bg-red-700 w-fit text-white">
                        je suis en 700
                    </div>
                    <div class="p-8 mt-2 bg-red-800 w-fit text-white">
                        je suis en 800
                    </div>
                    <div class="p-8 mt-2 bg-red-900 w-fit text-white">
                        je suis en 900
                    </div>
                </div>
                
                <div class={"ml-8"}>
                    <div class="p-8 mt-2 bg-red-pastel-50 w-fit">
                        je suis en 50
                    </div>
                    <div class="p-8 mt-2 bg-red-pastel-100 w-fit">
                        je suis en 100
                    </div>
                    <div class="p-8 mt-2 bg-red-pastel-200 w-fit">
                        je suis en 200
                    </div>
                    <div class="p-8 mt-2 bg-red-pastel-300 w-fit">
                        je suis en 300
                    </div>
                    <div class="p-8 mt-2 bg-red-pastel-400 w-fit">
                        je suis en 400
                    </div>
                    <div class="p-8 mt-2 bg-red-pastel-500 w-fit text-white">
                        je suis en 500
                    </div>
                    <div class="p-8 mt-2 bg-red-pastel-600 w-fit text-white">
                        je suis en 600
                    </div>
                    <div class="p-8 mt-2 bg-red-pastel-700 w-fit text-white">
                        je suis en 700
                    </div>
                    <div class="p-8 mt-2 bg-red-pastel-800 w-fit text-white">
                        je suis en 800
                    </div>
                    <div class="p-8 mt-2 bg-red-pastel-900 w-fit text-white">
                        je suis en 900
                    </div>
                </div>
                
                <div class={"ml-8"}>
                    <div class="p-8 mt-2 bg-green-50 w-fit">
                        je suis en 50
                    </div>
                    <div class="p-8 mt-2 bg-green-100 w-fit">
                        je suis en 100
                    </div>
                    <div class="p-8 mt-2 bg-green-200 w-fit">
                        je suis en 200
                    </div>
                    <div class="p-8 mt-2 bg-green-300 w-fit">
                        je suis en 300
                    </div>
                    <div class="p-8 mt-2 bg-green-400 w-fit">
                        je suis en 400
                    </div>
                    <div class="p-8 mt-2 bg-green-500 w-fit text-white">
                        je suis en 500
                    </div>
                    <div class="p-8 mt-2 bg-green-600 w-fit text-white">
                        je suis en 600
                    </div>
                    <div class="p-8 mt-2 bg-green-700 w-fit text-white">
                        je suis en 700
                    </div>
                    <div class="p-8 mt-2 bg-green-800 w-fit text-white">
                        je suis en 800
                    </div>
                    <div class="p-8 mt-2 bg-green-900 w-fit text-white">
                        je suis en 900
                    </div>
                </div>
                
                <div class={"ml-8"}>
                    <div class="p-8 mt-2 bg-green-pastel-50 w-fit">
                        je suis en 50
                    </div>
                    <div class="p-8 mt-2 bg-green-pastel-100 w-fit">
                        je suis en 100
                    </div>
                    <div class="p-8 mt-2 bg-green-pastel-200 w-fit">
                        je suis en 200
                    </div>
                    <div class="p-8 mt-2 bg-green-pastel-300 w-fit">
                        je suis en 300
                    </div>
                    <div class="p-8 mt-2 bg-green-pastel-400 w-fit">
                        je suis en 400
                    </div>
                    <div class="p-8 mt-2 bg-green-pastel-500 w-fit text-white">
                        je suis en 500
                    </div>
                    <div class="p-8 mt-2 bg-green-pastel-600 w-fit text-white">
                        je suis en 600
                    </div>
                    <div class="p-8 mt-2 bg-green-pastel-700 w-fit text-white">
                        je suis en 700
                    </div>
                    <div class="p-8 mt-2 bg-green-pastel-800 w-fit text-white">
                        je suis en 800
                    </div>
                    <div class="p-8 mt-2 bg-green-pastel-900 w-fit text-white">
                        je suis en 900
                    </div>
                </div>
                
                <div class={"ml-8"}>
                    <div class="p-8 mt-2 bg-yellow-50 w-fit">
                        je suis en 50
                    </div>
                    <div class="p-8 mt-2 bg-yellow-100 w-fit">
                        je suis en 100
                    </div>
                    <div class="p-8 mt-2 bg-yellow-200 w-fit">
                        je suis en 200
                    </div>
                    <div class="p-8 mt-2 bg-yellow-300 w-fit">
                        je suis en 300
                    </div>
                    <div class="p-8 mt-2 bg-yellow-400 w-fit">
                        je suis en 400
                    </div>
                    <div class="p-8 mt-2 bg-yellow-500 w-fit text-white">
                        je suis en 500
                    </div>
                    <div class="p-8 mt-2 bg-yellow-600 w-fit text-white">
                        je suis en 600
                    </div>
                    <div class="p-8 mt-2 bg-yellow-700 w-fit text-white">
                        je suis en 700
                    </div>
                    <div class="p-8 mt-2 bg-yellow-800 w-fit text-white">
                        je suis en 800
                    </div>
                    <div class="p-8 mt-2 bg-yellow-900 w-fit text-white">
                        je suis en 900
                    </div>
                </div>
                
                <div class={"ml-8"}>
                    <div class="p-8 mt-2 bg-yellow-pastel-50 w-fit">
                        je suis en 50
                    </div>
                    <div class="p-8 mt-2 bg-yellow-pastel-100 w-fit">
                        je suis en 100
                    </div>
                    <div class="p-8 mt-2 bg-yellow-pastel-200 w-fit">
                        je suis en 200
                    </div>
                    <div class="p-8 mt-2 bg-yellow-pastel-300 w-fit">
                        je suis en 300
                    </div>
                    <div class="p-8 mt-2 bg-yellow-pastel-400 w-fit">
                        je suis en 400
                    </div>
                    <div class="p-8 mt-2 bg-yellow-pastel-500 w-fit text-white">
                        je suis en 500
                    </div>
                    <div class="p-8 mt-2 bg-yellow-pastel-600 w-fit text-white">
                        je suis en 600
                    </div>
                    <div class="p-8 mt-2 bg-yellow-pastel-700 w-fit text-white">
                        je suis en 700
                    </div>
                    <div class="p-8 mt-2 bg-yellow-pastel-800 w-fit text-white">
                        je suis en 800
                    </div>
                    <div class="p-8 mt-2 bg-yellow-pastel-900 w-fit text-white">
                        je suis en 900
                    </div>
                </div>
                
                <div class={"ml-8"}>
                    <div class="p-8 mt-2 bg-orange-50 w-fit">
                        je suis en 50
                    </div>
                    <div class="p-8 mt-2 bg-orange-100 w-fit">
                        je suis en 100
                    </div>
                    <div class="p-8 mt-2 bg-orange-200 w-fit">
                        je suis en 200
                    </div>
                    <div class="p-8 mt-2 bg-orange-300 w-fit">
                        je suis en 300
                    </div>
                    <div class="p-8 mt-2 bg-orange-400 w-fit">
                        je suis en 400
                    </div>
                    <div class="p-8 mt-2 bg-orange-500 w-fit text-white">
                        je suis en 500
                    </div>
                    <div class="p-8 mt-2 bg-orange-600 w-fit text-white">
                        je suis en 600
                    </div>
                    <div class="p-8 mt-2 bg-orange-700 w-fit text-white">
                        je suis en 700
                    </div>
                    <div class="p-8 mt-2 bg-orange-800 w-fit text-white">
                        je suis en 800
                    </div>
                    <div class="p-8 mt-2 bg-orange-900 w-fit text-white">
                        je suis en 900
                    </div>
                </div>
                
                <div class={"ml-8"}>
                    <div class="p-8 mt-2 bg-orange-pastel-50 w-fit">
                        je suis en 50
                    </div>
                    <div class="p-8 mt-2 bg-orange-pastel-100 w-fit">
                        je suis en 100
                    </div>
                    <div class="p-8 mt-2 bg-orange-pastel-200 w-fit">
                        je suis en 200
                    </div>
                    <div class="p-8 mt-2 bg-orange-pastel-300 w-fit">
                        je suis en 300
                    </div>
                    <div class="p-8 mt-2 bg-orange-pastel-400 w-fit">
                        je suis en 400
                    </div>
                    <div class="p-8 mt-2 bg-orange-pastel-500 w-fit text-white">
                        je suis en 500
                    </div>
                    <div class="p-8 mt-2 bg-orange-pastel-600 w-fit text-white">
                        je suis en 600
                    </div>
                    <div class="p-8 mt-2 bg-orange-pastel-700 w-fit text-white">
                        je suis en 700
                    </div>
                    <div class="p-8 mt-2 bg-orange-pastel-800 w-fit text-white">
                        je suis en 800
                    </div>
                    <div class="p-8 mt-2 bg-orange-pastel-900 w-fit text-white">
                        je suis en 900
                    </div>
                </div>
                
                <div class={"ml-8"}>
                    <div class="p-8 mt-2 bg-orange-50 w-fit">
                        je suis en 50
                    </div>
                    <div class="p-8 mt-2 bg-orange-100 w-fit">
                        je suis en 100
                    </div>
                    <div class="p-8 mt-2 bg-orange-200 w-fit">
                        je suis en 200
                    </div>
                    <div class="p-8 mt-2 bg-orange-300 w-fit">
                        je suis en 300
                    </div>
                    <div class="p-8 mt-2 bg-orange-400 w-fit">
                        je suis en 400
                    </div>
                    <div class="p-8 mt-2 bg-orange-500 w-fit text-white">
                        je suis en 500
                    </div>
                    <div class="p-8 mt-2 bg-orange-600 w-fit text-white">
                        je suis en 600
                    </div>
                    <div class="p-8 mt-2 bg-orange-700 w-fit text-white">
                        je suis en 700
                    </div>
                    <div class="p-8 mt-2 bg-orange-800 w-fit text-white">
                        je suis en 800
                    </div>
                    <div class="p-8 mt-2 bg-orange-900 w-fit text-white">
                        je suis en 900
                    </div>
                </div>
                
                <div class={"ml-8"}>
                    <div class="p-8 mt-2 bg-orange-pastel-50 w-fit">
                        je suis en 50
                    </div>
                    <div class="p-8 mt-2 bg-orange-pastel-100 w-fit">
                        je suis en 100
                    </div>
                    <div class="p-8 mt-2 bg-orange-pastel-200 w-fit">
                        je suis en 200
                    </div>
                    <div class="p-8 mt-2 bg-orange-pastel-300 w-fit">
                        je suis en 300
                    </div>
                    <div class="p-8 mt-2 bg-orange-pastel-400 w-fit">
                        je suis en 400
                    </div>
                    <div class="p-8 mt-2 bg-orange-pastel-500 w-fit text-white">
                        je suis en 500
                    </div>
                    <div class="p-8 mt-2 bg-orange-pastel-600 w-fit text-white">
                        je suis en 600
                    </div>
                    <div class="p-8 mt-2 bg-orange-pastel-700 w-fit text-white">
                        je suis en 700
                    </div>
                    <div class="p-8 mt-2 bg-orange-pastel-800 w-fit text-white">
                        je suis en 800
                    </div>
                    <div class="p-8 mt-2 bg-orange-pastel-900 w-fit text-white">
                        je suis en 900
                    </div>
                </div>
                
                <div class={"ml-8"}>
                    <div class="p-8 mt-2 bg-violet-50 w-fit">
                        je suis en 50
                    </div>
                    <div class="p-8 mt-2 bg-violet-100 w-fit">
                        je suis en 100
                    </div>
                    <div class="p-8 mt-2 bg-violet-200 w-fit">
                        je suis en 200
                    </div>
                    <div class="p-8 mt-2 bg-violet-300 w-fit">
                        je suis en 300
                    </div>
                    <div class="p-8 mt-2 bg-violet-400 w-fit">
                        je suis en 400
                    </div>
                    <div class="p-8 mt-2 bg-violet-500 w-fit text-white">
                        je suis en 500
                    </div>
                    <div class="p-8 mt-2 bg-violet-600 w-fit text-white">
                        je suis en 600
                    </div>
                    <div class="p-8 mt-2 bg-violet-700 w-fit text-white">
                        je suis en 700
                    </div>
                    <div class="p-8 mt-2 bg-violet-800 w-fit text-white">
                        je suis en 800
                    </div>
                    <div class="p-8 mt-2 bg-violet-900 w-fit text-white">
                        je suis en 900
                    </div>
                </div>
                
                <div class={"ml-8"}>
                    <div class="p-8 mt-2 bg-violet-pastel-50 w-fit">
                        je suis en 50
                    </div>
                    <div class="p-8 mt-2 bg-violet-pastel-100 w-fit">
                        je suis en 100
                    </div>
                    <div class="p-8 mt-2 bg-violet-pastel-200 w-fit">
                        je suis en 200
                    </div>
                    <div class="p-8 mt-2 bg-violet-pastel-300 w-fit">
                        je suis en 300
                    </div>
                    <div class="p-8 mt-2 bg-violet-pastel-400 w-fit">
                        je suis en 400
                    </div>
                    <div class="p-8 mt-2 bg-violet-pastel-500 w-fit text-white">
                        je suis en 500
                    </div>
                    <div class="p-8 mt-2 bg-violet-pastel-600 w-fit text-white">
                        je suis en 600
                    </div>
                    <div class="p-8 mt-2 bg-violet-pastel-700 w-fit text-white">
                        je suis en 700
                    </div>
                    <div class="p-8 mt-2 bg-violet-pastel-800 w-fit text-white">
                        je suis en 800
                    </div>
                    <div class="p-8 mt-2 bg-violet-pastel-900 w-fit text-white">
                        je suis en 900
                    </div>
                </div>
            </div>
            <div class={"ml-4 mt-8"}>
                <Card variant="long" class="">
                    <p>Coucou je suis un test de card</p>
                    <p>Coucou je suis un test de card</p>
                    
                    <Badge variant="blue" class="mt-8">en cours</Badge>
                    <Badge variant="red" class="mt-8">en cours</Badge>
                    <Badge variant="green" class="mt-8">en cours</Badge>
                    <Badge variant="yellow" class="mt-8">en cours</Badge>
                    <Badge variant="purple" class="mt-8">en cours</Badge>
                    <Badge variant="orange" class="mt-8">en cours</Badge>
                    
                    <div class={"mt-8"}></div>
                    
                    <LinkText href="/">Lien interne vers l'accueil</LinkText>
                </Card>
            </div>
            
            <div class="mt-8">
                <Card variant="tall" containPosition="centerColumn" class="ml-4">
                    <p>Coucou je suis un test de card</p>
                    <p>Coucou je suis un test de card</p>
                    
                    <div class={"px-6 py-4 hover:cursor-pointer font-semibold bg-blue-500 hover:bg-blue-600 text-white rounded-lg w-fit mt-8"}>
                        Dashboard
                    </div>
                    <div class={"mb-4 uppercase font-semibold hover:cursor-pointer mt-4 px-6 py-4 bg-blue-500 hover:bg-blue-600 text-white rounded-lg w-fit"}>
                        Dashboard
                    </div>
                    
                    <div class={"px-4 py-2 hover:cursor-pointer font-semibold bg-red-500 hover:bg-red-600 text-white rounded-lg w-fit mt-8"}>
                        Dashboard
                    </div>
                    <div class={"mb-4 uppercase font-semibold hover:cursor-pointer mt-4 px-4 py-4 bg-red-500 hover:bg-red-600 text-white rounded-lg w-fit"}>
                        Dashboard
                    </div>
                    
                    <div class={"px-6 py-4 hover:cursor-pointer font-semibold bg-green-500 hover:bg-green-600 text-white rounded-lg w-fit mt-8"}>
                        Dashboard
                    </div>
                    <div class={"mb-4 uppercase font-semibold hover:cursor-pointer mt-4 px-6 py-4 bg-green-500 hover:bg-green-600 text-white rounded-lg w-fit"}>
                        Dashboard
                    </div>
                    
                    <div class={"px-4 py-2 hover:cursor-pointer font-semibold bg-yellow-500 hover:bg-yellow-600 text-white rounded-lg w-fit mt-8"}>
                        Dashboard
                    </div>
                    <div class={"mb-4 uppercase font-semibold hover:cursor-pointer mt-4 px-4 py-4 bg-yellow-500 hover:bg-yellow-600 text-white rounded-lg w-fit"}>
                        Dashboard
                    </div>
                    
                    <div class={"px-6 py-4 hover:cursor-pointer font-semibold bg-orange-500 hover:bg-orange-600 text-white rounded-lg w-fit mt-8"}>
                        Dashboard
                    </div>
                    <div class={"mb-4 uppercase font-semibold hover:cursor-pointer mt-4 px-6 py-4 bg-orange-500 hover:bg-orange-600 text-white rounded-lg w-fit"}>
                        Dashboard
                    </div>
                    
                    <div class={"px-4 py-2 hover:cursor-pointer font-semibold bg-violet-500 hover:bg-violet-600 text-white rounded-lg w-fit mt-8"}>
                        Dashboard
                    </div>
                    <div class={"mb-4 uppercase font-semibold hover:cursor-pointer mt-4 px-4 py-4 bg-violet-500 hover:bg-violet-600 text-white rounded-lg w-fit"}>
                        Dashboard
                    </div>
                    
                    <div class={"px-6 py-4 hover:cursor-pointer font-semibold bg-pink-500 hover:bg-pink-600 text-white rounded-lg w-fit mt-8"}>
                        Dashboard
                    </div>
                    <div class={"mb-4 uppercase font-semibold hover:cursor-pointer mt-4 px-6 py-4 bg-pink-500 hover:bg-pink-600 text-white rounded-lg w-fit"}>
                        Dashboard
                    </div>
                    
                    <div class={"px-4 py-2 hover:cursor-pointer font-semibold bg-indigo-500 hover:bg-indigo-600 text-white rounded-lg w-fit mt-8"}>
                        Dashboard
                    </div>
                    <div class={"mb-4 uppercase font-semibold hover:cursor-pointer mt-4 px-4 py-4 bg-indigo-500 hover:bg-indigo-600 text-white rounded-lg w-fit"}>
                        Dashboard
                    </div>
                    <LinkText href="/">Je suis encore un link text</LinkText>
                </Card>
            </div>
            
            <Card tag="article" variant="small" class="ml-8 mt-8">
                <div>Test card</div>
            </Card>
            <Card tag="section" variant="large" class="ml-8 mt-8">
                <div>Test card</div>
            </Card>
            <Card variant="extraLarge" class="ml-8 mt-8">
                <div>Test card</div>
            </Card>
            
            <Card variant="long" class={"mt-10 ml-8"}>
                <h2 class="text-lg font-semibold mb-4">Tests du composant Button</h2>
                
                <div class="space-y-4">
                    <div class="flex gap-4 flex-wrap">
                        <Button onClick$={() => console.log('Click! btn bleu')}>Button par défaut</Button>
                        <Button disabled onClick$={() => console.log('Click! btn bleu')}>Button par défaut desactiver</Button>
                        <div>
                            <Button size="sm" onClick$={() => console.log('Click! btn vert')} variant="success">Success Small</Button>
                        </div>
                        <div>
                            <Button size="sm" uppercase variant="error">Error Small Uppercase</Button>
                        </div>
                        <div>
                            <Button variant="warning">Warning Base</Button>
                        </div>
                        <div>
                            <Button uppercase variant="warning">Warning High Uppercase</Button>
                        </div>
                    </div>
                    
                    <div class="min-h-[40px]">
                        <form class="flex gap-4 items-start" onSubmit$={handler} preventdefault:submit>
                            <Button type="button" onClick$={() => console.log('Click! btn vert')} variant="success">Success not submit</Button>
                            <Button type="submit" size="sm" variant="success">Success submit</Button>
                        </form>
                    </div>
                    
                    <div class="flex gap-4 flex-wrap">
                        <div>
                            <Button size="sm" uppercase variant="purple">Purple Small</Button>
                        </div>
                        <div>
                            <Button size="sm" uppercase variant="pink">Pink Small Uppercase</Button>
                        </div>
                        <Button variant="indigo">Indigo Base</Button>
                        <Button variant="primary">Base Uppercase</Button>
                    </div>
                    
                    <div class="flex gap-4 flex-wrap">
                        <Button variant="success" class="mr-4">Custom Class</Button>
                        <Button variant="error">Gros bouton erreur</Button>
                    </div>
                    
                    <div class="space-y-4">
                        <h3 class="text-md font-semibold">Tests du composant Badge</h3>
                        <div class="flex gap-2 flex-wrap items-center">
                            <Badge variant="blue">Bleu par défaut</Badge>
                            <Badge variant="red" size="sm">Rouge small</Badge>
                            <Badge variant="green">Vert</Badge>
                            <Badge variant="yellow" size="sm">Jaune small</Badge>
                            <Badge variant="purple">Violet</Badge>
                            <Badge variant="orange">Orange</Badge>
                            <Badge variant="gray">Gris</Badge>
                            <Badge variant="indigo">Indigo</Badge>
                            <Badge variant="pink" size="sm">Rose small</Badge>
                        </div>
                    </div>
                </div>
                
                <div class="mb-10"></div>
                
                {/* Input basique */}
                
                <section class="mb-10">
                    <h2 class="mb-6 text-xl font-semibold text-gray-800">Input basique</h2>
                    <Input
                        label="Nom complet"
                        placeholder="Entrez votre nom"
                    />
                </section>
                
                <section class="w-[50%] mb-10">
                    <h2 class="mb-6 text-xl font-semibold text-gray-800">Input basique 50% de large</h2>
                    <Input
                        label="Nom complet"
                        placeholder="Entrez votre nom"
                    />
                </section>
                
                <section class="mb-10">
                    <h2 class="mb-6 text-xl font-semibold text-gray-800">Input filled</h2>
                    <Input
                        label="Email"
                        placeholder="Entrez votre email"
                        variant="filled"
                    />
                </section>
                
                <section class="mb-10">
                    <h2 class="mb-6 text-xl font-semibold text-gray-800">Input désactivé</h2>
                    <Input
                        label="Adresse"
                        placeholder="Entrez votre adresse"
                        state="disabled"
                    />
                </section>
                
                <section class="mb-10">
                    <h2 class="mb-6 text-xl font-semibold text-gray-800">Input avec icône</h2>
                    <Input
                        label="Recherche"
                        placeholder="Rechercher..."
                        leftIcon={<LockCloseIcon/>}
                    />
                </section>
                
                <section class="mb-10">
                    <h2 class="mb-6 text-xl font-semibold text-gray-800">Input outlined</h2>
                    <Input
                        variant="outlined"
                        label="Mot de passe"
                        placeholder="Entrez votre mot de passe"
                    />
                </section>
                
                <section class="mb-10">
                    <h2 class="mb-6 text-xl font-semibold text-gray-800">Input mini</h2>
                    <Input
                        variant="minimal"
                        label="Téléphone"
                        placeholder="Entrez votre numéro de téléphone"
                    />
                </section>
                
                <section class="mb-10">
                    <h2 class="mb-6 text-xl font-semibold text-gray-800">Input avec message d'erreur</h2>
                    <Input
                        label="Code postal"
                        placeholder="Entrez votre code postal"
                        state="error"
                        error="Code postal invalide"
                    />
                </section>
                
                <section class="mb-10">
                    <h2 class="mb-6 text-xl font-semibold text-gray-800">Input outlined succes</h2>
                    <Input
                        variant="outlined"
                        label="Date de naissance"
                        placeholder="Entrez votre date de naissance"
                        state="success"
                    />
                </section>
                
                <section class="mb-10">
                    <h2 class="mb-6 text-xl font-semibold text-gray-800">Input avec message d'aide</h2>
                    <Input
                        label="Ville"
                        placeholder="Entrez votre ville"
                        helper="Exemple : Paris, Lyon, Marseille"
                    />
                </section>
                
                <section class="mb-10">
                    <h2 class="mb-6 text-xl font-semibold text-gray-800">Input avec message d'aide et icône</h2>
                    <Input
                        label="Code de sécurité"
                        placeholder="Entrez votre code de sécurité"
                        helper="Exemple : 1234"
                        leftIcon={<LockCloseIcon/>}
                    />
                </section>
                
                <section class="mb-10">
                    <h2 class="mb-6 text-xl font-semibold text-gray-800">Input required</h2>
                    <Input
                        label="Numéro de carte"
                        placeholder="Entrez votre numéro de carte"
                        required
                    />
                </section>
            </Card>
            
            <div class="mt-12 mb-12">X</div>
        </>
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
