import React, { HTMLAttributes } from 'react';

export type ButtonColor = 'primary' | 'secondary' | 'danger' | 'gray';
export type ButtonVariant = 'solid' | 'outline';
export type ButtonSize = 'sm' | 'md' | 'lg';

const colorClassMap: Record<ButtonVariant, Record<ButtonColor, string>> = {
    solid: {
        primary: [
            'text-primary-0 bg-primary-500',
            'hover:bg-primary-600',
            'focus-visible:outline-none focus-visible:ring-4 focus-visible:ring-offset-0 focus-visible:ring-primary-500 focus-visible:ring-opacity-50',
        ].join(' '),
        secondary: [
            'text-secondary-0 bg-secondary-500',
            'hover:bg-secondary-600',
            'focus-visible:outline-none focus-visible:ring-4 focus-visible:ring-offset-0 focus-visible:ring-secondary-500 focus-visible:ring-opacity-50',
        ].join(' '),
        danger: [
            'text-danger-0 bg-danger-500',
            'hover:bg-danger-600',
            'focus-visible:outline-none focus-visible:ring-4 focus-visible:ring-offset-0 focus-visible:ring-danger-500 focus-visible:ring-opacity-50',
        ].join(' '),
        gray: [
            'text-gray-0 bg-gray-500',
            'hover:bg-gray-600',
            'focus-visible:outline-none focus-visible:ring-4 focus-visible:ring-offset-0 focus-visible:ring-gray-500 focus-visible:ring-opacity-50',
        ].join(' '),
    },
    outline: {
        primary: [
            'text-primary-500 ring-1 ring-primary-500',
            'hover:bg-primary-50 hover:text-primary-600 hover:ring-primary-600',
            'focus-visible:outline-none focus-visible:ring-4 focus-visible:ring-offset-0 focus-visible:ring-primary-400 focus-visible:ring-opacity-50',
        ].join(' '),
        secondary: [
            'text-secondary-500 ring-1 ring-secondary-500',
            'hover:bg-secondary-50 hover:text-secondary-600 hover:ring-secondary-600',
            'focus-visible:outline-none focus-visible:ring-4 focus-visible:ring-offset-0 focus-visible:ring-secondary-400 focus-visible:ring-opacity-50',
        ].join(' '),
        danger: [
            'text-danger-500 ring-1 ring-danger-500',
            'hover:bg-danger-50 hover:text-danger-600 hover:ring-danger-600',
            'focus-visible:outline-none focus-visible:ring-4 focus-visible:ring-offset-0 focus-visible:ring-danger-400 focus-visible:ring-opacity-50',
        ].join(' '),
        gray: [
            'text-gray-500 ring-1 ring-gray-500',
            'hover:bg-gray-50 hover:text-gray-600 hover:ring-gray-600',
            'focus-visible:outline-none focus-visible:ring-4 focus-visible:ring-offset-0 focus-visible:ring-gray-400 focus-visible:ring-opacity-50',
        ].join(' '),
    },
};

const sizeClassMap: Record<ButtonSize, string> = {
    sm: 'text-xs px-3',
    md: 'text-sm px-3.5',
    lg: 'text-md px-4',
};

export interface ButtonProps {
    color?: ButtonColor;
    variant?: ButtonVariant;
    size?: ButtonSize;
}

const Button: React.FC<ButtonProps & HTMLAttributes<HTMLButtonElement>> = ({
    className,
    variant,
    color,
    size,
    ...props
}) => {
    const baseClass = 'rounded py-1.5 font-medium transform active:scale-95 transition-transform duration-75';
    const colorClass = colorClassMap[variant ?? 'solid'][color ?? 'primary'];
    const sizeClass = sizeClassMap[size ?? 'md'];
    return (
        <button
            {...props}
            className={`${className ?? ''} ${baseClass} ${colorClass} ${sizeClass}`} />
    );
};

export default Button;
