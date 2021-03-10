import axios, { Method } from "axios";
import GeneratorRequest from "../models/generatorRequest";

const BASE_URI = process.env.VUE_APP_BASE_API_URI;
console.log(process.env.VUE_APP_BASE_API_URI);
const client = axios.create({
  baseURL: BASE_URI
});

const APIClient = {
  generate(request: GeneratorRequest) {
    return this.perform("post", `/generate`, JSON.stringify(request));
  },

  async perform(method: Method, resource: string, data: string) {
    return client({
      method,
      url: resource,
      data: data,
      headers: {}
    }).then(req => {
      return req.data;
    });
  }
};

export default APIClient;
