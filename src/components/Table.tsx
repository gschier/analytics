import { Paragraph } from './Typography';
import React, { HTMLAttributes, ReactNode } from 'react';
import classnames from 'classnames';

export interface TableProps {
  children: ReactNode;
  columns?: string[];
}

const cellClass = classnames('px-3', 'py-1.5');

export const Table: React.FC<TableProps & HTMLAttributes<HTMLElement>> = ({
  className,
  children,
  columns,
  ...props
}) => {
  return (
    <div
      className={classnames(
        'border border-gray-100 w-full overflow-hidden rounded overflow-x-auto',
        className,
      )}>
      <table {...props} className="w-full divide-y divide-gray-100">
        {columns && (
          <thead className="uppercase">
            <tr>
              {columns.map((c, i) => (
                <th
                  key={i}
                  className={classnames(
                    'text-left bg-gray-50 py-2',
                    cellClass,
                  )}>
                  <Paragraph size="sm" className="text-gray-600">
                    {c}
                  </Paragraph>
                </th>
              ))}
            </tr>
          </thead>
        )}
        <tbody>{children}</tbody>
      </table>
    </div>
  );
};

export const TableRow: React.FC<
  { children: ReactNode[] } & HTMLAttributes<HTMLElement>
> = ({ children, ...props }) => {
  return (
    <tr {...props}>
      {React.Children.map(children, (contents, i) => (
        <td
          key={i}
          className={classnames(
            cellClass,
            'text-gray-700 truncate w-full max-w-0',
          )}>
          {contents}
        </td>
      ))}
    </tr>
  );
};
