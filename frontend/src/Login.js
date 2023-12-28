import './App.css';

async function testButton() {
/*
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
*/
  const response = await fetch("/api/login", {
    method: "get",
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    }
  })

  console.log(response);
  const text = await response.text();
  console.log(text);
  //const raw = response.headers.raw()['set-cookie'];
  //console.log(raw);

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
function Login() {
  return (
    <div className="App">
      <header className="App-header">
        <p>
          Business-Quebec
        </p>
{/* From documentation https://react.dev/reference/react-dom/components/input  */}
        <div>
          Username
        </div>
        <input></input>
        <div>
          Password
        </div>
        <input></input>
        <button onClick={testButton}>
          Test Login!!!!
        </button>
      </header>
    </div>
  );
}

export default Login;
