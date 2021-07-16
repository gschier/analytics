import Text from './Text';
import React, { Children, HTMLAttributes, ReactNode, useMemo } from 'react';

export interface TableProps {
    children: (ReactNode[])[];
    columns?: string[];
}

const cellClass = 'px-3 py-2';

export const Table: React.FC<TableProps & HTMLAttributes<HTMLElement>> = ({
    className,
    children,
    columns,
    ...props
}) => {
    return (
        <div className={`${className ?? ''} border border-gray-100 w-full overflow-hidden rounded`}>
            <table {...props} className="w-full divide-y divide-gray-100">
                {columns && (
                    <thead className="uppercase">
                    <tr>
                        {columns.map((c, i) => (
                            <th key={i} className={`font-medium text-left bg-gray-50 ${cellClass}`}>
                                <Text size="sm">{c}</Text>
                            </th>
                        ))}
                    </tr>
                    </thead>
                )}
                <tbody>
                {children.map((r, i) => (
                    <TableRow key={i} className={i % 2 === 1 ? 'bg-gray-50' : ''}>
                        {r}
                    </TableRow>
                ))}
                </tbody>
            </table>
        </div>
    );
};

export const TableRow: React.FC<{ children: ReactNode[] } & HTMLAttributes<HTMLElement>> = ({
    children,
    ...props
}) => {
    return (
        <tr {...props}>
            {React.Children.map(children, ((td, i) => (
                <td key={i} className={`${cellClass} text-gray-700`}>
                    {td}
                </td>
            )))}
        </tr>
    );
};

export const TableCell: React.FC<HTMLAttributes<HTMLTableDataCellElement>> = ({
    className,
    ...props
}) => {
    return (
        <td {...props} className={`${className ?? ''} ${cellClass} text-gray-700`} />
    );
};
