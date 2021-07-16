import React, { HTMLAttributes } from 'react';

type StackSpace = 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 12 | 14 | 16;

const vSpaces: Record<StackSpace, string> = {
    0: 'space-y-0',
    1: 'space-y-1',
    2: 'space-y-2',
    3: 'space-y-3',
    4: 'space-y-4',
    5: 'space-y-5',
    6: 'space-y-6',
    7: 'space-y-7',
    8: 'space-y-8',
    9: 'space-y-8',
    10: 'space-y-8',
    12: 'space-y-12',
    14: 'space-y-14',
    16: 'space-y-16',
};

const hSpaces: Record<StackSpace, string> = {
    0: 'space-x-0',
    1: 'space-x-1',
    2: 'space-x-2',
    3: 'space-x-3',
    4: 'space-x-4',
    5: 'space-x-5',
    6: 'space-x-6',
    7: 'space-x-7',
    8: 'space-x-8',
    9: 'space-x-8',
    10: 'space-x-8',
    12: 'space-x-12',
    14: 'space-x-14',
    16: 'space-x-16',
};

export interface StackProps extends HTMLAttributes<HTMLDivElement> {
    justify?: JustifyValues;
    align?: AlignValues;
    space?: StackSpace;
    className?: string;
}

type JustifyValues = 'center' | 'between' | 'start' | 'end';
type AlignValues = 'center' | 'start' | 'end' | 'baseline';

const justifyValues: Record<JustifyValues, string> = {
    center: 'justify-center',
    between: 'justify-between',
    start: 'justify-start',
    end: 'justify-end',
};

const alignValues: Record<AlignValues, string> = {
    center: 'items-center',
    baseline: 'items-baseline',
    start: 'items-start',
    end: 'items-end',
};

export interface StackProps {
    space?: StackSpace;
    justify?: JustifyValues;
    align?: AlignValues;
}

export const VStack: React.FC<StackProps & HTMLAttributes<HTMLDivElement>> = ({
    className,
    space,
    justify,
    align,
    ...props
}) => {
    const spaceClass = space !== undefined ? vSpaces[space] : '';
    const justifyClass = justify ? justifyValues[justify] : '';
    const alignClass = align ? alignValues[align] : '';
    return (
        <div
            {...props}
            className={`${className ?? ''} flex flex-col ${spaceClass} ${justifyClass} ${alignClass}`}
        />
    );
};

export const HStack: React.FC<StackProps & HTMLAttributes<HTMLDivElement>> = ({
    className,
    space,
    justify,
    align,
    ...props
}) => {
    const spaceClass = space === undefined ? '' : hSpaces[space ?? 0];
    const justifyClass = justifyValues[justify ?? 'start'];
    const alignClass = alignValues[align ?? 'end'];
    return (
        <div
            {...props}
            className={`${className ?? ''} w-full flex flex-row ${spaceClass} ${justifyClass} ${alignClass}`}
        />
    );
};

