import type WorkingDay from "@/models/workingDay";
import axios, { type Method } from "axios";
import handleDates from "@/utils/parseDate";
import type GeneratorRequest from "../models/generatorRequest";
import type CreateRequest from "@/models/createRequest";

const BASE_URI = "http://localhost:8080"//process.env.VUE_APP_BASE_API_URI; TODO
const client = axios.create({
  baseURL: BASE_URI
});

client.interceptors.response.use(originalResponse => {
  handleDates(originalResponse.data);
  return originalResponse;
});

const APIClient = {
  generate(request: GeneratorRequest): Promise<WorkingDay[]> {
    return this.perform<WorkingDay[]>(
      "post",
      `/generate`,
      JSON.stringify(request)
    );
  },

  create(request: CreateRequest): Promise<string> {
    return this.perform<string>(
      "post",
      `/create`,
      JSON.stringify(request)
    );
  },

  async perform<T>(method: Method, resource: string, data: string): Promise<T> {
    return client({
      method,
      url: resource,
      data: data,
      headers: {}
    }).then<T>(req => {
      return req.data;
    });
  }
};

export default APIClient;
