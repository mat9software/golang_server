import '../App.css';

async function handleAction(e) {
    e.preventDefault();

    const form = e.target;
    const formData = new FormData(form);

    const formJson = Object.fromEntries(formData.entries());

    try {
        let fetchData = {
          method: formJson.method,
          headers: {
            "Content-Type": "application/json",
          }
        }
        if(formJson.method === "POST")
          fetchData.body = JSON.stringify(formJson.body)

        const response = await fetch(formJson.url, fetchData)

        console.log("Call ", formJson.method, " ", formJson.url, " status:", response.status)

/*
        if(response.status !== 200) {
          console.error("Call failed", response.status)
        }
*/
        //const ret = await response.json()
        //console.log(ret)
    } catch (e) {
        console.error(e)
    }

}

function ActionForm() {
return (
 <>
    <form  onSubmit={handleAction}>
      {/* Caution, no on change here, re-render might overwrite values.  */}
      <label><input checked type="radio" name="method" value="POST"/>POST</label>
      <label><input type="radio" name="method" value="GET"/>GET</label>
      <hr />
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

