import { useQuery } from 'react-query';

export interface Pageview {
  id: string;
  websiteId: string;
  sid: string;
  createdAt: Date;
  host: string;
  path: string;
  screenSize: string;
  countryCode: string;
  userAgent: string;
}

const usePageviews = () =>
  useQuery<Pageview[]>(
    'pageviews',
    async () => {
      const res = await fetch('/api/pageviews');
      return (await res.json()).map((pv: any) => {
        pv.createdAt = new Date(pv.createdAt);
        return pv;
      });
    },
    {},
  );

export default usePageviews;
