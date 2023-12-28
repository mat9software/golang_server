import '../App.css';

import ActionForm from './ActionForm';
import LoginForm from './LoginForm';
import WelcomeForm from './WelcomeForm';

import { useState } from 'react';

//--------------------------------------------------

const E_WELCOME = 0;
const E_LOGIN = 1;
const E_ACTION = 2;

function Admin() {
  const [pageStatus, setPageStatus] = useState(E_WELCOME);

  const handleWelcome = (success) => {
    if(success)
      setPageStatus(E_ACTION)
    else
      setPageStatus(E_LOGIN)
  }
  const handleLogin = (success) => {
    if(success)
      setPageStatus(E_ACTION)
  }

  return (
    <div className="App">
      <header className="App-header">
        <p>
          Business-Quebec
        </p>

        {pageStatus === E_WELCOME &&
         <WelcomeForm handleWelcome={handleWelcome}/>
        }
        {pageStatus === E_LOGIN &&
         <LoginForm handleLogin={handleLogin}/>
        }
        {pageStatus === E_ACTION &&
         <ActionForm/>
        }

      </header>
    </div>
  );
}

export default Admin;

