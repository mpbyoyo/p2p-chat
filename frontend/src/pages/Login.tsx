import { UserType } from '@/common'
import { useEffect, useState } from 'react'
import { ReadUsers } from '../../wailsjs/go/main/App'

const Login = () => {
    const [_availableUsers, setAvailableUsers] = useState<UserType[]>()
    
    useEffect(() => {
        ReadUsers()
            .then((users) => {
                if (Array.isArray(users)) setAvailableUsers(users)
            })
            .catch((e) => console.error(e))
    }, [])

    return (
        <div className="min-h-screen bg-purple-300">
            {/* determine if there are any created users. If so display them all as available logins. If not display a button to create an account */}
            <div></div>
        </div>
    )
}

export default Login
