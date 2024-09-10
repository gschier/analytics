import { useQuery } from 'react-query';

export interface PopularCount {
  name: string;
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
