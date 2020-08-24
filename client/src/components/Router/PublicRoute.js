import React, { useContext } from 'react';
import { Route, Redirect } from 'react-router-dom';
import { Store } from '../../store/index';

const PublicRoute = (props) => {
    const { globalState } = useContext(Store);

    return globalState.currentUser ?
        <Redirect to="/" /> : <Route {...props} />
}

export default PublicRoute;