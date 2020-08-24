import React, { useEffect, useState, useContext } from 'react';
import { auth } from '../../Firebase';
import { Store } from '../../store/index';

const Auth = ({ children }) => {
    const [Loading, setLoading] = useState(true);
    const { setGlobalState } = useContext(Store);

    useEffect(() => {
        auth.onAuthStateChanged(user => {
            setGlobalState({
                type: 'SET_CURRENT_USER',
                payload: { currentUser: user }
            });
            setLoading(false);
        });
    }, [setGlobalState, Loading])

    if (Loading) return (<div> Loading</div>);

    return <div>{children}</div>;
}

export default Auth;