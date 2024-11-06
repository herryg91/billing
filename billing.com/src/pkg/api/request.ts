import { UsertokenApi } from "@/repositories/usertoken-api";
import axios, { AxiosError, AxiosResponse } from "axios";
import { ApiErrorResponse, ForbiddenError, isApiResponseError, UnauthorizedError, UnknownError } from "./response";
import { getCookie, deleteCookie, setCookie } from "cookies-next";

const KEY_AUTH_TOKEN = "auth_token"
const KEY_REFRESH_TOKEN = "refresh_token"
const LoginPageRoute = "/login"


const client = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_HOST,
});

client.interceptors.request.use(
  (request) => {
    if(request === undefined || request.headers === undefined) 
      return request;
    const current_auth_token = getCookie(KEY_AUTH_TOKEN) ?? ""
    if(current_auth_token !== "") {
      request.headers['Authorization'] = "Bearer " + current_auth_token
    }

    if(request.headers['needauth'] !== null && request.headers['needauth'] !== undefined && request.headers['needauth'] !== ''){
      delete request.headers['needauth'] 
      request.params = (request.params === undefined || request.params === null) ? {__retry: 0} : {...request.params, __retry: 0}
    }

    return request;
  },
  (error) => {
      return Promise.reject(error);
  }
);

client.interceptors.response.use(
  (response) => {
      // console.log('Response:', response)
      return response;
  },
  async (error: AxiosError<ApiErrorResponse>) => {
    const max_retry = 3
    if (error.response?.status !== 401) return Promise.reject(error)

    if (Number(error.config?.params.__retry) >= max_retry) {
      // logout
      deleteCookie(KEY_AUTH_TOKEN)
      deleteCookie(KEY_REFRESH_TOKEN)
      window.location.href = LoginPageRoute
      return Promise.reject(error)
    }
      
      // Try request again with new token
      const currentRefreshToken = getCookie(KEY_REFRESH_TOKEN) ?? ""
      try {        
          const new_token = await UsertokenApi.RefreshToken(currentRefreshToken)
          setCookie(KEY_AUTH_TOKEN, new_token.data.auth_token)
          setCookie(KEY_REFRESH_TOKEN, new_token.data.refresh_token)

          const config = error.config;
          if(config === undefined) {
            deleteCookie(KEY_AUTH_TOKEN)
            deleteCookie(KEY_REFRESH_TOKEN)
            window.location.href = "/login";
            return Promise.reject("invalid config, cannot retry refresh token");
          }

          config.headers['Authorization'] = "Bearer " + new_token.data.auth_token;
          config.params.__retry = 1 + Number(error.config?.params.__retry)
          return new Promise((resolve, reject) => {
            client.request(config).then((response) => {
                  resolve(response);
              }).catch( (error) => {
                  reject(error);
              })
          });
      } catch (err) {
          if(isApiResponseError(err)){
              if(err.http_status === 403 && err.code === 11001){
                  deleteCookie(KEY_AUTH_TOKEN)
                  deleteCookie(KEY_REFRESH_TOKEN)
                  window.location.href = "/login";
              }
          }
          return Promise.reject(err);
      } 
  }
);

/* eslint-disable @typescript-eslint/no-explicit-any */
const request = (options: any) => {
  const onSuccess = (response: AxiosResponse) => {
    return response.data;
  };

  const onError = function (error: any) :Promise<ApiErrorResponse>{
    if (error.response.status === 401) {
      return Promise.reject(UnauthorizedError);
    }
   
    if (error.response.status === 403) {
      if(isApiResponseError(error.response.data)){
        if(error.response.data.http_status === 403 && error.response.data.code === 11001){
          return Promise.reject(error.response.data);
        }
        return Promise.reject(error.response.data);
      }
      return Promise.reject(ForbiddenError);
    }

    if(isApiResponseError(error.response.data)){
      return Promise.reject(error.response.data as ApiErrorResponse);
    }
    if(error.message) {
      return Promise.reject(UnknownError(error.message));
    }
    return Promise.reject(UnknownError(error.toString()));
  };

  return client({
    ...options,
  })
    .then(onSuccess)
    .catch(onError);
};

export default request;
