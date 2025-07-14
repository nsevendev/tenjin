# Déploiement Automatique avec GitHub Actions

## Configuration des Secrets GitHub

Pour que le déploiement automatique fonctionne, vous devez ajouter ces secrets dans votre repository GitHub :

### 1. Accéder aux Secrets
- Allez dans `Settings` > `Secrets and variables` > `Actions`
- Cliquez sur `New repository secret`

### 2. Secrets Requis

#### `IONOS_SSH_KEY`
- **Description** : Clé SSH privée pour se connecter au serveur Ionos
- **Valeur** : Contenu de votre fichier `~/.ssh/id_rsa` (clé privée)
- **Génération** :
  ```bash
  # Sur votre machine locale
  ssh-keygen -t rsa -b 4096 -C "github-actions@tenjin"
  # Copiez le contenu de ~/.ssh/id_rsa (clé PRIVÉE)
  cat ~/.ssh/id_rsa
  ```

#### `IONOS_HOST`
- **Description** : Adresse IP ou hostname du serveur Ionos
- **Valeur** : `votre-server.ionos.com` ou l'IP de votre serveur

#### `IONOS_USER`
- **Description** : Nom d'utilisateur SSH sur le serveur
- **Valeur** : `nseven` (ou votre nom d'utilisateur)

### 3. Configuration SSH sur le serveur Ionos

Sur votre serveur Ionos, ajoutez la clé publique :

```bash
# Sur le serveur Ionos
mkdir -p ~/.ssh
echo "VOTRE_CLE_PUBLIQUE_GITHUB" >> ~/.ssh/authorized_keys
chmod 600 ~/.ssh/authorized_keys
chmod 700 ~/.ssh
```

## Workflow de Déploiement

### Déclenchement
- **Push sur branche `preprod`** : Déploiement automatique
- **Manuel** : Via l'onglet Actions de GitHub

### Étapes du Déploiement
1. **Checkout** du code
2. **Configuration SSH** vers Ionos
3. **Git pull** sur le serveur
4. **Configuration .env** pour preprod
5. **Rebuild Docker** avec `make upb`
6. **Vérification** des services

### Environnements Configurés

#### Preprod
- **URLs** : 
  - Frontend: `https://tenjin.woopear.fr`
  - API: `https://tenjin-api.woopear.fr`
- **Base de données** : `tenjin_preprod`
- **Branche** : `preprod`

## Utilisation

### Créer la branche preprod
```bash
git checkout -b preprod
git push -u origin preprod
```

### Déployer en preprod
```bash
git checkout preprod
git merge main  # ou vos changements
git push origin preprod  # Déclenche le déploiement automatique
```

### Surveillance
- Consultez l'onglet `Actions` de GitHub pour voir les logs
- Les erreurs de déploiement sont visibles dans les logs d'action

## Dépannage

### Erreur SSH
- Vérifiez que la clé SSH est correcte dans les secrets
- Testez la connexion manuellement : `ssh nseven@votre-serveur.ionos.com`

### Erreur Docker
- Connectez-vous au serveur et vérifiez : `make logs-app`
- Vérifiez l'espace disque : `df -h`

### Variables d'environnement
- Les fichiers `.env` sont automatiquement configurés par l'Action
- Modifiez le workflow si vous ajoutez de nouvelles variables