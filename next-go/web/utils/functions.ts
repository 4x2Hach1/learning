import { API_SERVER } from "./const";
import { ApiUserSchema } from "./types";

export const authUser = async (token: string) => {
  const res = await fetch(API_SERVER + "/user/auth", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Authorization: token,
    },
  });

  const body = await res.json();
  const data = ApiUserSchema.parse(body);

  return data;
};
