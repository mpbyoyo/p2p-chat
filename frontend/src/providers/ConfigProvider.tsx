import { createContext, ReactNode, useContext, useState } from 'react'
import { ConfigType } from '@/common'

export const ConfigContext = createContext<ConfigType>({rememberUser: false})

const ConfigProvider = ({ children }: { children: ReactNode }) => {
    const config  = useContext(ConfigContext)
    const [configState, _setConfigState] = useState<ConfigType>(config)

    return <ConfigContext.Provider value={configState}>{children}</ConfigContext.Provider>
}

export default ConfigProvider
