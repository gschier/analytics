import usePageviews from './use-pageviews';

const useUniqueVisitors = () => {
  const sids = new Set();
  for (const pv of usePageviews().data || []) {
    sids.add(pv.sid);
  }
  return sids.size;
};

export default useUniqueVisitors;
