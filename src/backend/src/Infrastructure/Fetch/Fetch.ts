import axios from 'axios';
import {FetchError} from "./Error";

export class Fetch {
    static instance;

    static init() {
        axios.interceptors.request.use(function (config) {
            return config;
        }, function (error) {
            return Promise.reject(error);
        });

        axios.interceptors.response.use(function (response) {
            return response;
        }, function (error) {
            return Promise.reject(error);
        });

        Fetch.instance = axios.create({
            baseURL: 'http://localhost',
            timeout: 15000,
        });
    }

    static async get<T>(url: string): Promise<T | FetchError> {
        try {
            const response = await this.instance.get(url);
            return response.data;
        } catch(err) {
            return new FetchError(url, err.message);
        }
    }
}