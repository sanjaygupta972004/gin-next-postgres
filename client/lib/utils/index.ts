/* eslint-disable */
import { pickBy } from 'lodash';
import qs from 'qs';

const customPredicate = (value: any, _key: string) => {
  return value !== undefined || value === 0;
};

export const stringifyParams = (data: any) => {
  const { params, option } = data;
  const newParams = pickBy(params, customPredicate);
  return qs.stringify(newParams, {
    encode: false,
    skipNulls: true,
    strictNullHandling: true,
    ...option,
  });
};

export const formatNumberWithUnits = (num: number, decimals: number = 2) => {
  if (num === 0) return "0";

  const units = [
    { unit: "Q", threshold: 1e15 }, // Quadrillion
    { unit: "T", threshold: 1e12 }, // Trillion
    { unit: "B", threshold: 1e9 },  // Billion
    { unit: "M", threshold: 1e6 },  // Million
    { unit: "K", threshold: 1e3 },  // Thousand
    { unit: "", threshold: 1 },     // No unit
  ];

  const isNegative = num < 0;
  const absNum = Math.abs(num);

  let selectedUnit = units[units.length - 1]; // Default to no unit
  for (const unit of units) {
    if (absNum >= unit.threshold) {
      selectedUnit = unit;
      break;
    }
  }

  const scaled = absNum / selectedUnit.threshold;
  const formatted = scaled.toFixed(decimals).replace(/\.0+$/, "");

  return `${isNegative ? "-" : ""}${formatted}${selectedUnit.unit}`;
}