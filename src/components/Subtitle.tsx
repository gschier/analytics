import classnames from 'classnames';
import React, { HTMLAttributes } from 'react';

const Subtitle: React.FC<HTMLAttributes<HTMLHeadingElement>> = ({
    className,
    ...props
}) => {
    return (
        <h2
            {...props}
            className={classnames('text-2xl', 'font-bold', className)}
        />
    );
};

export default Subtitle;
