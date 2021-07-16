import React, { HTMLAttributes } from 'react';
import { HStack, VStack } from './Stacks';


export interface CardProps {
    title?: string;
}

const Card: React.FC<CardProps & HTMLAttributes<HTMLDivElement>> = ({
    className,
    title,
    children,
    ...props
}) => {
    return (
        <VStack {...props} className={`${className ?? ''} bg-gray-50 ring-1 ring-gray-100 rounded-lg divide-y divide-gray-100`}>
            {title && (
                <HStack className="px-3 py-2">
                    {title}
                </HStack>
            )}
            <VStack space={3} className="p-3 bg-gray-0 rounded-b-lg">
                {children}
            </VStack>
        </VStack>
    );
};

export default Card;
