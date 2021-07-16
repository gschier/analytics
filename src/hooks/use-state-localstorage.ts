import { Dispatch, SetStateAction, useState } from 'react';

export default function useStateLocalStorage<T>(
    key: string,
    defaultValue: T,
): [ T, Dispatch<SetStateAction<T>> ] {
    const savedV = localStorage[key];
    const [ v, setV ] = useState<T>(savedV ?? defaultValue);
    const updateV = (newV: SetStateAction<T>) => {
        localStorage[key] = newV;
        setV(newV);
    };
    return [ v, updateV ];
}