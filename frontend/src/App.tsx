import { Button } from '@/components/ui/button'
import { Greet } from '../wailsjs/go/main/App'

import React from 'react'

function App() {
    const [count, setCount] = React.useState(0)
    const [greet, setGreet] = React.useState('')

    return (
        <div className="min-h-screen bg-white grid place-items-center mx-auto py-8">
            <div className="text-blue-900 text-2xl font-bold flex flex-col items-center space-y-4">
                <h1>Vite + React + TS + Tailwind + shadcn/ui</h1>
                <Button onClick={() => setCount(count + 1)}>Count up ({count})</Button>
                <Button onClick={() => Greet('John').then(setGreet)}>Greet</Button>
                <p>{greet}</p>
                {greet && count > 0 && (
                    <Button
                        onClick={() => {
                            setGreet('')
                            setCount(0)
                        }}
                    >
                        Clear
                    </Button>
                )}
            </div>
        </div>
    )
}

export default App
