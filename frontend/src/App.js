import logo from './logo.svg';
import './App.css';

async function testButton() {
  const response = await fetch("/backend/signin", {
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

  console.log(response);
  const text = await response.text();
  console.log(text);
  //const raw = response.headers.raw()['set-cookie'];
  //console.log(raw);
  const raw = response.headers
  console.log(raw);

/*
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
      console.log(response.text)
  });
*/
}

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <button onClick={testButton}>
          Login!!!!
        </button>
      </header>
    </div>
  );
}

export default App;
