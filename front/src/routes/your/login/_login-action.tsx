import {routeAction$} from "@builder.io/qwik-city";

// eslint-disable-next-line qwik/loader-location
export const useLoginAction = routeAction$(async (data, requestEvent) => {
        console.log('📥 Données reçues:', data);
        
        try {
            const response = await fetch(`http://api:3000/api/v1/your/login`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    email: data.email,
                    password: data.password,
                })
            });
            
            if (!response.ok) {
                const errorData = await response.json();
                console.log(errorData);
                throw new Error('Erreur serveur');
            }
            
            const result = await response.json();
            
            // Gestion du cookie JWT
            const setCookieHeader = response.headers.get('set-cookie');
            if (setCookieHeader) {
                const cookieParts = setCookieHeader.split(';');
                const [nameValue] = cookieParts[0].split('=');
                const cookieName = nameValue;
                const cookieValue = cookieParts[0].split('=')[1];
                
                requestEvent.cookie.set(cookieName, cookieValue, {
                    maxAge: 86400,
                    path: '/',
                    httpOnly: true,
                    sameSite: 'strict',
                    secure: false
                });
            }
            
            return {
                success: true,
                user: result.data,
                message: result.message || 'Connexion réussie !'
            };
            
        } catch (error) {
            console.error('Erreur lors de la connexion:', error);
            return requestEvent.fail(500, {
                message: 'Une erreur est survenue lors de la connexion'
            });
        }
    },
);
