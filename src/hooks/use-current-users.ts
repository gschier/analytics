import { useQuery } from "react-query";

const useCurrentUsers = () =>
  useQuery<number>(
    [ "pageviews", "current" ],
    async () => {
      const res = await fetch("/api/live");
      return res.json();
    },
    {},
  );

export default useCurrentUsers;
