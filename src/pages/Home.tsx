import React from 'react';
import { formatRelative } from 'date-fns';
import { HStack, VStack } from '../components/Stacks';
import Title from '../components/Title';
import ParentSize from '@visx/responsive/lib/components/ParentSize';
import TestChart from '../components/TestChart';
import usePageviews from '../hooks/use-pageviews';
import Card from '../components/Card';
import useRollups from '../hooks/use-rollups';
import { Table, TableRow } from '../components/Table';
import { HugeText, Paragraph } from '../components/Typography';
import { capitalize } from '../util/text';
import Link from '../components/Link';
import useCurrentUsers from '../hooks/use-current-users';
import useUniqueVisitors from '../hooks/use-unique-visitors';
import useStateLocalStorage from '../hooks/use-state-localstorage';
import Button from '../components/Button';
import { Helmet } from 'react-helmet';
import usePopularPages from '../hooks/use-popular-pages';
import usePopularCountries from '../hooks/use-popular-countries';

const Home: React.FC = () => {
  const [theme, setTheme] = useStateLocalStorage<'dark' | 'light'>(
    'theme',
    'light',
  );
  const { data: pageviews } = usePageviews();
  const { data: rollups } = useRollups();
  const { data: popularPages } = usePopularPages();
  const { data: popularCountries } = usePopularCountries();
  const currentUsers = useCurrentUsers();
  const uniqueVisitors = useUniqueVisitors();

  return (
    <VStack space={6} className="m-4">
      <Helmet>
        <html className={theme} />
      </Helmet>
      <HStack>
        <Title>Analytics</Title>
        <Button
          onClick={() => setTheme(theme === 'dark' ? 'light' : 'dark')}
          className="ml-auto">
          {capitalize(theme)} Mode
        </Button>
      </HStack>

      <HStack space={3}>
        <Card title="Live" className="w-full">
          <HugeText>{currentUsers}</HugeText>
        </Card>
        <Card title="Visitors" className="w-full">
          <HugeText>{uniqueVisitors}</HugeText>
        </Card>
        <Card title="Views" className="w-full">
          <HugeText>{pageviews ? pageviews.length : <>&nbsp;</>}</HugeText>
        </Card>
      </HStack>

      <div className="w-full h-64">
        {rollups && (
          <ParentSize>
            {({ width, height }) => (
              <TestChart
                width={width}
                height={height}
                data={rollups.map((r) => ({
                  date: r.start.toISOString(),
                  close: r.unique,
                }))}
                data2={rollups.map((r) => ({
                  date: r.start.toISOString(),
                  close: r.total,
                }))}
              />
            )}
          </ParentSize>
        )}
      </div>

      <HStack space={3} align="start">
        {popularPages && (
          <Table columns={['Path', 'Unique', 'Total']} className="w-2/3">
            {popularPages.map((pp) => (
              <TableRow key={pp.path}>
                <Paragraph>{pp.path}</Paragraph>
                <Paragraph>{pp.unique}</Paragraph>
                <Paragraph>{pp.total}</Paragraph>
              </TableRow>
            ))}
          </Table>
        )}

        {popularCountries && (
          <Table columns={['Country', 'Unique', 'Total']} className="w-1/3">
            {popularCountries.map((pp) => (
              <TableRow key={pp.country}>
                <Paragraph>{pp.country}</Paragraph>
                <Paragraph>{pp.unique}</Paragraph>
                <Paragraph>{pp.total}</Paragraph>
              </TableRow>
            ))}
          </Table>
        )}
      </HStack>

      {pageviews && (
        <Table columns={['Date', 'Path', 'Country', 'Screen', 'SID']}>
          {pageviews.map((pv) => (
            <TableRow key={pv.id}>
              <Paragraph>
                {capitalize(formatRelative(pv.createdAt, new Date()))}
              </Paragraph>
              <Paragraph>
                <Link external to={`${pv.host}${pv.path}`}>
                  {`${pv.path}`}
                </Link>
              </Paragraph>
              <Paragraph>{pv.countryCode}</Paragraph>
              <Paragraph>{pv.screenSize}</Paragraph>
              <Paragraph>{pv.sid.slice(0, 5)}</Paragraph>
            </TableRow>
          ))}
        </Table>
      )}
    </VStack>
  );
};

export default Home;
