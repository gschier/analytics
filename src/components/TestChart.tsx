import React, { useCallback, useMemo } from 'react';
import { AreaClosed, Bar, Line } from '@visx/shape';
import appleStock, { AppleStock } from '@visx/mock-data/lib/mocks/appleStock';
import { curveMonotoneX } from '@visx/curve';
import { GridColumns, GridRows } from '@visx/grid';
import { scaleLinear, scaleTime } from '@visx/scale';
import { defaultStyles, TooltipWithBounds, withTooltip } from '@visx/tooltip';
import { WithTooltipProvidedProps } from '@visx/tooltip/lib/enhancers/withTooltip';
import { localPoint } from '@visx/event';
import { LinearGradient } from '@visx/gradient';
import { bisector, extent, max } from 'd3-array';
import { timeFormat } from 'd3-time-format';

type TooltipData = {
  first: AppleStock;
  second: AppleStock;
  firstY: number;
  secondY: number;
};

export const background = 'hsl(var(--color-primary-50))';
export const background2 = 'hsl(var(--color-primary-50))';
export const accentColor = 'hsl(var(--color-primary-400))';
export const accentColor2 = 'hsl(var(--color-primary-300))';
export const accentColorTwo = 'hsl(var(--color-primary-500))';
export const accentColorTwo2 = 'hsl(var(--color-primary-400))';
export const accentColorDark = 'hsl(var(--color-primary-600))';
const tooltipStyles = {
  ...defaultStyles,
  background,
  border: '1px solid hsl(var(--color-gray-100))',
  color: 'hsl(var(--color-gray-800))',
};

// util
const formatDate = timeFormat('%B %d');

// accessors
const getDate = (d: AppleStock) => new Date(d.date);
const getStockValue = (d: AppleStock) => d.close;
const bisectDate = bisector<AppleStock, Date>((d) => new Date(d.date)).left;

export type AreaProps = {
  width: number;
  height: number;
  margin?: { top: number; right: number; bottom: number; left: number };
  data?: AppleStock[];
  data2?: AppleStock[];
};

