import React, { HTMLAttributes } from 'react';
import classnames from 'classnames';
import Title from './Title';
import { HStack } from './Stacks';
import Link from './Link';
import Button from './Button';
import { Website } from '../hooks/use-websites';
import { useTheme } from '../hooks/use-theme';

export interface NavbarProps {
  website?: Website;
}

const Navbar: React.FC<NavbarProps & HTMLAttributes<HTMLElement>> = ({
  className,
  website,
  ...props
}) => {
  const [theme, setTheme] = useTheme();
  return (
    <nav
      {...props}
      className={classnames('rounded font-medium py-4 px-4 mb-4', className)}>
      <HStack>
        <Title>
          <HStack space={2} align="center">
            {website ? (
              <>
                <Link to="/">Home</Link>
                <div>/</div>
                <div>{website?.domain}</div>
              </>
            ) : (
              <div>Home</div>
            )}
          </HStack>
        </Title>
        <Button
          variant="ghost"
          className="ml-auto -mr-2 px-2"
          icon={theme === 'dark' ? 'moon' : 'sun'}
          onClick={() => setTheme(theme === 'dark' ? 'light' : 'dark')}
        />
      </HStack>
    </nav>
  );
};

export default Navbar;
