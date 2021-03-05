import axios, { AxiosRequestConfig } from 'axios';
import * as generatedSchema from './generated-schema';

export type schema = generatedSchema.definitions;

export const $axios = axios.create();

export function getItems(config?: AxiosRequestConfig) {
  return $axios.get<schema['service.ItemsListResponse']>('/api/items', config);
}
