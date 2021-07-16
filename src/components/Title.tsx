import React, { HTMLAttributes } from 'react';

const Title: React.FC<HTMLAttributes<HTMLHeadingElement>> = ({
    className,
    ...props
}) => {
    return <h1 {...props} className={`${className ?? ''} text-4xl font-bold text-gray-800`} />;
};

export default Title;
