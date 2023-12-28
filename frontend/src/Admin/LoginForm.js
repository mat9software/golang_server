import '../App.css';

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
return (
 <>
    <form onSubmit={(e) => handleSubmit(e, props.handleLogin)}>
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

export default LoginForm;

