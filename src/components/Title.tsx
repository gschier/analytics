import React, { HTMLAttributes } from 'react';
import classnames from 'classnames';

const Title: React.FC<HTMLAttributes<HTMLHeadingElement>> = ({
    className,
    ...props
}) => {
    return (
        <h1
            {...props}
            className={classnames(
                'text-4xl',
                'font-bold',
                'text-gray-800',
                className,
            )}
        />
    );
};

export default Title;
