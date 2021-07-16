import React, { HTMLAttributes } from 'react';

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
    const sizeClass = sizeClassMap[size ?? 'md'];
    return <p {...props} className={`${className ?? ''} ${sizeClass} font-normal text-gray-600`} />;
};

export default Text;
