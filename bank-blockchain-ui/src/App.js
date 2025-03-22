import React, { useEffect, useState } from "react";

function App() {
    const [blocks, setBlocks] = useState([]);

    // 定義 fetchBlocks() 來獲取區塊鏈數據
    const fetchBlocks = async () => {
        try {
            const response = await fetch("http://localhost:8080/blocks");
            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            const data = await response.json();
            setBlocks(data);
        } catch (error) {
            console.error("Fetch error:", error);
        }
    };

    useEffect(() => {
        fetchBlocks();
    }, []);

    const depositMoney = async () => {
        try {
            const response = await fetch("http://localhost:8080/deposit", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ data: "User deposits $200" }),
            });

            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }

            await fetchBlocks(); // **這裡確保 fetchBlocks() 存在**
        } catch (error) {
            console.error("Deposit failed:", error);
        }
    };

    return (
        <div>
            <h1>Bank Blockchain</h1>
            <button onClick={depositMoney}>Deposit $200</button>
            <ul>
                {blocks.map((block) => (
                    <li key={block.Index}>
                        <strong>Block #{block.Index}</strong>: {block.Data}
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default App;
