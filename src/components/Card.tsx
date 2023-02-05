import React, { HTMLAttributes } from 'react';
import { VStack } from './Stacks';
import classnames from 'classnames';
import { Paragraph } from './Typography';

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
        'ring-1 ring-gray-100 rounded divide-y divide-gray-100 overflow-hidden',
        className,
      )}>
      {title && (
        <Paragraph
          size="sm"
          className="px-3 py-2 bg-gray-50 text-gray-600 uppercase">
          {title}
        </Paragraph>
      )}
      <VStack space={3} className="p-3 rounded-b-lg">
        {children}
      </VStack>
    </VStack>
  );
};

export default Card;
