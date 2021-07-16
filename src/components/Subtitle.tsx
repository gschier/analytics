import React, { HTMLAttributes } from 'react';

const Subtitle: React.FC<HTMLAttributes<HTMLHeadingElement>> = ({
    className,
    ...props
}) => {
    return <h2 {...props} className={`${className ?? ''} text-2xl font-bold`} />;
};

export default Subtitle;
