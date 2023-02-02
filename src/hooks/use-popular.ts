import { useQuery } from 'react-query';

export interface PopularCount {
  country: string;
  screenSize: string;
  path: string | null;
  host: string | null;
  total: number | null;
  unique: number | null;
}

const usePopular = (siteId: string) =>
  useQuery<PopularCount[]>(
    ['pageviews', 'popular'],
    async () => {
      const res = await fetch(`/api/popular?site=${siteId}`);
      return res.json();
    },
    {},
  );

export default usePopular;
