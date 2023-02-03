import React from 'react';
import { HStack, VStack } from '../components/Stacks';
import Title from '../components/Title';
import { capitalize } from '../util/text';
import Link from '../components/Link';
import useStateLocalStorage from '../hooks/use-state-localstorage';
import Button from '../components/Button';
import { Helmet } from 'react-helmet';
import { useParams } from 'react-router-dom';
import useWebsites from '../hooks/use-websites';
import { useTheme } from '../hooks/use-theme';

const Home: React.FC = () => {
  const websites = useWebsites();
  const [theme, setTheme] = useTheme();
  return (
    <VStack space={6} className="m-4">
      <HStack>
        <Title>Analytics</Title>
        <Button
          onClick={() => setTheme(theme === 'dark' ? 'light' : 'dark')}
          className="ml-auto">
          {capitalize(theme)} Mode
        </Button>
      </HStack>

      <VStack>
        {websites.data?.map((website) => (
          <Link key={website.id} to={`/analytics/${website.id}`}>
            {website.domain}
          </Link>
        ))}
      </VStack>
    </VStack>
  );
};

export default Home;
