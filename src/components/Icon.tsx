import React from 'react';
import * as icons from '@heroicons/react/outline';
import classnames from 'classnames';

export type IconChoice = 'clock' | 'copy';
export type IconSize = 'sm' | 'md' | 'lg';

const sizeClassMap: Record<IconSize, string> = {
  sm: classnames('w-5 h-5'),
  md: classnames('w-6 h-6'),
  lg: classnames('w-7 h-7'),
};

export interface IconProps {
  icon: IconChoice;
  size?: IconSize;
  className?: string;
}

const Icon: React.FC<IconProps> = ({ className, icon, size, ...props }) => {
  const allProps = {
    ...props,
    className: classnames('w-6 h-6', sizeClassMap[size ?? 'md'], className),
  };
  switch (icon) {
    case 'clock':
      return <icons.ClockIcon {...allProps} />;
    case 'copy':
      return <icons.ClipboardCopyIcon {...allProps} />;
    default:
      throw new Error(`Unknown icon ${icon}`);
  }
};

export default Icon;
