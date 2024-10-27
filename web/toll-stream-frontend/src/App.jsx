import { useState } from "react";
import "./App.css";

function TollCalculator() {
  const [inputText, setInputText] = useState("");
  const [date, setdate] = useState("");
  const [isAPIDataFetched, setIsAPIDataFetched] = useState(false);
  const [toll, setToll] = useState(0);

  const handleShowToll = async () => {
    const formattedDate = date.split("-").reverse().join("-");

    try {
      const response = await fetch(
        `http://localhost:3001/get-invoice?obu=${inputText}&date=${formattedDate}`
      );
      if (!response.ok) {
        throw new Error("Failed to fetch data");
      }

      const data = await response.json();
      setToll(data.totalAmount);
      setIsAPIDataFetched(true);
    } catch (error) {
      console.error("Error fetching toll data:", error);
      setIsAPIDataFetched(false);
      alert('Error Occurred. Please enter correct data.')
    }
  };

  return (
    <div style={{ maxWidth: "400px" }}>
      <h2>Check Your Toll</h2>
      <div style={{ marginBottom: "15px" }}>
        <label>Enter OBU ID:</label>
        <input
          type="text"
          value={inputText}
          onChange={(e) => setInputText(e.target.value)}
          placeholder="Ex: 6f5ebyu5d32wevrytgf53e"
          style={{ width: "100%", padding: "8px", marginTop: "5px" }}
        />
      </div>

      <div style={{ marginBottom: "15px" }}>
        <label>Date:</label>
        <input
          type="date"
          value={date}
          onChange={(e) => setdate(e.target.value)}
          style={{ width: "100%", padding: "8px", marginTop: "5px" }}
        />
      </div>

      <div
        style={{
          display: "flex",
          justifyContent: "space-between",
          width: "100%",
        }}
      >
        <button
          onClick={handleShowToll}
          style={{
            padding: "1rem",
            backgroundColor: "#4CAF50",
            color: "white",
            border: "none",
            cursor: "pointer",
            height: "2.5rem",
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
          }}
        >
          Show Toll
        </button>

        {isAPIDataFetched ? (
          <h3 style={{ marginTop: "0.5rem" }}>₹ {toll}</h3>
        ) : (
          <></>
        )}
        
      </div>
    </div>
  );
}

export default TollCalculator;
