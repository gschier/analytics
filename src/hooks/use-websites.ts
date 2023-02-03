import { useQuery } from 'react-query';

export interface Website {
  id: string;
  domain: string;
}

const useWebsites = () =>
  useQuery<Website[]>(
    ['websites'],
    async () => {
      const res = await fetch('/api/websites');
      return res.json();
    },
    {},
  );

export default useWebsites;
