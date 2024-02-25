import React from 'react';
import { HStack, VStack } from '../components/Stacks';
import ParentSize from '@visx/responsive/lib/components/ParentSize';
import TestChart from '../components/TestChart';
import Card from '../components/Card';
import useRollups from '../hooks/use-rollups';
import { Table, TableRow } from '../components/Table';
import { HugeText, Paragraph } from '../components/Typography';
import Link from '../components/Link';
import useCurrentUsers from '../hooks/use-current-users';
import useSummaryStats from '../hooks/use-summary-stats';
import usePopular from '../hooks/use-popular';
import { dateBetween } from '../util/date';
import { useParams } from 'react-router-dom';
import Navbar from '../components/Navbar';
import useWebsite from '../hooks/use-website';
import usePopularEvents from '../hooks/use-popular-events';

const Site: React.FC = () => {
  const { id: websiteId } = useParams<{ id: string }>();
  const website = useWebsite(websiteId);
  const { data: rollups } = useRollups(websiteId);
  const { data: popularPaths } = usePopular(websiteId);
  const { data: popularEvents } = usePopularEvents(websiteId);
  const currentUsers = useCurrentUsers(websiteId);
  const summaryStats = useSummaryStats(websiteId);
  const countries = popularPaths?.filter((pp) => pp.country).slice(0, 6) ?? [];
  const screenSizes =
    popularPaths?.filter((pp) => pp.screenSize).slice(0, 6) ?? [];
  const paths = popularPaths?.filter((pp) => pp.path).slice(0, 10) ?? [];
  const events = popularEvents?.filter((pe) => pe.name) ?? [];

  return (
    <>
      <Navbar website={website} />
      <VStack space={6} className="mx-4 my-6">
        {currentUsers.data ||
          summaryStats?.unique ||
          (summaryStats?.total && (
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
          ))}

        {rollups?.some((r) => r.total > 0) && (
          <div className="w-full h-64">
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
          </div>
        )}

        <HStack collapse space={3} align="start">
          {countries.length > 0 && (
            <Table columns={['Country', 'Unique', 'Total']}>
              {countries.map((pp) => (
                <TableRow key={pp.country}>
                  <Paragraph>{pp.country}</Paragraph>
                  <Paragraph>{pp.unique}</Paragraph>
                  <Paragraph>{pp.total}</Paragraph>
                </TableRow>
              ))}
            </Table>
          )}
          {screenSizes.length > 0 && (
            <Table columns={['Screen Size', 'Unique', 'Total']}>
              {screenSizes.map((pp) => (
                <TableRow key={pp.screenSize}>
                  <Paragraph>{pp.screenSize}</Paragraph>
                  <Paragraph>{pp.unique}</Paragraph>
                  <Paragraph>{pp.total}</Paragraph>
                </TableRow>
              ))}
            </Table>
          )}
        </HStack>

        {paths.length > 0 && (
          <Table columns={['Path', 'Unique', 'Total']}>
            {paths.map((pp) => (
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
        {events.length && (
          <Table columns={['Event', 'Unique', 'Total']}>
            {events.map((pe, i) => (
              <TableRow key={i}>
                <Paragraph>{pe.name}</Paragraph>
                <Paragraph>{pe.unique}</Paragraph>
                <Paragraph>{pe.total}</Paragraph>
              </TableRow>
            ))}
          </Table>
        )}
      </VStack>
    </>
  );
};

export default Site;
