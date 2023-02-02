import { useQuery } from "react-query";

const useCurrentUsers = (siteId: string) =>
  useQuery<number>(
    [ "pageviews", "current" ],
    async () => {
      const res = await fetch(`/api/live?site=${siteId}`);
      return res.json();
    },
    {},
  );

export default useCurrentUsers;
