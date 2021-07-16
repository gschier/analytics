import React, { useState } from 'react';
import { HStack, VStack } from '../components/Stacks';
import Title from '../components/Title';
import ParentSize from '@visx/responsive/lib/components/ParentSize';
import TestChart from '../components/TestChart';
import Button from '../components/Button';

const Home: React.FC = () => {
    const [ count, setCount ] = useState<number>(0);

    return (
        <VStack space={3} className="m-4">
            <Title>Counter Example</Title>
            <div className="w-full h-64">
                <ParentSize>{({ width, height }) =>
                    <TestChart width={width} height={height} />
                }</ParentSize>
            </div>
            <HStack space={2}>
                <Button onClick={() => setCount(count + 1)}>
                    Count {count}
                </Button>
                <Button onClick={() => setCount(count + 1)} variant="outline" color="secondary">
                    Count {count}
                </Button>
            </HStack>
        </VStack>
    );
};

export default Home;