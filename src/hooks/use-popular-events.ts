import { useQuery } from 'react-query';

export interface PopularCount {
  country: string;
  screenSize: string;
  name: string;
  platform: string | null;
  version: string | null;
  total: number | null;
  unique: number | null;
}

const usePopularEvents = (siteId: string) =>
  useQuery<PopularCount[]>(
    ['pageviews', 'popular_events'],
    async () => {
      const res = await fetch(`/api/popular_events?site=${siteId}`);
      return res.json();
    },
    {},
  );

export default usePopularEvents;
