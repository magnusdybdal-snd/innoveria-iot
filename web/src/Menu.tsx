import innLogo from './assets/innoveria.png'
import viteLogo from '/vite.svg'
import Box from '@mui/material/Box'
import Button from '@mui/material/Button'
import Divider from '@mui/material/Divider'
import Typography from '@mui/material/Typography'
import Avatar from '@mui/material/Avatar'
import { Link as RouterLink } from 'react-router'
import MenuBox from './templates/menuBox.tsx'
import SubPages from './pages/subPageList.tsx'
import MainPages from './pages/mainPageList.tsx'

export default function Menu() {
  return (
      <Box
          sx={{
              backgroundColor: 'primary.light',
              color: 'primary.main',
          }}
      >
          <div className="w-64 bg-gray-900 text-white flex flex-col h-screen">
              {/* Top menu logo */}
              <a href="/">
                  <img
                      src={innLogo}
                      alt="Innoveria logo"
                      style={{
                          width: '310px',
                          height: 'auto'
                      }}
                  />
              </a>
              {/* User info */}
              <div className="flex items-center gap-3 p-4">
                  <Avatar
                      alt="User"
                      src={viteLogo}
                      style={{
                          width: '60px',
                          height: 'auto'
                      }}
                  />
                  <div>
                      <Typography variant="h3">
                          Username
                      </Typography>
                      <Typography variant="h6">
                          email@email.com
                      </Typography>
                  </div>
                  <Divider
                      sx={{
                          backgroundColor: 'primary.main',
                      }}
                      variant="middle"
                  />
              </div>
              {/* Menu navigation */}
              <nav className="flex-1 flex flex-col mt-4 space-y-2">
                  {Array.from(MainPages.entries()).map(([category, page]) => (
                      <MenuBox
                          key={category}
                          title={category}
                          mainPage={page}
                          subPages={SubPages.get(category) ?? []}
                          add={category !== "devices"} // example logic
                      />
                  ))}
              </nav>

              {/* Logout pinned to bottom */}
              <div className="p-4">
                  <Button
                      component={RouterLink}
                      to="/Login"
                      variant="outlined"
                      sx={{
                          backgroundColor: 'secondary.main',
                          color: 'white',
                          '&:hover': {
                              backgroundColor: 'secondary.dark',
                          },
                          borderRadius: 2,
                          margin: 2
                      }}
                  >
                      Log out
                  </Button>
              </div>
          </div>
      </Box>
  )
}
