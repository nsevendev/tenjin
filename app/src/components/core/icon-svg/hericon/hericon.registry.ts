import { IconRegistry } from './hericon.type';

/**
 * Tous les paths viennent de https://heroicons.com (outline/solid 24)
 * Ajoute d’autres icônes en copiant-colle leur `d` ici. (le d du <path>)
 */
export const ICONS: IconRegistry = {
  // ——— CORE / NAV
  home: {
    variant: 'outline',
    paths: ['M2.25 12l9.39-9.39a1.5 1.5 0 012.12 0L23.25 12M4.5 9.75V21a.75.75 0 00.75.75H9.75V15a.75.75 0 01.75-.75h3a.75.75 0 01.75.75v6.75h4.5A.75.75 0 0020.25 21V9.75'],
  },
  dashboard: {
    variant: 'outline',
    paths: ['M3.75 3.75h6.5v6.5h-6.5zM13.75 3.75h6.5v10.5h-6.5zM3.75 13.75h6.5v6.5h-6.5zM13.75 16.75h6.5v3.5h-6.5z'],
  },
  search: {
    variant: 'outline',
    paths: ['M21 21l-4.35-4.35M10.5 18a7.5 7.5 0 100-15 7.5 7.5 0 000 15z'],
  },
  settings: {
    variant: 'outline',
    paths: ['M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.54-.89 3.31.88 2.42 2.42a1.724 1.724 0 001.066 2.573c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.89 1.54-.88 3.31-2.42 2.42a1.724 1.724 0 00-2.573 1.066c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.54.89-3.31-.88-2.42-2.42A1.724 1.724 0 003.317 13.675c-1.756-.426-1.756-2.924 0-3.35A1.724 1.724 0 004.383 7.75c-.89-1.54.88-3.31 2.42-2.42.958.554 2.149.214 2.573-1.013zM12 9a3 3 0 110 6 3 3 0 010-6z'],
  },
  bell: {
    variant: 'outline',
    paths: ['M14.25 18.75a2.25 2.25 0 11-4.5 0M4.5 15h15l-1.5-3V9a6 6 0 10-12 0v3L4.5 15z'],
  },

  // ——— UTILITAIRES
  plus: { variant: 'outline', paths: ['M12 4.5v15M4.5 12h15'] },
  minus: { variant: 'outline', paths: ['M5 12h14'] },
  x: { variant: 'outline', paths: ['M6 18L18 6M6 6l12 12'] },
  check: { variant: 'outline', paths: ['M4.5 12.75l6 6 9-13.5'] },
  edit: { variant: 'outline', paths: ['M16.862 4.487a2.25 2.25 0 113.182 3.183L7.5 20.25 3 21l.75-4.5 13.112-12.013z'] },
  trash: { variant: 'outline', paths: ['M6 7.5h12M9 7.5V6a3 3 0 016 0v1.5M7.5 7.5l1 12a2 2 0 001.995 1.8h5.01A2 2 0 0017.5 19.5l1-12'] },
  eye: { variant: 'outline', paths: ['M2.25 12s3.75-6.75 9.75-6.75S21.75 12 21.75 12 18 18.75 12 18.75 2.25 12 2.25 12z M12 9.75a2.25 2.25 0 110 4.5 2.25 2.25 0 010-4.5z'] },
  'eye-slash': { variant: 'outline', paths: ['M3 3l18 18M9.88 9.88A3 3 0 0012 15a3 3 0 002.12-.88M6.75 6.75C4.5 8.25 2.25 12 2.25 12s3.75 6.75 9.75 6.75a10.5 10.5 0 005.25-1.41'] },

  // ——— FICHIERS / DOCS
  document: {
    variant: 'outline',
    paths: ['M12 3H6.75A1.5 1.5 0 005.25 4.5v15A1.5 1.5 0 006.75 21h10.5A1.5 1.5 0 0018.75 19.5V9L12 3z M12 3v6h6'],
  },
  'document-check': {
    variant: 'outline',
    paths: ['M9 12l2 2 4-4 M12 3H6.75A1.5 1.5 0 005.25 4.5v15A1.5 1.5 0 006.75 21h10.5A1.5 1.5 0 0018.75 19.5V9L12 3z M12 3v6h6'],
  },
  'clipboard-document': {
    variant: 'outline',
    paths: ['M9 2.25h6v3H9z M7.5 4.5h9A1.5 1.5 0 0118 6v12a1.5 1.5 0 01-1.5 1.5h-9A1.5 1.5 0 016 18V6A1.5 1.5 0 017.5 4.5z'],
  },
  'arrow-down-tray': {
    variant: 'outline',
    paths: ['M3 16.5h18M12 3v12m0 0l3.75-3.75M12 15L8.25 11.25'],
  },

  // ——— FORMATION (cours, sessions, quiz, attestations, stages)
  'academic-cap': { // formation
    variant: 'outline',
    paths: ['M12 3l9 5.25-9 5.25L3 8.25 12 3z M12 13.5V21m-3-3h6'],
  },
  'presentation-chart-bar': { // suivi/analytics
    variant: 'outline',
    paths: ['M3 4.5h18v12a2.25 2.25 0 01-2.25 2.25H5.25A2.25 2.25 0 013 16.5v-12z M7.5 13.5V9.75M12 13.5V7.5M16.5 13.5V10.5'],
  },
  'book-open': { // cours
    variant: 'outline',
    paths: ['M12 6.75c-2.25-1.5-6-1.5-9 0v9c3-1.5 6-1.5 9 0m0-9c2.25-1.5 6-1.5 9 0v9c-3-1.5-6-1.5-9 0'],
  },
  'queue-list': { // sessions / modules
    variant: 'outline',
    paths: ['M3.75 6.75h16.5M3.75 12h16.5M3.75 17.25h10.5'],
  },
  'question-mark-circle': { // quiz
    variant: 'outline',
    paths: ['M8.625 9.75A3.375 3.375 0 0112 6.375 3.375 3.375 0 0115.375 9.75c0 1.875-2.25 2.25-2.25 3.75M12 18h.007'],
  },
  'certificate': { // attestation (rosette)
    variant: 'outline',
    paths: ['M12 3l1.5 3 3 .75-2.25 2.25.5 3-2.75-1.5L9.25 12l.5-3L7.5 6.75 10.5 6 12 3z M9 15l-3 6 4.5-3 4.5 3-3-6'],
  },
  'briefcase': { // stage
    variant: 'outline',
    paths: ['M9 6V4.5A1.5 1.5 0 0110.5 3h3A1.5 1.5 0 0115 4.5V6m-9 0h12A1.5 1.5 0 0119.5 7.5v9A1.5 1.5 0 0118 18H6a1.5 1.5 0 01-1.5-1.5v-9A1.5 1.5 0 016 6z'],
  },

  // ——— RECRUTEMENT (offres, candidats, matching, entretiens/vidéo)
  'building-office': { variant: 'outline', paths: ['M3 21h18M6 21V3h12v18M9 21v-3m6 3v-3M9 6h6M9 9h6M9 12h6'] },
  'user': { variant: 'outline', paths: ['M15.75 7.5a3.75 3.75 0 11-7.5 0 3.75 3.75 0 017.5 0z M4.5 20.25a7.5 7.5 0 0115 0'] },
  'user-group': { variant: 'outline', paths: ['M18 13.5a3 3 0 10-3-3m3 3v.75A4.5 4.5 0 0113.5 18H6a4.5 4.5 0 01-4.5-4.5V13.5m9-3a3 3 0 100-6'] },
  'sparkles': { variant: 'outline', paths: ['M9 4.5l1.5 3L14 9l-3.5 1.5L9 14l-1.5-3.5L4 9l3.5-1.5L9 4.5zM16.5 4.5l.75 1.5 1.5.75-1.5.75-.75 1.5-.75-1.5-1.5-.75 1.5-.75.75-1.5z'] }, // matching
  'video-camera': { variant: 'outline', paths: ['M3.75 7.5h10.5a1.5 1.5 0 011.5 1.5v1.5l4.5-3v9l-4.5-3V15a1.5 1.5 0 01-1.5 1.5H3.75A1.5 1.5 0 012.25 15V9A1.5 1.5 0 013.75 7.5z'] }, // entretiens
  'calendar': { variant: 'outline', paths: ['M6.75 3v3M17.25 3v3M3 8.25h18M4.5 7.5v12A1.5 1.5 0 006 21h12a1.5 1.5 0 001.5-1.5v-12'] },
  'document-text': { variant: 'outline', paths: ['M12 3v6h6M8.25 12h7.5M8.25 15.75h7.5M6.75 3h5.25l6 6V19.5A1.5 1.5 0 0117.25 21H6.75A1.5 1.5 0 015.25 19.5V4.5A1.5 1.5 0 016.75 3z'] }, // offre
  'tag': { variant: 'outline', paths: ['M2.25 12l9.75 9.75a1.5 1.5 0 002.121 0L21.75 14.13a1.5 1.5 0 000-2.121L12 2.25H6A3.75 3.75 0 002.25 6v6z'] }, // compétences

  // ——— MESSAGERIE / STREAMING
  'chat-bubble-left-right': { variant: 'outline', paths: ['M2.25 12.75a4.5 4.5 0 014.5-4.5h9a4.5 4.5 0 014.5 4.5v.75a3 3 0 01-3 3H9.75L6 19.5v-3H6a3 3 0 01-3-3v-.75z'] },
  'play-circle': { variant: 'outline', paths: ['M12 21a9 9 0 100-18 9 9 0 000 18z M10 9.75l5.25 2.25L10 14.25V9.75z'] },
  'pause-circle': { variant: 'outline', paths: ['M12 21a9 9 0 100-18 9 9 0 000 18z M10 9v6M14 9v6'] },

  // ——— SÉCURITÉ
  'lock-closed': { variant: 'outline', paths: ['M6.75 10.5V7.5A5.25 5.25 0 0117.25 7.5v3M6 10.5h12v9A1.5 1.5 0 0116.5 21h-9A1.5 1.5 0 016 19.5v-9z'] },
  'shield-check': { variant: 'outline', paths: ['M12 3l7.5 3v6c0 5.25-3.75 7.5-7.5 9-3.75-1.5-7.5-3.75-7.5-9V6L12 3z M9.75 12.75l1.5 1.5 3-3'] },

  // ——— ÉTAT / FEEDBACK
  'information-circle': { variant: 'outline', paths: ['M12 8.25h.007M11.25 12h1.5v4.5h-1.5z'] },
  'exclamation-triangle': { variant: 'outline', paths: ['M12 3.75L1.5 20.25h21L12 3.75z M12 9v4.5M12 16.5h.007'] },
  'arrow-path': { variant: 'outline', paths: ['M4.5 9A7.5 7.5 0 0121 9m0 0l-3-3m3 3l-3 3M3 15a7.5 7.5 0 0016.5 0m0 0l-3 3m3-3l-3-3'] },

  // ——— TABLEAU / LISTE / VUE
  'table-cells': { variant: 'outline', paths: ['M3 6h18M3 12h18M3 18h18M3 6v12M9 6v12M15 6v12'] },
  'view-columns': { variant: 'outline', paths: ['M3 6h6v12H3zM9 6h6v12H9zM15 6h6v12h-6z'] },
  'funnel': { variant: 'outline', paths: ['M3 5.25h18L13.5 12v6l-3 1.5v-7.5L3 5.25z'] },
};
