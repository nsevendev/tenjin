/**
 * Utilitaires de validation pour les formulaires et données
 */

export interface ValidationResult {
  isValid: boolean;
  errors: string[];
}

/**
 * Valide un mot de passe
 */
export function validatePassword(password: string): ValidationResult {
  const errors: string[] = [];
  
  if (!password) {
    errors.push('Mot de passe requis');
  }
  if (password.length < 8) {
    errors.push('Mot de passe trop court (minimum 8 caractères)');
  }
  if (password.length > 128) {
    errors.push('Mot de passe trop long');
  }
  if (!/[A-Z]/.test(password)) {
    errors.push('Doit contenir au moins une majuscule');
  }
  if (!/[a-z]/.test(password)) {
    errors.push('Doit contenir au moins une minuscule');
  }
  if (!/[0-9]/.test(password)) {
    errors.push('Doit contenir au moins un chiffre');
  }
  
  return {
    isValid: errors.length === 0,
    errors
  };
}
