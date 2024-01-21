"use client";

import { API_SERVER } from "@/utils/const";
import { authUser } from "@/utils/functions";
import { ApiUser } from "@/utils/types";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";

export default function Page() {
  const route = useRouter();
  const [user, setUser] = useState<ApiUser>();
  const [title, setTitle] = useState("");
  const [date, setDate] = useState("");
  const [location, setLocation] = useState("");
  const [detail, setDetail] = useState("");
  const [err, setErr] = useState<Error>();

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
        <p className="my-8 flex justify-between">
          <a href={`/memory`}>戻る</a>
        </p>
        <div>
          <p>新規投稿</p>
          <p>{err?.message}</p>
          <div>
            <p>タイトル</p>
            <input
              type="text"
              value={title}
              onInput={(e: React.ChangeEvent<HTMLInputElement>) =>
                setTitle(e.target.value)
              }
            />
          </div>
          <div>
            <p>日付</p>
            <input
              type="date"
              value={date}
              onInput={(e: React.ChangeEvent<HTMLInputElement>) =>
                setDate(e.target.value)
              }
            />
          </div>
          <div>
            <p>場所</p>
            <input
              type="url"
              value={location}
              onInput={(e: React.ChangeEvent<HTMLInputElement>) =>
                setLocation(e.target.value)
              }
            />
          </div>
          <div>
            <p>詳細</p>
            <textarea
              value={detail}
              onChange={(e: React.ChangeEvent<HTMLTextAreaElement>) =>
                setDetail(e.target.value)
              }
            />
          </div>
        </div>
        <button
          onClick={async () => {
            const token = localStorage.getItem("jwt");
            if (!token) {
              localStorage.removeItem("jwt");
              route.push("/");
            } else {
              try {
                await fetch(API_SERVER + "/memory", {
                  method: "POST",
                  headers: {
                    "Content-Type": "application/json",
                    Authorization: token,
                  },
                  body: JSON.stringify({
                    title: title,
                    date: date,
                    location: location,
                    detail: detail,
                  }),
                });

                route.push("/memory");
              } catch (error) {
                if (error instanceof Error) {
                  setErr(error);
                }
              }
            }
          }}
        >
          投稿
        </button>
      </div>
    </main>
  );
}
