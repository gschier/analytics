import React, { HTMLAttributes } from 'react';
import { HStack, VStack } from './Stacks';
import classnames from 'classnames';

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
    <VStack
      {...props}
      className={classnames(
        'bg-gray-50 dark:bg-primary-100 ring-1 ring-gray-100 rounded divide-y divide-gray-100',
        className,
      )}>
      {title && <HStack className="px-3 py-2">{title}</HStack>}
      <VStack
        space={3}
        className="p-3 bg-gray-0 dark:bg-primary-50 rounded-b-lg">
        {children}
      </VStack>
    </VStack>
  );
};

export default Card;
