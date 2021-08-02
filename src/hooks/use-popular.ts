import { useQuery } from 'react-query';

export interface PopularCount {
  country: string;
  screenSize: string;
  path: string | null;
  host: string | null;
  total: number | null;
  unique: number | null;
}

const usePopular = () =>
  useQuery<PopularCount[]>(
    ['pageviews', 'popular'],
    async () => {
      const res = await fetch('/api/popular');
      return res.json();
    },
    {},
  );

export default usePopular;
