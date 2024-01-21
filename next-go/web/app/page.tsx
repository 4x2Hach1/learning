"use client";

import { API_SERVER } from "@/utils/const";
import { ApiLoginSchema } from "@/utils/types";
import { useRouter, redirect } from "next/navigation";
import React, { useEffect, useState } from "react";

export default function Home() {
  const [islogin, setIslogin] = useState<boolean>(true);
  const [name, setName] = useState<string>("");
  const [email, setEmail] = useState<string>("user@example.com");
  const [password, setPassword] = useState<string>("password");
  const [err, setErr] = useState<Error>();
  const [toTop, setToTop] = useState(false);
  const router = useRouter();

  useEffect(() => {
    if (toTop) {
      redirect("/memory");
    }
  }, [toTop]);

  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      {(() => {
        if (islogin) {
          return (
            <div>
              <p className="mb-8 ">ログイン</p>
              <p>{err?.message}</p>
              <div>
                <p>メールアドレス</p>
                <input
                  type="email"
                  value={email}
                  onInput={(e: React.ChangeEvent<HTMLInputElement>) =>
                    setEmail(e.target.value)
                  }
                />
              </div>
              <div>
                <p>パスワード</p>
                <input
                  type="password"
                  value={password}
                  onInput={(e: React.ChangeEvent<HTMLInputElement>) =>
                    setPassword(e.target.value)
                  }
                />
              </div>
              <p className="mb-8">
                <button
                  onClick={async () => {
                    try {
                      const res = await fetch(API_SERVER + "/login", {
                        method: "POST",
                        headers: { "Content-Type": "application/json" },
                        body: JSON.stringify({
                          email: email,
                          password: password,
                        }),
                      });

                      const body = await res.json();
                      const data = ApiLoginSchema.parse(body);
                      localStorage.setItem("jwt", data);
                      setToTop(true);
                    } catch (error) {
                      if (error instanceof Error) {
                        setErr(error);
                      }
                    }
                  }}
                >
                  ログイン
                </button>
              </p>
              <p>
                <button
                  onClick={() => {
                    setEmail("");
                    setPassword("");
                    setErr(undefined);
                    setIslogin(false);
                  }}
                >
                  新規作成
                </button>
              </p>
            </div>
          );
        } else {
          return (
            <div>
              <p className="mb-8">新規作成</p>
              <p>{err?.message}</p>
              <div>
                <p>名前</p>
                <input
                  type="text"
                  value={name}
                  onInput={(e: React.ChangeEvent<HTMLInputElement>) =>
                    setName(e.target.value)
                  }
                />
              </div>
              <div>
                <p>メールアドレス</p>
                <input
                  type="email"
                  value={email}
                  onInput={(e: React.ChangeEvent<HTMLInputElement>) =>
                    setEmail(e.target.value)
                  }
                />
              </div>
              <div>
                <p>パスワード</p>
                <input
                  type="password"
                  value={password}
                  onInput={(e: React.ChangeEvent<HTMLInputElement>) =>
                    setPassword(e.target.value)
                  }
                />
              </div>
              <p className="mb-8">
                <button
                  onClick={async () => {
                    try {
                      await fetch(API_SERVER + "/user", {
                        method: "POST",
                        headers: { "Content-Type": "application/json" },
                        body: JSON.stringify({
                          name: name,
                          email: email,
                          password: password,
                        }),
                      });

                      router.refresh();
                      setName("");
                      setEmail("");
                      setPassword("");
                      setErr(undefined);
                      setIslogin(true);
                    } catch (error) {
                      if (error instanceof Error) {
                        setErr(error);
                      }
                    }
                  }}
                >
                  作成
                </button>
              </p>
              <p>
                <button
                  onClick={() => {
                    setName("");
                    setEmail("");
                    setPassword("");
                    setErr(undefined);
                    setIslogin(true);
                  }}
                >
                  ログイン
                </button>
              </p>
            </div>
          );
        }
      })()}
    </main>
  );
}
