import React, { createContext, useReducer } from 'react';

const initialState = {
    currentUser: null
}

const reducer = (state, action) => {
    switch (action.type) {
        case 'SET_CURRENT_USER':
            return { ...state, currentUser: action.payload.currentUser };
        default:
            return state;
    }
}

export const Store = createContext({
    globalState: initialState,
    setGlobalState: () => null
});

const StoreProvider = ({ children }) => {
    const [globalState, setGlobalState] = useReducer(reducer, initialState);
    return (
        <Store.Provider value={{ globalState, setGlobalState }}>
            {children}
        </Store.Provider>
    );
}

export default StoreProvider;