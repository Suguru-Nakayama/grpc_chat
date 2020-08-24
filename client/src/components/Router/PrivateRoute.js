import React, { useContext } from 'react';
import { Route, Redirect } from 'react-router-dom';
import { Store } from '../../store/index';

const PrivateRoute = (props) => {
    const { globalState } = useContext(Store);

    return globalState.currentUser ?
        <Route {...props} /> : <Redirect to="/login" />;
}

export default PrivateRoute;