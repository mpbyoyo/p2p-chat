import { useContext, useEffect } from 'react'
import Login from './pages/Login'
import { UserContext } from './providers/UserProvider'
import { ConfigContext } from './providers/ConfigProvider'

import { RequestDir } from '../wailsjs/go/main/FileManager'

const App = () => {
    const { username } = useContext(UserContext)
    const { rememberUser: _rememberUser } = useContext(ConfigContext)

    // frontend does not ever need to know the directory. All dir ops will be handled by GO. This is just for testing
    const getDir = async () => {
        const dir = await RequestDir()
        console.log(dir)
    }

    useEffect(() => {
        getDir()
    }, [])

    // chat app

    // anyone can either host a dedicated server that others can join or can join someone's server (given the IP).

    // How to handle DMs? Either DMs are direct p2p connections or are per-server. Or both.
    // Should be both. "Known" is someone you have friended within a server. "Contact" is someone you've shared a key directly with (can only establish a p2p connection with a "Contact").

    // should be able to authenticate that who is sending a message is actually who they say they are.
    // when a user makes a connection (server or p2p) both parties agree on a passkey that must be sent by both parties before the message is actually sent.

    // should be able to create a local account that encrypts all data related to the account that requires a password to unlock

    return (
        <div>
            {!username && <Login />}
        </div>
    )
}

export default App
