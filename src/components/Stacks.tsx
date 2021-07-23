import classnames from 'classnames';
import React, { HTMLAttributes } from 'react';

type StackSpace = 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 12 | 14 | 16;

const spaces: Record<StackSpace, string> = {
  0: 'gap-0',
  1: 'gap-1',
  2: 'gap-2',
  3: 'gap-3',
  4: 'gap-4',
  5: 'gap-5',
  6: 'gap-6',
  7: 'gap-7',
  8: 'gap-8',
  9: 'gap-8',
  10: 'gap-8',
  12: 'gap-12',
  14: 'gap-14',
  16: 'gap-16',
};

export interface StackProps extends HTMLAttributes<HTMLDivElement> {
  justify?: JustifyValue;
  align?: AlignValue;
  space?: StackSpace;
  className?: string;
}

type JustifyValue = 'center' | 'between' | 'start' | 'end';
type AlignValue = 'center' | 'start' | 'end' | 'baseline';
type WrapValue = true | false | 'reverse';

const justifyValues: Record<JustifyValue, string> = {
  center: 'justify-center',
  between: 'justify-between',
  start: 'justify-start',
  end: 'justify-end',
};

const alignValues: Record<AlignValue, string> = {
  center: 'items-center',
  baseline: 'items-baseline',
  start: 'items-start',
  end: 'items-end',
};

const wrapValue = (v?: WrapValue): string => {
  if (v === true) return 'flex-wrap';
  else if (v === false) return 'flex-nowrap';
  else if (v === 'reverse') return 'flex-wrap-reverse';
  else return '';
};

export interface StackProps {
  space?: StackSpace;
  justify?: JustifyValue;
  align?: AlignValue;
  wrap?: WrapValue;
}

export const VStack: React.FC<StackProps & HTMLAttributes<HTMLDivElement>> = ({
  className,
  space,
  justify,
  align,
  wrap,
  ...props
}) => (
  <div
    {...props}
    className={classnames(
      'flex flex-col',
      space !== undefined ? spaces[space] : '',
      justify ? justifyValues[justify] : '',
      align ? alignValues[align] : '',
      wrapValue(wrap),
      className,
    )}
  />
);

export const HStack: React.FC<
  StackProps & { collapse?: boolean } & HTMLAttributes<HTMLDivElement>
> = ({ className, space, justify, align, wrap, collapse, ...props }) => (
  <div
    {...props}
    className={classnames(
      'w-full flex',
      collapse && 'flex-col sm:flex-row',
      !collapse && 'flex-row',
      space === undefined ? '' : spaces[space],
      justifyValues[justify ?? 'start'],
      alignValues[align ?? 'end'],
      wrapValue(wrap),
      className,
    )}
  />
);
