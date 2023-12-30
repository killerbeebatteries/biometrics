import { useEffect, useState } from "react";
import "./App.css";

// TODO:
// - Make it easier to add and modify records
// - Add a graph of data
// - Add support for comment annotations

function App() {
  const [biometrics, setBPAndWeight] = useState(false);

  function getBPAndWeight() {
    fetch("http://localhost:3001")
      .then((response) => {
        return response.text();
      })
      .then((data) => {
        setBPAndWeight(data);
      });
  }

  function createBPAndWeight() {
    let date = prompt("Enter date in 20240124 format.");
    let time = prompt("Enter time in 24 hour format.");
    let sys  = prompt("Enter systolic pressure.");
    let dia  = prompt("Enter diasystolic pressure.");
    let bp   = prompt("Enter blood pressure.");
    let weight_total  = prompt("Enter weight in KG.");
    let weight_fat    = prompt("Enter body fat in KG.");
    let weight_muscle = prompt("Enter muscle in KG.");
    let comment       = prompt("Enter a comment.");
    fetch("http://localhost:3001/biometrics", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ date, time, sys, dia, bp, weight_total, weight_fat, weight_muscle, comment }),
    })
      .then((response) => {
        return response.text();
      })
      .then((data) => {
        alert(data);
        getBPAndWeight();
      });
  }

  function deleteBPAndWeight() {
    let id = prompt("Enter biometrics id");
    fetch(`http://localhost:3001/biometrics/${id}`, {
      method: "DELETE",
    })
      .then((response) => {
        return response.text();
      })
      .then((data) => {
        alert(data);
        getBPAndWeight();
      });
  }

  function updateBPAndWeight() {
    let date = prompt("Enter date in 20240124 format.");
    let time = prompt("Enter time in 24 hour format.");
    let sys  = prompt("Enter systolic pressure.");
    let dia  = prompt("Enter diasystolic pressure.");
    let bp   = prompt("Enter blood pressure.");
    let weight_total  = prompt("Enter weight in KG.");
    let weight_fat    = prompt("Enter body fat in KG.");
    let weight_muscle = prompt("Enter muscle in KG.");
    let comment       = prompt("Enter a comment.");
    fetch(`http://localhost:3001/biometrics/${id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ date, time, sys, dia, bp, weight_total, weight_fat, weight_muscle, comment}),
    })
      .then((response) => {
        return response.text();
      })
      .then((data) => {
        alert(data);
        getBPAndWeight();
      });
  }

  useEffect(() => {
    getBPAndWeight();
  }, []);
  return (
    <div>
      {biometrics ? biometrics : "There is no biometrics data available"}
      <br />
      <button onClick={createBPAndWeight}>Add biometrics</button>
      <br />
      <button onClick={deleteBPAndWeight}>Delete biometrics</button>
      <br />
      <button onClick={updateBPAndWeight}>Update biometrics</button>
    </div>
  );
}
export default App;
