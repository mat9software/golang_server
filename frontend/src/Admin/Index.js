import '../App.css';

import ActionForm from './ActionForm';
import LoginForm from './LoginForm';

import { useState } from 'react';

//--------------------------------------------------

const E_LOADING = 0;
const E_LOGIN = 1;
const E_ACTION = 2;

function Loading() {
  return (
    <div>Loading</div>
  )
}

function Admin() {
  //const [pageStatus, setPageStatus] = useState(E_LOADING);
  //const [pageStatus, setPageStatus] = useState(E_LOGIN);
  const [pageStatus, setPageStatus] = useState(E_ACTION);

  return (
    <div className="App">
      <header className="App-header">
        <p>
          Business-Quebec
        </p>

        {pageStatus === E_LOADING &&
         <Loading/>
        }
        {pageStatus === E_LOGIN &&
         <LoginForm/>
        }
        {pageStatus === E_ACTION &&
         <ActionForm/>
        }

      </header>
    </div>
  );
}

export default Admin;

