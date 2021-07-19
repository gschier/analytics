import React, { HTMLAttributes } from 'react';
import classnames from 'classnames';
import { Link as RouterLink } from 'react-router-dom';
import Button from './Button';

export interface LinkProps {
  to: string;
  button?: boolean;
  external?: boolean;
}

const Link: React.FC<LinkProps & HTMLAttributes<HTMLElement>> = ({
  className,
  children,
  button,
  to,
  external,
  ...props
}) => {
  const allProps = {
    className: classnames(
      'text-primary-500',
      'hover:text-primary-600 hover:underline',
      className,
    ),
    ...props,
  };

  const realChildren = button ? <Button>{children}</Button> : children;

  const hasProto = to.match(/^https?:\/\//);
  if (external || hasProto) {
    return (
      <a
        href={hasProto ? to : `https://${to}`}
        target="_blank"
        rel="noopener noreferrer"
        {...allProps}>
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
