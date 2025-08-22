import { useStore, useSignal, $, QRL } from "@builder.io/qwik";

// Types pour les règles de validation
type ValidationRule<T> = {
    required?: boolean;
    minLength?: number;
    maxLength?: number;
    pattern?: RegExp;
    email?: boolean;
    custom?: (value: T) => string | undefined; // Pas de QRL pour les fonctions synchrones
};

// Configuration du formulaire
type FormConfig<T extends Record<string, any>> = {
    initialValues: T;
    validationRules?: Partial<Record<keyof T, ValidationRule<T[keyof T]>>>;
};

// Type de retour du hook
type UseFormReturn<T extends Record<string, any>> = {
    values: T;
    errors: Partial<Record<keyof T, string>>;
    touched: Partial<Record<keyof T, boolean>>;
    isValid: { value: boolean };
    setFieldValue: QRL<(fieldName: keyof T, value: any) => void>;
    setFieldError: QRL<(fieldName: keyof T, error: string) => void>;
    clearFieldError: QRL<(fieldName: keyof T) => void>;
    validateForm: QRL<() => boolean>;
    reset: QRL<() => void>;
    getFieldProps: QRL<(fieldName: keyof T) => {
        name: string;
        value: any;
        error?: string;
        onInput$: QRL<(event: Event) => void>;
        onBlur$: QRL<() => void>;
    }>;
};

export function useForm<T extends Record<string, any>>(
    config: FormConfig<T>
): UseFormReturn<T> {
    // Stores réactifs
    const values = useStore<T>({ ...config.initialValues });
    const errors = useStore<Partial<Record<keyof T, string>>>({});
    const touched = useStore<Partial<Record<keyof T, boolean>>>({});
    const isValid = useSignal(true);
    
    // Validation d'un champ (fonction synchrone)
    const validateField = $((fieldName: keyof T, value: any): string | undefined => {
        const rules = config.validationRules?.[fieldName];
        if (!rules) return undefined;
        
        // Required
        if (rules.required && (!value || (typeof value === 'string' && value.trim() === ''))) {
            return 'Ce champ est requis';
        }
        
        // Si le champ est vide et non requis, on skip les autres validations
        if (!value && !rules.required) return undefined;
        
        // Email
        if (rules.email && !/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(value)) {
            return 'Email invalide';
        }
        
        // Pattern
        if (rules.pattern && !rules.pattern.test(value)) {
            return 'Format invalide';
        }
        
        // Min length
        if (rules.minLength && value.length < rules.minLength) {
            return `Minimum ${rules.minLength} caractères`;
        }
        
        // Max length
        if (rules.maxLength && value.length > rules.maxLength) {
            return `Maximum ${rules.maxLength} caractères`;
        }
        
        // Custom validation
        if (rules.custom) {
            return rules.custom(value);
        }
        
        return undefined;
    });
    
    // Définir la valeur d'un champ avec validation
    const setFieldValue = $((fieldName: keyof T, value: any) => {
        values[fieldName] = value;
        
        // Validation en temps réel si le champ a déjà été touché
        if (touched[fieldName]) {
            const error = validateField(fieldName, value);
            if (error) {
                error.then(err => {
                    errors[fieldName] = err;
                });
            } else {
                delete errors[fieldName];
            }
        }
        
        // Mettre à jour l'état de validité globale
        updateFormValidity();
    });
    
    // Définir une erreur manuellement (pour les erreurs serveur)
    const setFieldError = $((fieldName: keyof T, error: string) => {
        errors[fieldName] = error;
        touched[fieldName] = true;
        updateFormValidity();
    });
    
    // Effacer l'erreur d'un champ
    const clearFieldError = $((fieldName: keyof T) => {
        delete errors[fieldName];
        updateFormValidity();
    });
    
    // Mettre à jour l'état de validité du formulaire
    const updateFormValidity = $(() => {
        isValid.value = Object.keys(errors).length === 0;
    });
    
    // Valider tout le formulaire
    const validateForm = $((): boolean => {
        const fieldNames = Object.keys(values) as (keyof T)[];
        let formIsValid = true;
        
        // Reset des erreurs
        Object.keys(errors).forEach(key => delete errors[key as keyof T]);
        
        // Validation de tous les champs
        for (const fieldName of fieldNames) {
            const error = validateField(fieldName, values[fieldName]);
            if (error) {
                error.then(err => {
                    errors[fieldName] = err
                });
                formIsValid = false;
            }
        }
        
        // Marquer tous les champs comme touchés
        fieldNames.forEach(fieldName => {
            touched[fieldName] = true;
        });
        
        isValid.value = formIsValid;
        return formIsValid;
    });
    
    // Reset du formulaire
    const reset = $(() => {
        Object.assign(values, config.initialValues);
        Object.keys(errors).forEach(key => delete errors[key as keyof T]);
        Object.keys(touched).forEach(key => delete touched[key as keyof T]);
        isValid.value = true;
    });
    
    // Helper pour obtenir les props d'un champ
    const getFieldProps = $((fieldName: keyof T) => {
        return {
            name: String(fieldName), // Important pour les formulaires HTML
            value: values[fieldName] || '',
            error: touched[fieldName] ? errors[fieldName] : undefined,
            onInput$: $((event: Event) => {
                const target = event.target as HTMLInputElement;
                setFieldValue(fieldName, target.value);
            }),
            onBlur$: $(() => {
                touched[fieldName] = true;
                const error = validateField(fieldName, values[fieldName]);
                if (error) {
                    error.then(err => {
                        errors[fieldName] = err;
                    });
                } else {
                    delete errors[fieldName];
                }
                updateFormValidity();
            })
        };
    });
    
    return {
        values,
        errors,
        touched,
        isValid,
        setFieldValue,
        setFieldError,
        clearFieldError,
        validateForm,
        reset,
        getFieldProps
    };
}
