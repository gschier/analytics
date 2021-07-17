import React from 'react';
import { formatRelative } from 'date-fns';
import { VStack } from '../components/Stacks';
import Title from '../components/Title';
import ParentSize from '@visx/responsive/lib/components/ParentSize';
import TestChart from '../components/TestChart';
import usePageviews from '../hooks/use-pageviews';
import Card from '../components/Card';
import useRollups from '../hooks/use-rollups';
import { Table, TableRow } from '../components/Table';
import { Paragraph } from '../components/Typography';
import { capitalize } from '../util/text';

const Home: React.FC = () => {
  const { data: pageviews } = usePageviews();
  const { data: rollups } = useRollups();

  return (
    <VStack space={3} className="m-4">
      <Title>Analytics</Title>

      <div className="w-full h-64">
        {rollups && (
          <ParentSize>
            {({ width, height }) => (
              <TestChart
                width={width}
                height={height}
                data={rollups.map((r) => ({
                  date: r.start.toISOString(),
                  close: r.count,
                }))}
              />
            )}
          </ParentSize>
        )}
      </div>

      <div className="w-full h-64">
        {pageviews && (
          <Card title="Recent Views">
            <Table columns={['Date', 'Country', 'Screen', 'SID']}>
              {pageviews.slice(0, 100).map((pv) => (
                <TableRow key={pv.id}>
                  <Paragraph>
                    {capitalize(formatRelative(pv.createdAt, new Date()))}
                  </Paragraph>
                  <Paragraph>{pv.countryCode}</Paragraph>
                  <Paragraph>{pv.screenSize}</Paragraph>
                  <Paragraph>{pv.sid}</Paragraph>
                </TableRow>
              ))}
            </Table>
          </Card>
        )}
      </div>
    </VStack>
  );
};

export default Home;
