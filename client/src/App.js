import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Top from './pages/Top';
import Signup from './pages/Signup';
import Login from './pages/Login';
import Auth from './components/Auth';

function App() {
  return (
    <Router>
      <Switch>
        <Route exact path="/signup" component={Signup} />
        <Route exact path="/login" component={Login} />
        <Auth>
          <Switch>
            <Route exact path="/" component={Top} />
          </Switch>
        </Auth>
      </Switch>
    </Router>
  );
}

export default App;