export default withTooltip<AreaProps, TooltipData>(
  ({
    width,
    height,
    margin = { top: 0, right: 0, bottom: 0, left: 0 },
    showTooltip,
    hideTooltip,
    tooltipData,
    tooltipTop = 0,
    tooltipLeft = 0,
    data: stock = appleStock.slice(800),
    data2: stock2 = appleStock.slice(800),
  }: AreaProps & WithTooltipProvidedProps<TooltipData>) => {
    if (width < 10) return null;
    if (stock.length === 0) return null;

    // stock?.forEach(d => d.close = 10);

    // bounds
    const innerWidth = width - margin.left - margin.right;
    const innerHeight = height - margin.top - margin.bottom;

    // scales
    const dateScale = useMemo(
      () =>
        scaleTime({
          range: [margin.left, innerWidth + margin.left],
          domain: extent(stock, getDate) as [Date, Date],
        }),
      [innerWidth, margin.left, stock, stock2],
    );
    const stockValueScale = useMemo(
      () =>
        scaleLinear({
          range: [innerHeight + margin.top, margin.top],
          domain: [0, (max(stock2, getStockValue) || 0) + innerHeight / 10],
          nice: true,
        }),
      [margin.top, innerHeight, stock, stock2],
    );

    // tooltip handler
    const handleTooltip = useCallback(
      (
        event:
          | React.TouchEvent<SVGRectElement>
          | React.MouseEvent<SVGRectElement>,
      ) => {
        const { x } = localPoint(event) || { x: 0 };
        const x0 = dateScale.invert(x);
        const index = bisectDate(stock, x0, 1);
        const index2 = bisectDate(stock2, x0, 1);
        const d0 = stock[index - 1];
        const d1 = stock[index];
        const d20 = stock2[index2 - 1];
        const d21 = stock2[index2];
        let d = d0;
        let d2 = d20;
        if (d1 && getDate(d1)) {
          d =
            x0.valueOf() - getDate(d0).valueOf() >
            getDate(d1).valueOf() - x0.valueOf()
              ? d1
              : d0;
        }
        if (d21 && getDate(d21)) {
          d2 =
            x0.valueOf() - getDate(d20).valueOf() >
            getDate(d1).valueOf() - x0.valueOf()
              ? d21
              : d20;
        }
        const top1 = stockValueScale(getStockValue(d));
        const top2 = stockValueScale(getStockValue(d2));
        const topAvg = (top2 + top1) / 2;
        showTooltip({
          tooltipData: {
            first: d,
            firstY: top1,
            second: d2,
            secondY: top2,
          },
          tooltipLeft: x,
          tooltipTop: topAvg,
        });
      },
      [showTooltip, stockValueScale, dateScale],
    );

    return (
      <div style={{ touchAction: 'pan-y' }}>
        <svg width={width} height={height} className="rounded">
          <rect
            x={0}
            y={0}
            width={width}
            height={height}
            fill="url(#area-background-gradient)"
          />
          <LinearGradient
            id="area-background-gradient"
            from={background}
            to={background2}
          />
          <LinearGradient
            id="area-gradient"
            from={accentColor}
            to={accentColor2}
          />
          <LinearGradient
            id="area-gradient-2"
            from={accentColorTwo}
            to={accentColorTwo2}
          />
          <GridRows
            left={margin.left}
            scale={stockValueScale}
            width={innerWidth}
            strokeDasharray="1,3"
            stroke={accentColor}
            strokeOpacity={0}
            pointerEvents="none"
          />
          <GridColumns
            top={margin.top}
            scale={dateScale}
            height={innerHeight}
            strokeDasharray="1,3"
            stroke={accentColor}
            strokeOpacity={0.2}
            pointerEvents="none"
          />
          <AreaClosed<AppleStock>
            data={stock2}
            x={(d) => dateScale(getDate(d)) ?? 0}
            y={(d) => stockValueScale(getStockValue(d)) ?? 0}
            yScale={stockValueScale}
            strokeWidth={1}
            stroke="url(#area-gradient)"
            fill="url(#area-gradient)"
            curve={curveMonotoneX}
          />
          <AreaClosed<AppleStock>
            data={stock}
            x={(d) => dateScale(getDate(d)) ?? 0}
            y={(d) => stockValueScale(getStockValue(d)) ?? 0}
            yScale={stockValueScale}
            strokeWidth={1}
            stroke="url(#area-gradient-2)"
            fill="url(#area-gradient-2)"
            curve={curveMonotoneX}
          />
          <Bar
            x={margin.left}
            y={margin.top}
            width={innerWidth}
            height={innerHeight}
            fill="transparent"
            rx={14}
            onTouchStart={handleTooltip}
            onTouchMove={handleTooltip}
            onMouseMove={handleTooltip}
            onMouseLeave={() => hideTooltip()}
          />
          {tooltipData && (
            <g>
              <Line
                from={{ x: tooltipLeft, y: margin.top }}
                to={{
                  x: tooltipLeft,
                  y: innerHeight + margin.top,
                }}
                stroke={accentColorDark}
                strokeWidth={2}
                pointerEvents="none"
                strokeDasharray="5,2"
              />
              <circle
                cx={tooltipLeft}
                cy={tooltipData.firstY + 1}
                r={4}
                fill={accentColorTwo2}
                stroke="white"
                strokeWidth={2}
                pointerEvents="none"
              />
              <circle
                cx={tooltipLeft}
                cy={tooltipData.secondY + 1}
                r={4}
                fill={accentColor2}
                stroke="white"
                strokeWidth={2}
                pointerEvents="none"
              />
            </g>
          )}
        </svg>
        {tooltipData && (
          <div>
            <TooltipWithBounds
              key={Math.random()}
              top={tooltipTop - 38}
              left={tooltipLeft + 10}
              style={tooltipStyles}>
              <div className="font-semibold">
                <div className="text-xs pb-3 text-gray-700">
                  {formatDate(getDate(tooltipData.first))}
                </div>
                <div>
                  <span className="text-primary-400">Views</span>{' '}
                  {getStockValue(tooltipData.second)}
                </div>
                <div>
                  <span className="text-primary-500">Visitors</span>{' '}
                  {getStockValue(tooltipData.first)}
                </div>
              </div>
            </TooltipWithBounds>
            {/*<Tooltip*/}
            {/*  top={innerHeight + margin.top - 4}*/}
            {/*  left={tooltipLeft - 8}*/}
            {/*  style={{*/}
            {/*    ...tooltipStyles,*/}
            {/*    minWidth: 72,*/}
            {/*    textAlign: 'center',*/}
            {/*    transform: 'translateX(-50%)',*/}
            {/*  }}>*/}
            {/*  <span className="text-xs mb-2 text-gray-700">*/}
            {/*    {formatDate(getDate(tooltipData.first))}*/}
            {/*  </span>*/}
            {/*  <br />*/}
            {/*  Visitors {getStockValue(tooltipData.first)}*/}
            {/*  <br />*/}
            {/*  Views {getStockValue(tooltipData.second)}*/}
            {/*</Tooltip>*/}
          </div>
        )}
      </div>
    );
  },
);
