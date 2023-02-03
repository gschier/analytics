import useStateLocalStorage from './use-state-localstorage';

export const useTheme = () => {
  return useStateLocalStorage<'dark' | 'light'>('theme', 'light');
};
