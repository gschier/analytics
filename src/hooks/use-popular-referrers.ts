import { useQuery } from 'react-query';

export interface PopularReferrerCount {
  referrer: string;
  total: number | null;
  unique: number | null;
}

const usePopularReferrers = (siteId: string) =>
  useQuery<PopularReferrerCount[]>(
    ['pageviews', 'popular_referrers'],
    async () => {
      const res = await fetch(`/api/popular_referrers?site=${siteId}`);
      return res.json();
    },
    {},
  );

export default usePopularReferrers;
