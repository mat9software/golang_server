import '../App.css';

import { useState } from "react"

async function handleSubmit(e, handleLogin) {
    e.preventDefault()

    const form = e.target
    const formData = new FormData(form)
    const formJson = Object.fromEntries(formData.entries())

    try {
        const fetchData = {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(formJson),
        }

        const response = await fetch('/api/login', fetchData)
        if(response.status === 200) {
          handleLogin(true)
          return
        }
        //const ret = await response.json()
        //console.log(ret)
    } catch (e) {
        console.error(e)
    }
    handleLogin(false)
}

function LoginForm(props) {
const [showFailedLogin, setShowFailedLogin] = useState(false)

const handleLogin = (success) => {
  props.handleLogin(success)// Propagate result
  setShowFailedLogin(!success)// Display error if necessary
}

return (
 <>
    <form onSubmit={(e) => handleSubmit(e, handleLogin)}>
      <label>
        Username: <input name="username" />
      </label>
      <hr />
      <label>
        Password: <input type="password" name="password" />
      </label>
      <hr />
      { showFailedLogin &&
        <div style={{ color:"red" }} >Login Failed.</div>
      }
      <button type="submit">Submit</button>
    </form>
 </>
);
}

export default LoginForm;

