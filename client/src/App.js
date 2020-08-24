import React from 'react';
import { BrowserRouter, Switch } from 'react-router-dom';
import Top from './pages/Top';
import Signup from './pages/Signup';
import Login from './pages/Login';
import Auth from './components/Router/Auth';
import PrivateRoute from './components/Router/PrivateRoute';
import PublicRoute from './components/Router/PublicRoute';

function App() {
  return (
    <BrowserRouter>
      <Auth>
        <Switch>
          <PublicRoute path="/signup" component={Signup} />
          <PublicRoute path="/login" component={Login} />
          <PrivateRoute exact path="/" component={Top} />
        </Switch>
      </Auth>
    </BrowserRouter>
  );
}

export default App;
