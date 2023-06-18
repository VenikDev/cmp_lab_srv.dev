// export const HOST_V1 = "http://localhost:8080/api/v1"
export const HOST_V1 = "/api/v1"
//
// This code declares an enum called "TypeRequest" with two possible values:
// POST and GET. Each value is assigned a string value: 'POST' and 'GET',
// respectively. The enum allows these values to be used as type annotations
// in TypeScript, ensuring that only these specific values are used when
// defining variables, function parameters, or return types.
export enum TypeRequest {
  POST = 'POST',
  GET = 'GET'
}