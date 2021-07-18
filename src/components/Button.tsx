import React, { HTMLAttributes } from 'react';
import classnames from 'classnames';
import Icon, { IconChoice } from './Icon';

export type ButtonColor = 'primary' | 'secondary' | 'danger' | 'gray';
export type ButtonVariant = 'solid' | 'outline';
export type ButtonSize = 'sm' | 'md' | 'lg';

const colorClassMap: Record<ButtonVariant, Record<ButtonColor, string>> = {
  solid: {
    primary: classnames(
      'text-white bg-primary-500 border border-primary-500',
      'focus-visible:outline-none focus-visible:ring-4 focus-visible:ring-offset-0 focus-visible:ring-primary-500 focus-visible:ring-opacity-50',
      'hover:bg-primary-600 hover:border-primary-600',
    ),
    secondary: classnames(
      'text-white bg-secondary-500 border border-secondary-500',
      'focus-visible:outline-none focus-visible:ring-4 focus-visible:ring-offset-0 focus-visible:ring-secondary-500 focus-visible:ring-opacity-50',
      'hover:bg-secondary-600 hover:border-secondary-600',
    ),
    danger: classnames(
      'text-white bg-danger-500 border border-danger-500',
      'focus-visible:outline-none focus-visible:ring-4 focus-visible:ring-offset-0 focus-visible:ring-danger-500 focus-visible:ring-opacity-50',
      'hover:bg-danger-600 hover:border-danger-600',
    ),
    gray: classnames(
      'text-white bg-gray-500 border border-gray-500',
      'focus-visible:outline-none focus-visible:ring-4 focus-visible:ring-offset-0 focus-visible:ring-gray-500 focus-visible:ring-opacity-50',
      'hover:bg-gray-600 hover:border-gray-600',
    ),
  },
  outline: {
    primary: classnames(
      'text-primary-500 ring-1 ring-primary-500',
      'hover:bg-primary-50 hover:text-primary-600 hover:ring-primary-600',
      'focus-visible:outline-none focus-visible:ring-4 focus-visible:ring-offset-0 focus-visible:ring-primary-400 focus-visible:ring-opacity-50',
    ),
    secondary: classnames(
      'text-secondary-500 ring-1 ring-secondary-500',
      'hover:bg-secondary-50 hover:text-secondary-600 hover:ring-secondary-600',
      'focus-visible:outline-none focus-visible:ring-4 focus-visible:ring-offset-0 focus-visible:ring-secondary-400 focus-visible:ring-opacity-50',
    ),
    danger: classnames(
      'text-danger-500 ring-1 ring-danger-500',
      'hover:bg-danger-50 hover:text-danger-600 hover:ring-danger-600',
      'focus-visible:outline-none focus-visible:ring-4 focus-visible:ring-offset-0 focus-visible:ring-danger-400 focus-visible:ring-opacity-50',
    ),
    gray: classnames(
      'text-gray-500 ring-1 ring-gray-500',
      'hover:bg-gray-50 hover:text-gray-600 hover:ring-gray-600',
      'focus-visible:outline-none focus-visible:ring-4 focus-visible:ring-offset-0 focus-visible:ring-gray-400 focus-visible:ring-opacity-50',
    ),
  },
};

const sizeClassMap: Record<ButtonSize, string> = {
  sm: classnames('text-xs px-2.5'),
  md: classnames('text-sm px-3'),
  lg: classnames('text-md px-3.5'),
};

export interface ButtonProps {
  color?: ButtonColor;
  variant?: ButtonVariant;
  size?: ButtonSize;
  icon?: IconChoice;
}

const Button: React.FC<ButtonProps & HTMLAttributes<HTMLButtonElement>> = ({
  className,
  children,
  variant,
  color,
  size,
  icon,
  ...props
}) => {
  return (
    <button
      {...props}
      style={{ touchAction: 'manipulation' }}
      className={classnames(
        'rounded-md py-1.5 font-medium transform transition-transform duration-75 flex items-center justify-center',
        'active:scale-95',
        colorClassMap[variant ?? 'solid'][color ?? 'primary'],
        sizeClassMap[size ?? 'md'],
        className,
      )}>
      {icon && (
        <Icon
          icon={icon}
          className={classnames({ 'mr-1': !!children })}
          size="sm"
        />
      )}
      {children}
    </button>
  );
};

export default Button;
