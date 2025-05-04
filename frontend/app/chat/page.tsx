"use client";

import { useEffect, useState, useRef } from "react";

interface ChatMessage {
  type: string;
  sender_id: string;
  recipient_id?: string;
  content: string;
  attachment_url?: string;
  timestamp: string;
}

const ChatPage = () => {
  const [messages, setMessages] = useState<ChatMessage[]>([]);
  const [input, setInput] = useState("");
  const [wsConnected, setWsConnected] = useState(false);
  const wsRef = useRef<WebSocket | null>(null);
  // Replace with your actual user id, fetched from session or context
  const userID = "user1";

  useEffect(() => {
    const ws = new WebSocket(`ws://localhost:5000/ws/chat?user_id=${userID}`);
    wsRef.current = ws;

    ws.onopen = () => {
      setWsConnected(true);
      console.log("WebSocket connected");
    };

    ws.onmessage = (event) => {
      const msg = JSON.parse(event.data) as ChatMessage;
      setMessages(prev => [...prev, msg]);
    };

    ws.onclose = () => {
      setWsConnected(false);
      console.log("WebSocket disconnected");
    };

    return () => {
      ws.close();
    };
  }, [userID]);

  const sendMessage = () => {
    if (wsRef.current && wsConnected) {
      const message: ChatMessage = {
        type: "broadcast", // or "private" with a recipient_id
        sender_id: userID,
        content: input,
        timestamp: new Date().toISOString(),
      };
      wsRef.current.send(JSON.stringify(message));
      setInput("");
    }
  };

  return (
    <div style={{ padding: "1rem" }}>
      <h1>Chat</h1>
      <div style={{ border: "1px solid #ccc", padding: "1rem", height: "300px", overflowY: "scroll" }}>
        {messages.map((msg, index) => (
          <div key={index}>
            <strong>{msg.sender_id}</strong>: {msg.content} <small>({msg.timestamp})</small>
          </div>
        ))}
      </div>
      <div style={{ marginTop: "1rem" }}>
        <input
          type="text"
          value={input}
          onChange={(e) => setInput(e.target.value)}
          placeholder="Type your message"
          style={{ width: "80%" }}
        />
        <button onClick={sendMessage} style={{ width: "18%" }}>Send</button>
      </div>
    </div>
  );
};

export default ChatPage;
