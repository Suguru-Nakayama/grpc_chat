import React, { useEffect, useState } from 'react';
import { auth } from '../Firebase';
import { Redirect } from 'react-router-dom';

const Auth = ({ children }) => {
    const [isLogin, setIsLogin] = useState(false);
    const [Loading, setLoading] = useState(true);

    useEffect(() => {
        auth.onAuthStateChanged(user => {
            if (user) {
                setIsLogin(true);
            } else {
                setIsLogin(false);
            }
            setLoading(false);
        })
    }, [])

    return Loading ? < div > Loading</div > : (
        isLogin ? <div>{children}</div> : <Redirect to="/login" />
    );
}

export default Auth;