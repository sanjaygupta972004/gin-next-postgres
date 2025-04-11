import axios, { AxiosError, AxiosInstance, AxiosRequestConfig, AxiosResponse } from 'axios';
import { camelizeKeys, decamelizeKeys } from 'humps';
import isNil from 'lodash.isnil';
import { stringifyParams } from '@/lib/utils';
import { CookiesStorage } from '@/lib/storage/cookie';

const defaultAxiosConfig: AxiosRequestConfig = {
  baseURL: process.env.API_URL,
  timeout: 90000,
  paramsSerializer: {
    // eslint-disable-next-line
    serialize: (params: any) => {
      return stringifyParams({
        params: decamelizeKeys({ ...params }),
        option: {
          encode: !isNil(params?.tags) || false,
        },
      });
    },
  },
};

export const getAccessToken = () => `Bearer ${CookiesStorage.getAccessToken()}`;

const transformResponse = (response: AxiosResponse) => {
  if (response?.data) {
    return { ...response, data: camelizeKeys(response.data) };
  }
  return response;
};

// eslint-disable-next-line
const wrapApiErrors = async (error: AxiosError, axiosInstance: AxiosInstance) => {
  const status = error.response?.status || error.status;
  if (!status) {
    throw new Error('Connection with API server is broken');
  }
  return Promise.reject(error.response?.data);
};

const api = axios.create({
  ...defaultAxiosConfig,
  headers: {
    ...defaultAxiosConfig.headers,
  },
});

api.interceptors.request.use((config) => {
  const authorization = getAccessToken();
  if (!authorization.includes('null')) {
    config.headers!.Authorization = authorization;
  }
  if (config.data instanceof FormData) {
    return config;
  }
  config.headers['Content-Type'] = 'application/json';
  // if (config.data) {
  //   config.data = { ...decamelizeKeys(config.data) };
  // }
  // if (config.params) {
  //   config.params = { ...decamelizeKeys(config.params) };
  // }
  return config;
});

api.interceptors.response.use(transformResponse, (error) => {
  const errorResponse = error;

  return wrapApiErrors(errorResponse, api);
});

export default api;
