import React, { HTMLAttributes } from 'react';
import { HStack, VStack } from '../components/Stacks';
import Title from '../components/Title';
import Text from '../components/Text';
import Button, { ButtonColor, ButtonSize, ButtonVariant } from '../components/Button';
import Subtitle from '../components/Subtitle';
import { Helmet } from 'react-helmet';
import { capitalize } from '../util/text';
import Input from '../components/Input';
import { Table, TableRow } from '../components/Table';
import ParentSize from '@visx/responsive/lib/components/ParentSize';
import TestChart from '../components/TestChart';
import Card from '../components/Card';
import useStateLocalStorage from '../hooks/use-state-localstorage';

const Design: React.FC = () => {
    const [ theme, setTheme ] = useStateLocalStorage<'dark' | 'light'>('theme', 'light');
    const colors: ButtonColor[] = [ 'primary', 'secondary', 'danger', 'gray' ];
    const buttonVariants: ButtonVariant[] = [ 'solid', 'outline' ];
    const buttonSizes: ButtonSize[] = [ 'lg', 'md', 'sm' ];

    return (
        <VStack space={3} className="m-4">
            <Helmet>
                <html className={theme} />
            </Helmet>
            <HStack>
                <Title>Design System</Title>
                <Button onClick={() => setTheme(theme === 'dark' ? 'light' : 'dark')} className="ml-auto">
                    {capitalize(theme)}
                </Button>
            </HStack>
            <Text>This is the design system for the analytics tool.</Text>

            <Subtitle>Cards</Subtitle>
            <HStack wrap space={3} align="start">
                <Card title="Example Card" className="max-w-sm">
                    <Input placeholder="Email" />
                    <Input type="password" placeholder="Password" />
                    <Button className="w-full">Submit</Button>
                </Card>
                <Card title="Example Card" className="max-w-xs">
                    <div className="w-full h-32">
                        <ParentSize>{({ width, height }) =>
                            <TestChart width={width} height={height} />
                        }</ParentSize>
                    </div>
                </Card>
                <Card title="Example Card" className="max-w-xs">
                    <Text>
                        Hello, this is some text within an example card. It's long enough to showcase how text may wrap
                        inside a card but the text itself is not interesting at all.
                    </Text>
                </Card>
            </HStack>

            <Subtitle>Tables</Subtitle>
            <VStack space={2}>
                <Table columns={[ 'Name', 'Description' ]}>
                    <TableRow>
                        <Text>Jane Cooper</Text>
                        <Text>Regional Paradigm Technician</Text>
                    </TableRow>
                    <TableRow>
                        <Text>Cody Fisher</Text>
                        <Text>Product Directives Officer</Text>
                    </TableRow>
                    <TableRow>
                        <Text>Esther Howard</Text>
                        <Text>Forward Response Developer</Text>
                    </TableRow>
                </Table>
            </VStack>

            <Subtitle>Charts</Subtitle>
            <HStack space={2}>
                <div className="w-1/3 h-32">
                    <ParentSize>{({ width, height }) =>
                        <TestChart width={width} height={height} />
                    }</ParentSize>
                </div>
                <div className="w-2/3 h-32">
                    <ParentSize>{({ width, height }) =>
                        <TestChart width={width} height={height} />
                    }</ParentSize>
                </div>
            </HStack>

            <Subtitle>Buttons</Subtitle>
            <VStack space={2}>
                {colors.map(c => (
                    <HStack space={2} key={c}>
                        {buttonSizes.map(s =>
                            buttonVariants.map(v => (
                                <Button key={v + s + c} variant={v} size={s} color={c}>{capitalize(v)}</Button>
                            )),
                        )}
                    </HStack>
                ))}
            </VStack>

            <Subtitle>Inputs</Subtitle>
            <VStack space={3} className="max-w-sm">
                <Input label="Small Input" placeholder="Some value" size="sm" />
                <Input label="Medium Input" placeholder="Some value" size="md" />
                <Input label="Large Input" placeholder="Some value" size="lg" />
                <Input textarea label="Textarea" placeholder="Some value" size="lg" />
                <Input label="Errored Input" placeholder="Some value" error="That's not right" defaultValue="My bad" />
            </VStack>

            <Subtitle>Colors</Subtitle>
            <VStack space={1}>
                <HStack space={1}>
                    <ColorBlock className="bg-primary-0" />
                    <ColorBlock className="bg-primary-50" />
                    <ColorBlock className="bg-primary-100" />
                    <ColorBlock className="bg-primary-200" />
                    <ColorBlock className="bg-primary-300" />
                    <ColorBlock className="bg-primary-400" />
                    <ColorBlock className="bg-primary-500" />
                    <ColorBlock className="bg-primary-600" />
                    <ColorBlock className="bg-primary-700" />
                    <ColorBlock className="bg-primary-800" />
                    <ColorBlock className="bg-primary-900" />
                </HStack>
                <HStack space={1}>
                    <ColorBlock className="bg-secondary-0" />
                    <ColorBlock className="bg-secondary-50" />
                    <ColorBlock className="bg-secondary-100" />
                    <ColorBlock className="bg-secondary-200" />
                    <ColorBlock className="bg-secondary-300" />
                    <ColorBlock className="bg-secondary-400" />
                    <ColorBlock className="bg-secondary-500" />
                    <ColorBlock className="bg-secondary-600" />
                    <ColorBlock className="bg-secondary-700" />
                    <ColorBlock className="bg-secondary-800" />
                    <ColorBlock className="bg-secondary-900" />
                </HStack>
                <HStack space={1}>
                    <ColorBlock className="bg-danger-0" />
                    <ColorBlock className="bg-danger-50" />
                    <ColorBlock className="bg-danger-100" />
                    <ColorBlock className="bg-danger-200" />
                    <ColorBlock className="bg-danger-300" />
                    <ColorBlock className="bg-danger-400" />
                    <ColorBlock className="bg-danger-500" />
                    <ColorBlock className="bg-danger-600" />
                    <ColorBlock className="bg-danger-700" />
                    <ColorBlock className="bg-danger-800" />
                    <ColorBlock className="bg-danger-900" />
                </HStack>
                <HStack space={1}>
                    <ColorBlock className="bg-gray-0" />
                    <ColorBlock className="bg-gray-50" />
                    <ColorBlock className="bg-gray-100" />
                    <ColorBlock className="bg-gray-200" />
                    <ColorBlock className="bg-gray-300" />
                    <ColorBlock className="bg-gray-400" />
                    <ColorBlock className="bg-gray-500" />
                    <ColorBlock className="bg-gray-600" />
                    <ColorBlock className="bg-gray-700" />
                    <ColorBlock className="bg-gray-800" />
                    <ColorBlock className="bg-gray-900" />
                </HStack>
            </VStack>
        </VStack>
    );
};

const ColorBlock: React.FC<HTMLAttributes<HTMLDivElement>> = ({ className, ...props }) => (
    <div {...props} className={`${className ?? ''} w-8 h-6 rounded`} />
);

export default Design;