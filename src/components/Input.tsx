import React, { HTMLAttributes, ReactNode, useMemo } from 'react';
import { VStack } from './Stacks';

export type InputSize = 'sm' | 'md' | 'lg';

const sizeClassMap: Record<InputSize, string> = {
    sm: 'text-xs',
    md: 'text-sm',
    lg: 'text-md',
};

export interface InputProps {
    name?: string;
    type?: 'text' | 'number' | 'email' | 'url' | 'password';
    size?: InputSize;
    label?: ReactNode;
    textarea?: boolean;
    error?: string;
    defaultValue?: string;
    autoFocus?: boolean;
}

const Input: React.FC<InputProps & HTMLAttributes<HTMLInputElement | HTMLTextAreaElement>> = ({
    className,
    label,
    size,
    type,
    textarea,
    error,
    name,
    ...props
}) => {
    const id = useMemo(() => `input-${name ?? 'unknown'}-${Math.random()}`, []);
    const sizeClass = sizeClassMap[size ?? 'md'];
    const errorClass = error && '!ring-danger-400 !focus:ring-danger-400 text-danger-800 text-opacity-80';
    const baseClass = [
        'bg-gray-0 rounded py-1.5 px-3 ring-1 ring-gray-200 placeholder-gray-300 w-full',
        'focus:outline-none focus:ring-2 focus:ring-offset-0 focus:ring-primary-400',
    ].join(' ');

    return (
        <VStack className={className}>
            {label && <label htmlFor={id} className="font-semibold text-gray-500 text-sm mb-1">{label}</label>}
            {textarea ? (
                <textarea {...props} id={id} name={name} className={`${baseClass} ${sizeClass} ${errorClass} h-16`} />
            ) : (
                <input {...props} id={id} name={name} type={type} className={`${baseClass} ${sizeClass} ${errorClass}`} />
            )}
            {error && <div className="text-danger-500 text-sm mt-1">{error}</div>}
        </VStack>
    );
};

export default Input;
