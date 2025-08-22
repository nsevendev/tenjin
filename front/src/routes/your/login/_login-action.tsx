import {routeAction$, z, zod$} from "@builder.io/qwik-city";

// Schéma de validation pour le serveur
const LoginSchema = z.object({
    email: z.string().email("Email invalide"),
    password: z.string().min(1, "Mot de passe requis")
});

// eslint-disable-next-line qwik/loader-location
export const useLoginAction = routeAction$(async (data, { fail, cookie }) => {
    console.log('📥 Données reçues:', data); // Pour debug
    console.log('📥 Type:', typeof data);
    console.log('📥 Keys:', Object.keys(data));
    
    try {
        // Validation côté serveur
        const validatedData = LoginSchema.parse(data);
        
        const response = await fetch(`http://api:3000/api/v1/your/login`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                email: validatedData.email,
                password: validatedData.password,
            })
        });
        
        console.log(response);
        
        if (!response.ok) {
            const errorData = await response.json();
            console.log(errorData);
            // Gestion des erreurs spécifiques
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
        console.log(result);
        // ✅ Gestion du cookie JWT (si nécessaire côté Qwik)
        const setCookieHeader = response.headers.get('set-cookie');
        if (setCookieHeader) {
            console.log('🍪 Cookie reçu:', setCookieHeader);
            
            // Parser le cookie (format: nom=valeur; options...)
            const cookieParts = setCookieHeader.split(';');
            const [nameValue] = cookieParts[0].split('=');
            const cookieName = nameValue;
            const cookieValue = cookieParts[0].split('=')[1];
            
            // ✅ Définir le cookie côté Qwik pour le navigateur
            cookie.set(cookieName, cookieValue, {
                maxAge: 86400, // 24h
                path: '/',
                httpOnly: true,
                sameSite: 'strict',
                secure: false // true en production avec HTTPS
            });
            
            console.log('🍪 Cookie défini pour le navigateur');
        }
        
        // Redirection après connexion réussie
        return {
            success: true,
            user: result.data,
            message: result.message || 'Connexion réussie !'
        };
        
    } catch (error) {
        console.error('Erreur lors de la connexion:', error);
        
        if (error instanceof z.ZodError) {
            return fail(400, {
                fieldErrors: error.flatten().fieldErrors,
                message: 'Données invalides'
            });
        }
        
        return fail(500, {
            message: 'Une erreur est survenue lors de la connexion'
        });
    }
}, zod$(LoginSchema));

