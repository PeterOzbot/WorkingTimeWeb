import WorkingDay from "@/models/workingDay";
import axios, { Method } from "axios";
import handleDates from "@/utils/parseDate";
import GeneratorRequest from "../models/generatorRequest";

const BASE_URI = process.env.VUE_APP_BASE_API_URI;
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
