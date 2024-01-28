"use client";

import { API_SERVER } from "@/utils/const";
import { ApiHeavyCheckSchema, ApiHeavySchema } from "@/utils/types";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";

export default function Heavy() {
  const router = useRouter();
  const [modal, setModal] = useState(false);
  const [isRunning, setIsRunning] = useState(false);
  const [status, setStatus] = useState(0);
  const [err, setErr] = useState<Error>();
  const [key, setKey] = useState<string>("");
  const [check, setCheck] = useState<NodeJS.Timeout>();

  useEffect(() => {
    if (100 <= status) {
      clearInterval(check);
      setCheck(undefined);
      setIsRunning(false);
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [status]);

  return (
    <>
      <p className="my-8">
        <button onClick={() => setModal(true)}>heavy</button>
      </p>
      {modal && (
        <div className="fixed inset-0 w-screen h-screen bg-gray-800 bg-opacity-25 flex justify-center items-center">
          <div className="bg-slate-100 w-80 h-40 flex flex-col">
            <div className="text-center	grow py-8">
              <p>run heavy{isRunning && " running"}</p>
              {isRunning && (
                <p>
                  {status}%
                  <div className="w-full bg-neutral-200">
                    <div
                      className="bg-blue-600 p-0.5 text-center text-xs font-medium leading-none"
                      style={{ width: `${status}%` }}
                    ></div>
                  </div>
                </p>
              )}
              <p>{err?.message}</p>
            </div>
            {status < 100 ? (
              <div className="flex justify-between px-4 py-4">
                <button onClick={() => setModal(false)}>キャンセル</button>
                {!isRunning && (
                  <button
                    onClick={async () => {
                      setIsRunning(true);
                      const token = localStorage.getItem("jwt");
                      if (!token) {
                        localStorage.removeItem("jwt");
                        router.push("/");
                      } else {
                        try {
                          // run heavy
                          const res = await fetch(API_SERVER + "/heavy", {
                            method: "POST",
                            headers: {
                              "Content-Type": "application/json",
                              Authorization: token,
                            },
                          });
                          const body = await res.json();
                          const newKey = ApiHeavySchema.parse(body);
                          setKey(newKey);

                          // set heavy checker
                          setCheck(
                            setInterval(async () => {
                              try {
                                const res = await fetch(
                                  API_SERVER + `/heavy/${newKey}`,
                                  {
                                    method: "GET",
                                    headers: {
                                      "Content-Type": "application/json",
                                      Authorization: token,
                                    },
                                  }
                                );
                                const body = await res.json();
                                const newStatus =
                                  ApiHeavyCheckSchema.parse(body);
                                setStatus(newStatus);
                              } catch (error) {
                                if (error instanceof Error) {
                                  setErr(error);
                                }
                              }
                            }, 3 * 1000)
                          );
                        } catch (error) {
                          if (error instanceof Error) {
                            setErr(error);
                          }
                        }
                      }
                    }}
                  >
                    実行
                  </button>
                )}
              </div>
            ) : (
              <div className="flex justify-between px-4 py-4">
                <button
                  onClick={async () => {
                    const token = localStorage.getItem("jwt");
                    if (!token) {
                      localStorage.removeItem("jwt");
                      router.push("/");
                    } else {
                      try {
                        const res = await fetch(API_SERVER + `/heavy/${key}`, {
                          method: "DELETE",
                          headers: {
                            "Content-Type": "application/json",
                            Authorization: token,
                          },
                        });
                        if (res.ok) {
                          setStatus(0);
                          setErr(undefined);
                          setKey("");
                          setModal(false);
                        } else {
                          const body = await res.json();
                          throw new Error(body.message);
                        }
                      } catch (error) {
                        if (error instanceof Error) {
                          setErr(error);
                        }
                      }
                    }
                  }}
                >
                  閉じる
                </button>
              </div>
            )}
          </div>
        </div>
      )}
    </>
  );
}
