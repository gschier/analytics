import { useQuery } from 'react-query';

export interface CountryCount {
  country: string;
  total: number;
  unique: number;
}

const usePopularCountries = () =>
  useQuery<CountryCount[]>(
    ['pageviews', 'popular', 'countries'],
    async () => {
      const res = await fetch('/api/rollups/pageviews/countries');
      return res.json();
    },
    {},
  );

export default usePopularCountries;
