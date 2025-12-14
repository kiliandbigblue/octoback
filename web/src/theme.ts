import { createTheme, MantineColorsTuple } from '@mantine/core';

// Warm coral/terracotta - appetizing and inviting
const coral: MantineColorsTuple = [
  '#ffeee9',
  '#fddcd3',
  '#f9b7a4',
  '#f68f71',
  '#f36d47',
  '#f15829',
  '#f24d1b',
  '#d83f0e',
  '#c13608',
  '#a82a00',
];

// Sage green - fresh and natural
const sage: MantineColorsTuple = [
  '#eef8f3',
  '#dfede6',
  '#bddacc',
  '#97c6af',
  '#78b596',
  '#64ab86',
  '#57a67d',
  '#46916a',
  '#3a815c',
  '#29704c',
];

// Deep charcoal for text
const charcoal: MantineColorsTuple = [
  '#f5f5f5',
  '#e7e7e7',
  '#cdcdcd',
  '#b2b2b2',
  '#9a9a9a',
  '#8b8b8b',
  '#848484',
  '#717171',
  '#656565',
  '#2d3436',
];

export const theme = createTheme({
  colors: {
    coral,
    sage,
    charcoal,
  },
  primaryColor: 'coral',
  fontFamily: '"Bricolage Grotesque", sans-serif',
  headings: {
    fontFamily: '"Bricolage Grotesque", sans-serif',
    fontWeight: '700',
  },
  radius: {
    xs: '4px',
    sm: '8px',
    md: '12px',
    lg: '20px',
    xl: '32px',
  },
  defaultRadius: 'md',
  other: {
    backgroundCream: '#fffbf7',
    backgroundDark: '#1a1a1a',
  },
});
