import { useContext } from 'react'
import Login from './pages/Login'
import { UserContext } from './providers/UserProvider'
// import { ConfigContext } from './providers/ConfigProvider'

import {} from '../wailsjs/go/main/App'


const App = () => {
    const { username } = useContext(UserContext)
    // const { _rememberUser } = useContext(ConfigContext)

    // chat app

    // anyone can either host a dedicated server that others can join or can join someone's server (given the IP).

    // How to handle DMS? Either DMS are direct p2p connections or are per-server. Or both.

    // should be able to authenticate that who is sending a message is actually who they say they are.
    // when a user makes a connection (server or p2p) both parties agree on a passkey that must be sent by both parties before the message is actually sent.

    // should be able to create a local account that encrypts all data related to the account that requires a password to unlock

    return <div>{!username && <Login />}</div>
}

export default App
