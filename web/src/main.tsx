import {BrowserRouter, Routes, Route} from "react-router";
import { createRoot } from 'react-dom/client'
import CssBaseline from '@mui/material/CssBaseline'
import {StyledEngineProvider, ThemeProvider} from '@mui/material/styles'
import './index.css'
import Home from './pages/Home.tsx'
import Login from './pages/Login.tsx'
import Dashboard from './pages/Dashboard.tsx'
import Sensors from './pages/Sensors.tsx'
import { Color } from './Theme/color.tsx'

createRoot(document.getElementById('root')!).render(
    <BrowserRouter>
        <StyledEngineProvider injectFirst>
            <ThemeProvider theme={ Color }>
                <CssBaseline />
                <Routes>
                    <Route path="Login" element={<Login />} />
                    <Route path="/" element={<Home />} >
                        <Route path="Dashboard" element={<Dashboard />} />
                        <Route path="Devices" element={<Dashboard />} >
                            <Route index element={<Home />} />
                            <Route path="Sensors" element={<Sensors />} />
                        </Route>
                        <Route path="Reports" element={<Home />} />
                    </Route>
                </Routes>
            </ThemeProvider>
        </StyledEngineProvider>
    </BrowserRouter>
)