import classnames from 'classnames';
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

const Input: React.FC<
    InputProps & HTMLAttributes<HTMLInputElement | HTMLTextAreaElement>
> = ({ className, label, size, type, textarea, error, name, ...props }) => {
    const id = useMemo(() => `input-${name ?? 'unknown'}-${Math.random()}`, []);
    const sizeClass = sizeClassMap[size ?? 'md'];
    const errorClass =
        error &&
        classnames(
            '!ring-danger-400',
            '!focus:ring-danger-400',
            'text-danger-800',
            'text-opacity-80',
        );
    const baseClass = classnames(
        'bg-gray-0',
        'rounded',
        'py-1.5',
        'px-3',
        'border',
        'border-gray-200',
        'placeholder-gray-300',
        'w-full',
        'focus:outline-none',
        'focus:ring-1',
        'focus:ring-offset-0',
        'focus:ring-primary-400',
        'focus:border-primary-400',
    );

    return (
        <VStack className={className}>
            {label && (
                <label
                    htmlFor={id}
                    className="font-semibold text-gray-500 text-sm mb-1">
                    {label}
                </label>
            )}
            <div>
                {textarea ? (
                    <textarea
                        {...props}
                        style={{ touchAction: 'manipulation' }}
                        id={id}
                        name={name}
                        className={classnames(
                            baseClass,
                            sizeClass,
                            errorClass,
                            'h-16',
                        )}
                    />
                ) : (
                    <input
                        {...props}
                        style={{ touchAction: 'manipulation' }}
                        id={id}
                        name={name}
                        type={type}
                        className={classnames(baseClass, sizeClass, errorClass)}
                    />
                )}
            </div>
            {error && (
                <div className="text-danger-500 text-sm mt-1">{error}</div>
            )}
        </VStack>
    );
};

export default Input;
