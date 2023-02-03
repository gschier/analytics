import React from 'react';
import { VStack } from '../components/Stacks';
import Link from '../components/Link';
import useWebsites from '../hooks/use-websites';
import Navbar from '../components/Navbar';

const Home: React.FC = () => {
  const websites = useWebsites();
  return (
    <>
      <Navbar />
      <VStack className="mx-4">
        {websites.data?.map((website) => (
          <div key={website.id}>
            <Link to={`/analytics/${website.id}`}>{website.domain}</Link>
          </div>
        ))}
      </VStack>
    </>
  );
};

export default Home;
