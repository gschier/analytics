import useStateLocalStorage from './use-state-localstorage';
import { useEffect, useState } from 'react';
import { act } from 'react-dom/test-utils';

export const useTheme = (): [
  'dark' | 'light',
  (theme: 'dark' | 'light') => void,
] => {
  const [storedTheme, setStoredTheme] = useStateLocalStorage<
    'dark' | 'light' | null
  >('theme', null);
  const [activeTheme, setActiveTheme] = useState<'dark' | 'light'>(
    storedTheme ?? (prefersDark() ? 'dark' : 'light'),
  );
  setThemeClass(activeTheme);

  const setTheme = (theme: 'dark' | 'light') => {
    setStoredTheme(theme);
    setActiveTheme(theme);
    setThemeClass(theme);
  };

  useEffect(() => {
    window
      .matchMedia('(prefers-color-scheme: dark)')
      .addEventListener('change', () => {
        if (storedTheme === null) {
          setActiveTheme(prefersDark() ? 'dark' : 'light');
        }
      });
  }, []);

  return [activeTheme, setTheme];
};

const prefersDark = () =>
  window.matchMedia('(prefers-color-scheme: dark)').matches;

const setThemeClass = (theme: 'light' | 'dark') => {
  if (theme === 'dark') {
    document.documentElement.classList.add('dark');
  } else {
    document.documentElement.classList.remove('dark');
  }
};
