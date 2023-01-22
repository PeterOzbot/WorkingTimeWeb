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

  create(request: CreateRequest): Promise<void> {
    // return this.perform<string>(
    //   "post",
    //   `/create`,
    //   JSON.stringify(request)
    // );
    return client({
      method: 'POST',
      url: '/create',
      responseType: 'blob', // important
      data: JSON.stringify(request),
      headers: {}
    }).then(response => {


      //return response.data;

      const fileName = response.headers["file-name"]?.toString() ?? "file.xlsx";

      // create file link in browser's memory
      const href = URL.createObjectURL(response.data);

      // create "a" HTML element with href to file & click
      const link = document.createElement('a');
      link.href = href;
      link.setAttribute('download', fileName); //or any other extension
      document.body.appendChild(link);
      link.click();

      // clean up "a" element & remove ObjectURL
      document.body.removeChild(link);
      URL.revokeObjectURL(href);
    });
  },

  async perform<T>(method: Method, resource: string, data: string): Promise<T> {
    return client({
      method,
      url: resource,
      data: data,
      headers: {}
    }).then<T>(response => {
      return response.data;
    });
  }
};

export default APIClient;
