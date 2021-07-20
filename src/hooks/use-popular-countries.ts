import { useQuery } from 'react-query';

export interface CountryCount {
  country: string;
  screenSize: string;
  path: string | null;
  host: string | null;
  total: number | null;
  unique: number | null;
}

const usePopularCountries = () =>
  useQuery<CountryCount[]>(
    ['pageviews', 'popular', 'countries'],
    async () => {
      const res = await fetch('/api/rollups/pageviews/popular');
      return res.json();
    },
    {},
  );

export default usePopularCountries;
