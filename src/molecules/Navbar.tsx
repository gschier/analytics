import React, { HTMLAttributes } from 'react';

export interface NavbarProps {
}

const Navbar: React.FC<NavbarProps & HTMLAttributes<HTMLElement>> = ({
    className,
    ...props
}) => {
    return (
        <nav
            {...props}
            className={`${className ?? ''} rounded font-medium py-2 px-4`}>
        </nav>
    );
};

export default Navbar;
