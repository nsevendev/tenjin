import {routeAction$, validator$, ValidatorReturn} from "@builder.io/qwik-city";
import {ValueOrPromise} from "@builder.io/qwik";

// eslint-disable-next-line qwik/loader-location
export const useLoginAction = routeAction$(
    async (data, { fail, cookie }) => {
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
                if (response.status === 401) {
                    return fail(401, {
                        fieldErrors: {
                            email: 'Email ou mot de passe incorrect',
                            password: 'Email ou mot de passe incorrect'
                        },
                        message: 'Identifiants incorrects'
                    });
                }
                
                if (response.status === 404) {
                    return fail(404, {
                        fieldErrors: { email: 'Aucun compte trouvé avec cet email' },
                        message: 'Compte introuvable'
                    });
                }
                
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
                
                cookie.set(cookieName, cookieValue, {
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
            return fail(500, {
                message: 'Une erreur est survenue lors de la connexion'
            });
        }
    },
    validator$((ev, data): ValueOrPromise<ValidatorReturn> => {
        const formData = data as Record<string, unknown>;
        const errors: Record<string, string> = {};
        
        if (!formData.email || typeof formData.email !== 'string') {
            errors.email = "Email requis";
        } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(formData.email)) {
            errors.email = "Email invalide";
        }
        
        if (!formData.password || typeof formData.password !== 'string' || formData.password.length < 1) {
            errors.password = "Mot de passe requis";
        }
        
        if (Object.keys(errors).length > 0) {
            return {
                success: false,
                status: 400,
                error: errors
            };
        }
        
        return { success: true };
    })
);
