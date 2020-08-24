import React from 'react';
import { auth } from '../Firebase';
import Button from '@material-ui/core/Button';

const LogOut = () => {
    const handleLogout = () => {
        logout();
    }

    const logout = () => {
        try {
            auth.signOut();
        } catch (error) {
            console.log(`ログアウト時にエラーが発生しました (${error})`);
        }
    }

    return (
        <Button color="inherit" onClick={handleLogout}>LogOut</Button>
    );
}

export default LogOut;