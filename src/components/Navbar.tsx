import React, { HTMLAttributes } from 'react';
import classnames from 'classnames';

export interface NavbarProps {}

const Navbar: React.FC<NavbarProps & HTMLAttributes<HTMLElement>> = ({
  className,
  ...props
}) => (
  <nav
    {...props}
    className={classnames('rounded font-medium py-2 px-4', className)}
  />
);

export default Navbar;
