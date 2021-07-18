import usePageviews from './use-pageviews';

const useCurrentUsers = () => {
  const sids = new Set();
  for (const pv of usePageviews().data || []) {
    const millis = Date.now() - pv.createdAt.getTime();
    if (millis > 1000 * 60 * 5) {
      continue;
    }
    sids.add(pv.sid);
  }
  return sids.size;
};

export default useCurrentUsers;
