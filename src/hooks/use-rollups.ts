import { useQuery } from 'react-query';

export interface Rollup {
  start: Date;
  end: Date;
  total: number;
  unique: number;
}

const useRollups = () =>
  useQuery<Rollup[]>(
    'rollups',
    async () => {
      const res = await fetch('/api/rollups/pageviews');
      return (await res.json()).map((r: any) => {
        r.start = new Date(r.start);
        r.end = new Date(r.end);
        return r;
      });
    },
    {},
  );

export default useRollups;
