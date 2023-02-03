import useWebsites, { Website } from './use-websites';

const useWebsite = (websiteId: string) => {
  const websites = useWebsites();
  return websites.data?.find((website: Website) => website.id === websiteId);
};

export default useWebsite;
