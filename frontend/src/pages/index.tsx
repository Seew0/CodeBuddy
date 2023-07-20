import { useState, ChangeEvent, KeyboardEvent } from "react";
import axios from "axios";
import Head from "next/head";

interface Message {
  text: string;
  sender: "You" | "CodeBuddy";
}

interface ApiResponse {
  answer: string;
  status: string;
}

const Home: React.FC = () => {
  const [messages, setMessages] = useState<Message[]>([
    { text: "Sup I am CodeBuddy. Ask me your question", sender: "CodeBuddy" },
  ]);
  const [inputValue, setInputValue] = useState<string>("");

  const handleInputChange = (event: ChangeEvent<HTMLInputElement>) => {
    setInputValue(event.target.value);
  };

  const handleUserMessage = async () => {
    if (inputValue.trim() === "") return;

    const userMessage: Message = { text: inputValue, sender: "You" };

    setMessages([...messages, userMessage]);
    setInputValue("");

    try {
      const response = await axios.post<ApiResponse>(
        "http://localhost:5000/api/answer",
        {
          query: inputValue, 
        }
      );

      const botReply: Message = { text: response.data.answer, sender: "CodeBuddy" };

      setMessages([...messages, userMessage, botReply]);
    } catch (error) {
      console.error("Error fetching AI response:", error);
    }
  };

  const handleKeyPress = (event: KeyboardEvent<HTMLInputElement>) => {
    if (event.key === "Enter") {
      handleUserMessage();
    }
  };

  return (
    <>
      <Head>
        <title>CodeBuddy</title>
      </Head>
      <div className="flex flex-col items-center justify-center h-screen bg-gray-600">
        <h1 className=" text-5xl text-slate-300 p-3">CodeBuddy</h1>
        <div className=" bg-slate-800 p-4 w-screen h-screen overflow-y-auto">
          <div className="chat-messages space-y-4">
            {messages.map((message, index) => (
              <div
                key={index}
                className={`message p-2 whitespace-pre-wrap text-slate-300 ${
                  message.sender === "You"
                    ? "bg-blue-800 self-end"
                    : "bg-gray-700"
                }`}
              >
                {message.sender + ": " + message.text}
              </div>
            ))}
          </div>
        </div>
        <div className="You-input mt-4 max-w-screen-2xl">
          <div className="flex">
            <input
              type="text"
              value={inputValue}
              onChange={handleInputChange}
              onKeyPress={handleKeyPress}
              placeholder="Ask your question"
              className="border border-gray-400 bg-gray-700 p-3 flex-grow w-screen mb-2"
            />
            <button
              onClick={handleUserMessage}
              className=" bg-purple-600 p-3 ml-2 mb-2"
            >
              Ask
            </button>
          </div>
        </div>
      </div>
    </>
  );
};

export default Home;
