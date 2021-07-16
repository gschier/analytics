import React, { HTMLAttributes } from 'react';
import classnames from 'classnames';

export type TextSize = 'sm' | 'md' | 'lg';

export interface TextProps {
    size?: TextSize;
}

const sizeClassMap: Record<TextSize, string> = {
    sm: 'text-sm',
    md: 'text-base',
    lg: 'text-lg',
};

const Text: React.FC<TextProps & HTMLAttributes<HTMLHeadingElement>> = ({
    className,
    size,
    ...props
}) => {
    return (
        <p
            {...props}
            className={classnames(
                'font-normal',
                'text-gray-600',
                sizeClassMap[size ?? 'md'],
                className,
            )}
        />
    );
};

export default Text;
