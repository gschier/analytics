import React, { HTMLAttributes, ReactNode } from 'react';
import { VStack } from './Stacks';

export type InputSize = 'sm' | 'md' | 'lg';

const sizeClassMap: Record<InputSize, string> = {
    sm: 'text-xs',
    md: 'text-sm',
    lg: 'text-md',
};

export interface InputProps {
    type?: 'text' | 'number' | 'email' | 'url';
    size?: InputSize;
    label?: ReactNode;
    textarea?: boolean;
    error?: string;
    defaultValue?: string;
}

const Input: React.FC<InputProps & HTMLAttributes<HTMLInputElement | HTMLTextAreaElement>> = ({
    className,
    label,
    size,
    type,
    textarea,
    error,
    ...props
}) => {
    const sizeClass = sizeClassMap[size ?? 'md'];
    const errorClass = error && [
        '!ring-danger-500 !focus:ring-danger-500 text-danger-800',
    ].join(' ');
    const baseClass = [
        'bg-gray-0 rounded py-1.5 px-3 ring-1 ring-gray-200 placeholder-gray-300 w-full',
        'focus:outline-none focus:ring-2 focus:ring-offset-0 focus:ring-primary-500',
    ].join(' ');

    return (
        <VStack className={className}>
            {label && <label className="font-semibold text-gray-400 text-sm mb-1">{label}</label>}
            {textarea ? (
                <textarea {...props} className={`${baseClass} ${sizeClass} ${errorClass} h-16`} />
            ) : (
                <input {...props} type={type} className={`${baseClass} ${sizeClass} ${errorClass}`} />
            )}
            {error && <div className="text-danger-500 text-sm mt-1">{error}</div>}
        </VStack>
    );
};

export default Input;