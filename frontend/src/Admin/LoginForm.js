import '../App.css';

async function handleLogin(e) {
    e.preventDefault();

    const form = e.target;
    const formData = new FormData(form);

    //mdtmp fetch('/some-api', { method: form.method, body: formData });

    const formJson = Object.fromEntries(formData.entries());
    console.log(formJson);
}

function LoginForm() {
return (
 <>
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

export default LoginForm;

