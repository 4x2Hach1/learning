"use client";

import { API_SERVER } from "@/utils/const";
import { authUser } from "@/utils/functions";
import { ApiMemory, ApiMemorySchema, ApiUser } from "@/utils/types";
import { useParams, useRouter } from "next/navigation";
import { useEffect, useState } from "react";

export default function Page() {
  const route = useRouter();
  const params = useParams<{ id: string }>();
  const [user, setUser] = useState<ApiUser>();
  const [memory, setMemory] = useState<ApiMemory>();

  useEffect(() => {
    (async () => {
      const token = localStorage.getItem("jwt");
      if (!token) {
        localStorage.removeItem("jwt");
        route.push("/");
      } else {
        try {
          const resUser = await authUser(token);
          setUser(resUser);

          const res = await fetch(API_SERVER + `/memory/${params.id}`, {
            method: "GET",
            headers: {
              "Content-Type": "application/json",
              Authorization: token,
            },
          });

          const body = await res.json();
          const resMemory = ApiMemorySchema.parse(body);
          setMemory(resMemory);
        } catch (error) {
          console.log(error);
          route.push("/");
        }
      }
    })();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <div>
        <p>diary</p>
        <div>
          {user?.name ?? "---"}さん
          <button
            className="pl-8"
            onClick={() => {
              localStorage.removeItem("jwt");
              route.push("/");
            }}
          >
            ログアウト
          </button>
        </div>
        <table>
          <thead>
            <tr>
              <th>No.</th>
              <th>タイトル</th>
              <th>日付</th>
              <th>場所</th>
              <th>詳細</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>{memory?.id}</td>
              <td>{memory?.title}</td>
              <td>{memory?.date}</td>
              <td>{memory?.location}</td>
              <td>{memory?.detail}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </main>
  );
}
