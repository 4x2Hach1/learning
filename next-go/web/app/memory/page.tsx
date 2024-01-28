"use client";

import { API_SERVER } from "@/utils/const";
import { authUser } from "@/utils/functions";
import { ApiMemories, ApiMemoriesSchema, ApiUser } from "@/utils/types";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import Heavy from "./heavy";

export default function Page() {
  const route = useRouter();
  const [user, setUser] = useState<ApiUser>();
  const [memories, setMemories] = useState<ApiMemories>();

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

          const res = await fetch(API_SERVER + "/memories", {
            method: "GET",
            headers: {
              "Content-Type": "application/json",
              Authorization: token,
            },
          });

          const body = await res.json();
          const resMemories = ApiMemoriesSchema.parse(body);
          setMemories(resMemories);
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
        <p>memory</p>
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
        <Heavy />
        <p className="my-8">
          <a href="/memory/new">投稿</a>
        </p>
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
            {memories?.map((memory, index) => (
              <tr key={index}>
                <td>
                  <a href={`/memory/${memory.id}`}>{memory.id}</a>
                </td>
                <td>{memory.title}</td>
                <td>{memory.date}</td>
                <td>{memory.location}</td>
                <td>{memory.detail}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </main>
  );
}
