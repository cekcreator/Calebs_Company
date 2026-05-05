"use client";

import { useEffect, useRef, useState } from "react";
import { createWebSocket } from "@/lib/api";

export function useWebSocket(path: string) {
  const ws = useRef<WebSocket | null>(null);
  const [messages, setMessages] = useState<string[]>([]);
  const [connected, setConnected] = useState(false);

  useEffect(() => {
    ws.current = createWebSocket(path);
    ws.current.onopen = () => setConnected(true);
    ws.current.onclose = () => setConnected(false);
    ws.current.onmessage = (e) => setMessages((prev) => [...prev, e.data]);
    return () => ws.current?.close();
  }, [path]);

  const send = (msg: string) => ws.current?.send(msg);

  return { messages, connected, send };
}
