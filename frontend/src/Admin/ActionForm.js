import '../App.css';

async function handleAction(e) {
    e.preventDefault();

    const form = e.target;
    const formData = new FormData(form);

    //mdtmp fetch('/some-api', { method: form.method, body: formData });

    const formJson = Object.fromEntries(formData.entries());
    console.log(formJson);
}

function ActionForm() {
return (
 <>
    <form  onSubmit={handleAction}>
      <div> Check box, POST or GET</div>
      <label>
        URL <input name="url" />
      </label>
      <hr />
      <label>
        Body: <textarea name="body" />
      </label>
      <hr />
      <button type="submit">Submit</button>
    </form>
 </>
);
}

export default ActionForm;

