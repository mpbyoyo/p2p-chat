import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import UserProvider from './providers/UserProvider.tsx'
import ConfigProvider from './providers/ConfigProvider.tsx'

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <React.StrictMode>
        <ConfigProvider>
            <UserProvider>
                <App />
            </UserProvider>
        </ConfigProvider>
    </React.StrictMode>
)
