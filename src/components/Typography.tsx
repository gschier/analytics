import React, { HTMLAttributes } from 'react';
import classnames from 'classnames';

export type TextSize = 'sm' | 'md' | 'lg' | 'xl';

export interface TextProps {
  size?: TextSize;
}

const sizeClassMap: Record<TextSize, string> = {
  sm: classnames('text-sm font-normal'),
  md: classnames('text-base font-normal'),
  lg: classnames('text-lg font-normal'),
  xl: classnames('text-2xl font-normal'),
};

export const Paragraph: React.FC<
  TextProps & HTMLAttributes<HTMLHeadingElement>
> = ({ className, size, ...props }) => {
  return (
    <p
      {...props}
      className={classnames(
        'font-normal text-gray-600',
        sizeClassMap[size ?? 'md'],
        className,
      )}
    />
  );
};

export const InlineText: React.FC<TextProps & HTMLAttributes<HTMLSpanElement>> =
  ({ className, size, ...props }) => {
    return (
      <span
        {...props}
        className={classnames(sizeClassMap[size ?? 'md'], className)}
      />
    );
  };

export const HugeText: React.FC<HTMLAttributes<HTMLSpanElement>> = (props) => (
  <InlineText {...props} size="xl" />
);
