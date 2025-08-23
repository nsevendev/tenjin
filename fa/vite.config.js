import { defineConfig } from 'vite';

export default defineConfig({
  server: {
    allowedHosts: [
      'tenjin.local',
      'localhost',
      '.local'  // Permet tous les sous-domaines .local
    ],
    host: '0.0.0.0',
    port: 3000,
    hmr: {
      host: 'tenjin.local'  // Pour le hot-reload via Traefik
    }
  }
});
