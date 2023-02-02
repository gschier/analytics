import usePopular, { PopularCount } from "./use-popular";

const useSummaryStats = (siteId: string): PopularCount | null => {
  const counts = usePopular(siteId);
  const all = counts.data?.find(c => c.path === null && c.country === null && c.host === null && c.screenSize === null);
  return all ? all : null;
};

export default useSummaryStats;
