import { createContext, ReactNode, useContext, useState } from 'react'
import { UserType } from '../common/types'

export const UserContext = createContext<UserType>({
    username: null,
})

const UserProvider = ({ children }: { children: ReactNode }) => {
    const user = useContext(UserContext)
    const [userState, _setUserState] = useState<UserType>(user)

    return <UserContext.Provider value={userState}>{children}</UserContext.Provider>
}

export default UserProvider
