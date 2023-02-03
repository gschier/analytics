import React from 'react';
import { HStack, VStack } from '../components/Stacks';
import Title from '../components/Title';
import ParentSize from '@visx/responsive/lib/components/ParentSize';
import TestChart from '../components/TestChart';
import Card from '../components/Card';
import useRollups from '../hooks/use-rollups';
import { Table, TableRow } from '../components/Table';
import { HugeText, Paragraph } from '../components/Typography';
import { capitalize } from '../util/text';
import Link from '../components/Link';
import useCurrentUsers from '../hooks/use-current-users';
import useSummaryStats from '../hooks/use-summary-stats';
import Button from '../components/Button';
import { Helmet } from 'react-helmet';
import usePopular from '../hooks/use-popular';
import { dateBetween } from '../util/date';
import { useParams } from 'react-router-dom';
import useWebsite from '../hooks/use-website';
import { useTheme } from '../hooks/use-theme';

const Site: React.FC = () => {
  const { id: websiteId } = useParams<{ id: string }>();
  const website = useWebsite(websiteId);
  const [theme, setTheme] = useTheme();
  const { data: rollups } = useRollups(websiteId);
  const { data: popular } = usePopular(websiteId);
  const currentUsers = useCurrentUsers(websiteId);
  const summaryStats = useSummaryStats(websiteId);

  return (
    <VStack space={6} className="m-4">
      <Helmet>
        <html className={theme} />
      </Helmet>
      <HStack>
        <Title>{website?.domain}</Title>
        <Button
          onClick={() => setTheme(theme === 'dark' ? 'light' : 'dark')}
          className="ml-auto">
          {capitalize(theme)} Mode
        </Button>
      </HStack>

      <HStack space={3}>
        <Card title="Live" className="w-full">
          <HugeText>{currentUsers.data}</HugeText>
        </Card>
        <Card title="Visitors" className="w-full">
          <HugeText>{summaryStats?.unique}</HugeText>
        </Card>
        <Card title="Views" className="w-full">
          <HugeText>{summaryStats?.total}</HugeText>
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
                  date: dateBetween(r.start, r.end).toISOString(),
                  close: r.unique,
                }))}
                data2={rollups.map((r) => ({
                  date: dateBetween(r.start, r.end).toISOString(),
                  close: r.total,
                }))}
              />
            )}
          </ParentSize>
        )}
      </div>

      {popular && (
        <HStack collapse space={3} align="start">
          <Table columns={['Country', 'Unique', 'Total']}>
            {popular
              .filter((pp) => pp.country)
              .slice(0, 6)
              .map((pp) => (
                <TableRow key={pp.country}>
                  <Paragraph>{pp.country}</Paragraph>
                  <Paragraph>{pp.unique}</Paragraph>
                  <Paragraph>{pp.total}</Paragraph>
                </TableRow>
              ))}
          </Table>
          <Table columns={['Screen Size', 'Unique', 'Total']}>
            {popular
              .filter((pp) => pp.screenSize)
              .slice(0, 6)
              .map((pp) => (
                <TableRow key={pp.screenSize}>
                  <Paragraph>{pp.screenSize}</Paragraph>
                  <Paragraph>{pp.unique}</Paragraph>
                  <Paragraph>{pp.total}</Paragraph>
                </TableRow>
              ))}
          </Table>
        </HStack>
      )}

      {popular && (
        <Table columns={['Path', 'Unique', 'Total']}>
          {popular
            .filter((pp) => pp.path)
            .slice(0, 10)
            .map((pp) => (
              <TableRow key={pp.path}>
                <Link external to={`${pp.host}${pp.path}`}>
                  {`${pp.path}`}
                </Link>
                <Paragraph>{pp.unique}</Paragraph>
                <Paragraph>{pp.total}</Paragraph>
              </TableRow>
            ))}
        </Table>
      )}
    </VStack>
  );
};

export default Site;
