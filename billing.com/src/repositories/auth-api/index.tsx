import request from "@/pkg/api/request"
import { ApiSuccessResponse } from "@/pkg/api/response"
import { AuthenticatedUser, UserToken } from "./entity";

const Login = (email: string, password: string) :Promise<ApiSuccessResponse<UserToken>> => {
  return request({
    url: `user/login`,
    method: "POST",
    data:{
      email: email,
      password: password,
    }
  });
};

const GetAuthenticated = () :Promise<ApiSuccessResponse<AuthenticatedUser>> => {
  return request({
    url: `user/current`,
    method: "GET",
    headers: {needauth: true},
  });
};



export const AuthApi = {
  Login,
  GetAuthenticated
};
