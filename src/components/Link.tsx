import React, { HTMLAttributes } from 'react';
import classnames from 'classnames';
import { Link as RouterLink } from 'react-router-dom';

export interface LinkProps {
  to: string;
}

const Link: React.FC<LinkProps & HTMLAttributes<HTMLElement>> = ({
  className,
  to,
  ...props
}) => {
  const allProps = {
    className: classnames(
      'text-secondary-500',
      'hover:text-secondary-700',
      'hover:underline',
      className,
    ),
    ...props,
  };

  if (to.match(/^https?:\/\//)) {
    return <a href={to} {...allProps} />;
  } else {
    return <RouterLink to={to} {...allProps} />;
  }
};

export default Link;
