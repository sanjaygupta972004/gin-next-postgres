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
