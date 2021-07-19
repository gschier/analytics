import { useQuery } from 'react-query';

export interface PageCount {
  path: string;
  host: string;
  total: number;
  unique: number;
}

const usePopularPages = () =>
  useQuery<PageCount[]>(
    ['pageviews', 'popular', 'pages'],
    async () => {
      const res = await fetch('/api/rollups/pageviews/paths');
      return res.json();
    },
    {},
  );

export default usePopularPages;
