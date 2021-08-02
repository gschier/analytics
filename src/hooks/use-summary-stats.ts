import usePopular, { PopularCount } from "./use-popular";

const useSummaryStats = (): PopularCount | null => {
  const counts = usePopular();
  const all = counts.data?.find(c => c.path === null && c.country === null && c.host === null && c.screenSize === null);
  return all ? all : null;
};

export default useSummaryStats;
