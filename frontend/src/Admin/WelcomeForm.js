import '../App.css';

import { useEffect } from 'react';

async function validateLoginStatus() {
    try {
        const fetchData = {
          method: "GET",
        }

        const response = await fetch('/api/welcome', fetchData)
        if (response.status === 200) 
          return true;
    } catch (e) {
        console.error(e)
    }
    return false;
}

function WelcomeForm(props) {

useEffect(() => {
   (async () => {
     const ret = await validateLoginStatus()
     props.handleWelcome(ret)
   })()
});

return (
 <>
    <div>Loading</div>
 </>
);
}

export default WelcomeForm;

