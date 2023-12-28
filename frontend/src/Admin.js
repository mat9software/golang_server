import './App.css';
import { useState } from 'react';

async function handleLogin(e) {
    e.preventDefault();

    const form = e.target;
    const formData = new FormData(form);

    //mdtmp fetch('/some-api', { method: form.method, body: formData });

    const formJson = Object.fromEntries(formData.entries());
    console.log(formJson);
}
async function ActionCb() {
}

function Actions() {
return (
 <>
        <div> Check box, POST or GET</div>
        <div>URL</div>
        <input></input>
        <div>Body</div>
        <textarea>
        </textarea>
        <button onClick={ActionCb}>
          Submit
        </button>
 </>
);
}

function LoginForm() {
return (
 <>
{/* From documentation https://react.dev/reference/react-dom/components/input  */}
    <form  onSubmit={handleLogin}>
      <label>
        Username: <input name="username" />
      </label>
      <hr />
      <label>
        Password: <input name="password" />
      </label>
      <hr />
      <button type="submit">Submit</button>
    </form>
 </>
);
}

function Admin() {
  const [isLoggedIn, setLoggedIn] = useState(false);

  return (
    <div className="App">
      <header className="App-header">
        <p>
          Business-Quebec
        </p>
        {isLoggedIn
        ? <Actions/>
        : <LoginForm/>
        }
      </header>
    </div>
  );
}

export default Admin;

