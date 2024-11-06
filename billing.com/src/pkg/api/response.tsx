export type ApiSuccessResponse<T> = {
    http_status: number;
    process_time: string;
    data: T;
  }
  export type ApiErrorResponse = {
    http_status:  number;
    grpc_status:  number;
    code: number;
    message: string;
    other_errors: ApiErrorOther[];
  }
  
  export type ApiErrorOther = {
    code: number;
    field: string;
    message: string;
  }
  
  export const UnknownError = (err :string) => {
    const errUnknown: ApiErrorResponse= {
        http_status:  503,
        grpc_status:  14,
        code: 999999,
        message: "Service Unavailable: " + err,
        other_errors: [],
    }
    return errUnknown;
  }
  
  export const NetworkError: ApiErrorResponse = {
    http_status: 503,
    grpc_status: 14,
    code: 999503,
    message: "Service Unavailable",
    other_errors: [],
  }
  
  export const UnauthorizedError: ApiErrorResponse = {
    http_status: 401,
    grpc_status: 16,
    code: 999401,
    message: "You are unauthorized",
    other_errors: [],
  }
  
  export const ForbiddenError: ApiErrorResponse = {
    http_status: 403,
    grpc_status: 7,
    code: 999403,
    message: "You don't have access",
    other_errors: [],
  }
  
  /* eslint-disable @typescript-eslint/no-explicit-any */
  export function isApiResponseError(input: any): input is ApiErrorResponse {
    if (!input){
      return false
    }
    return typeof input.http_status === 'number'
        && input.http_status !== undefined
        && input.http_status !== null
        && input.http_status > 0
        && typeof input.message === 'string'
  }
    