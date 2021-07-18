import React, { HTMLAttributes } from 'react';
import classnames from 'classnames';
import { Link as RouterLink } from 'react-router-dom';
import Button from './Button';

export interface LinkProps {
  to: string;
  button?: boolean;
}

const Link: React.FC<LinkProps & HTMLAttributes<HTMLElement>> = ({
  className,
  children,
  button,
  to,
  ...props
}) => {
  const allProps = {
    className: classnames(
      'text-secondary-500',
      'hover:text-secondary-700 hover:underline',
      className,
    ),
    ...props,
  };

  const realChildren = button ? <Button>{children}</Button> : children;

  if (to.match(/^https?:\/\//)) {
    return (
      <a href={to} {...allProps}>
        {realChildren}
      </a>
    );
  }

  return (
    <RouterLink to={to} {...allProps}>
      {realChildren}
    </RouterLink>
  );
};

export default Link;
