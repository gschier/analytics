import React, { useEffect, useState } from 'react';
import { VStack } from '../components/Stacks';
import Title from '../components/Title';
import ParentSize from '@visx/responsive/lib/components/ParentSize';
import TestChart from '../components/TestChart';

const Home: React.FC = () => {
    const [ pageviews, setPageviews ] = useState<any[]>([]);
    useEffect(() => {
        fetch('/api/pageviews').then(async res => {
            const data = await res.json();
            setPageviews(data);
        }, err => {
            console.log('ERROR', err.message);
        });
    }, []);

    return (
        <VStack space={3} className="m-4">
            <Title>Analytics</Title>
            <div className="w-full h-64">
                <ParentSize>{({ width, height }) =>
                    <TestChart width={width} height={height} data={pageviews.map(pv => ({
                        date: pv.Start,
                        close: pv.Count,
                    }))} />
                }</ParentSize>
            </div>
        </VStack>
    );
};

export default Home;