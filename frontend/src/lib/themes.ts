// 4-color swatch data: [bg, primary, secondary, accent]
// Used by ThemePicker to render the color strip preview.

export interface ThemeDef {
  name: string;
  dark: boolean;
  colors: [string, string, string, string];
}

export const LIGHT_THEMES: ThemeDef[] = [
  { name: 'light',      dark: false, colors: ['#f2f2f2', '#570df8', '#f000b8', '#37cdbe'] },
  { name: 'cupcake',    dark: false, colors: ['#faf7f5', '#65c3c8', '#ef9fbc', '#eeaf3a'] },
  { name: 'bumblebee',  dark: false, colors: ['#ffffff', '#e0a82e', '#181830', '#36d399'] },
  { name: 'emerald',    dark: false, colors: ['#f9fafb', '#66cc8a', '#377cfb', '#36d399'] },
  { name: 'corporate',  dark: false, colors: ['#f9f9f9', '#4b6bfb', '#7b92b2', '#36d399'] },
  { name: 'retro',      dark: false, colors: ['#e4d8b4', '#ef9995', '#a4cbb4', '#45c97a'] },
  { name: 'cyberpunk',  dark: false, colors: ['#ffee00', '#ff7598', '#75d1f0', '#adff2f'] },
  { name: 'valentine',  dark: false, colors: ['#fae8f0', '#e96d9b', '#a991f7', '#36d399'] },
  { name: 'garden',     dark: false, colors: ['#e9e7e7', '#5c7f67', '#cc4c4c', '#5c7f67'] },
  { name: 'aqua',       dark: false, colors: ['#345da7', '#09ecf3', '#966fb3', '#36d399'] },
  { name: 'lofi',       dark: false, colors: ['#ffffff', '#0d0d0d', '#1a1a1a', '#36d399'] },
  { name: 'pastel',     dark: false, colors: ['#f8f0fb', '#d1a3e8', '#f4a7b9', '#a8d8a8'] },
  { name: 'fantasy',    dark: false, colors: ['#ffffff', '#6e0b75', '#007ebd', '#36d399'] },
  { name: 'winter',     dark: false, colors: ['#f0f4fc', '#047aff', '#463aa1', '#36d399'] },
];

export const DARK_THEMES: ThemeDef[] = [
  { name: 'dark',       dark: true,  colors: ['#1d232a', '#661ae4', '#d926a9', '#1fb2a6'] },
  { name: 'synthwave',  dark: true,  colors: ['#1a103c', '#e779c1', '#58c7f3', '#36d399'] },
  { name: 'halloween',  dark: true,  colors: ['#1d0500', '#f28c18', '#6d3a63', '#36d399'] },
  { name: 'forest',     dark: true,  colors: ['#0e1c0e', '#1eb854', '#1fd65f', '#1eb854'] },
  { name: 'black',      dark: true,  colors: ['#000000', '#ffffff', '#999999', '#36d399'] },
  { name: 'luxury',     dark: true,  colors: ['#09090b', '#c4a46b', '#9a9a9a', '#36d399'] },
  { name: 'dracula',    dark: true,  colors: ['#282a36', '#ff79c6', '#bd93f9', '#50fa7b'] },
  { name: 'night',      dark: true,  colors: ['#0f1729', '#38bdf8', '#818cf8', '#4ade80'] },
  { name: 'coffee',     dark: true,  colors: ['#20161f', '#db924b', '#263e3f', '#36d399'] },
  { name: 'dim',        dark: true,  colors: ['#2a2e37', '#9FE88D', '#FF7D5C', '#45AEEE'] },
  { name: 'sunset',     dark: true,  colors: ['#f9d8b6', '#316794', '#8fbc8f', '#ef9c25'] },
];

export const ALL_THEMES: ThemeDef[] = [...LIGHT_THEMES, ...DARK_THEMES];

export function getThemeDef(name: string): ThemeDef | undefined {
  return ALL_THEMES.find(t => t.name === name);
}
