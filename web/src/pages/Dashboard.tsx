import Box from '@mui/material/Box'
import AddBox from '../templates/addCompBox.tsx'
import Path from '../templates/path.tsx'
import Menu from "../Menu.tsx"

export default function Home() {
    return (
        <div className="flex h-screen">
            <Menu/>
            <Box
                sx={{
                    backgroundColor: 'primary.dark',
                    color: 'primary.main',
                }}
                className="flex-1 overflow-auto"
            >
                <Box
                    sx={{
                        paddingLeft: 2,
                        paddingRight: 2,
                        paddingTop: 2
                    }}
                    className="flex-1 overflow-auto"
                >
                    <Path/>
                </Box>
                <AddBox/>
            </Box>
        </div>
    )
}