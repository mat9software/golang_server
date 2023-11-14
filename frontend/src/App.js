import logo from './logo.svg';
import './App.css';

function testButton() {
  fetch("/backend/signin", {
    method: "post",
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    },

    body: JSON.stringify({
      username: "user1",
      password: "password1"
    })
  })
  .then( (response) => { 
      console.log(response)
  });
}

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
        <button onClick={testButton}>
          Login!!!!
        </button>
      </header>
    </div>
  );
}

export default App;
