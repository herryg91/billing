import request from "@/pkg/api/request"
import { ApiSuccessResponse } from "@/pkg/api/response"
import { TokenClaim, UserToken } from "./param";

const ValidateToken = (auth_token: string) :Promise<ApiSuccessResponse<TokenClaim>> => {
    return request({
      url: `usertoken/validate`,
      method: "POST",
      data: {
        auth_token: auth_token
      },
    });
};

const RefreshToken = (refresh_token: string) :Promise<ApiSuccessResponse<UserToken>> => {
    return request({
      url: `usertoken/refresh`,
      method: "POST",
      data: {
        refresh_token: refresh_token
      },
    });
};

export const UsertokenApi = {
    ValidateToken,
    RefreshToken
};
  