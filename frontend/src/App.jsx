import { useState, useEffect } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

// TODO:
// - Update the update function
// - Add support to add metrics
// - Add a graph of data
// - Add support for comment annotations

function App() {
  const [biometrics, setBPAndWeight] = useState(false);

  function getBPAndWeight() {
    fetch('http://localhost:3001')
      .then(response => {
        return response.text();
      })
      .then(data => {
        setBPAndWeight(data);
      });
  }

  function createBPAndWeight() {
    let name = prompt('Enter biometrics name');
    let email = prompt('Enter biometrics email');
    fetch('http://localhost:3001/biometrics', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({name, email}),
    })
      .then(response => {
        return response.text();
      })
      .then(data => {
        alert(data);
        getBPAndWeight();
      });
  }

  function deleteBPAndWeight() {
    let id = prompt('Enter biometrics id');
    fetch(`http://localhost:3001/biometrics/${id}`, {
      method: 'DELETE',
    })
      .then(response => {
        return response.text();
      })
      .then(data => {
        alert(data);
        getBPAndWeight();
      });
  }

  function updateBPAndWeight() {
    let id = prompt('Enter biometrics id');
    let name = prompt('Enter new biometrics name');
    let email = prompt('Enter new biometrics email');
    fetch(`http://localhost:3001/biometrics/${id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({name, email}),
    })
      .then(response => {
        return response.text();
      })
      .then(data => {
        alert(data);
        getBPAndWeight();
      });
  }

  useEffect(() => {
    getBPAndWeight();
  }, []);
  return (
    <div>
      {biometrics ? biometrics : 'There is no biometrics data available'}
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
