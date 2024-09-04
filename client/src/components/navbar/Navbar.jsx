import { useState, useContext } from 'react';
import { AppBar, Box, Toolbar, IconButton, Typography, MenuItem, Menu, Drawer, Divider, Tooltip, Button } from '@mui/material';
import MenuIcon from '@mui/icons-material/Menu';
import LoginIcon from '@mui/icons-material/Login';
import LogoutIcon from '@mui/icons-material/Logout';
import ArrowCircleLeftIcon from '@mui/icons-material/ArrowCircleLeft';
import AccountCircle from '@mui/icons-material/AccountCircle';
import HomeIcon from '@mui/icons-material/Home';
import MovieIcon from '@mui/icons-material/Movie';
import LightModeIcon from '@mui/icons-material/LightMode';
import DarkModeIcon from '@mui/icons-material/DarkMode';
import { Link } from 'react-router-dom';
import PropTypes from 'prop-types';
import { AuthContext } from '../../hooks/AuthContext';

const Navbar = ({ toggleTheme, isDarkMode }) => {
    const { isAuthenticated, logout } = useContext(AuthContext);
    const [anchorEl, setAnchorEl] = useState(null);
    const [drawerOpen, setDrawerOpen] = useState(false);

    const isMenuOpen = Boolean(anchorEl);

    const handleMenuClose = () => {
        setAnchorEl(null);
    };

    const toggleDrawer = (open) => (event) => {
        if (event.type === 'keydown' && (event.key === 'Tab' || event.key === 'Shift')) {
            return;
        }
        setDrawerOpen(open);
    };

    const menuId = 'primary-search-account-menu';
    const renderMenu = (
        <Menu
            anchorEl={anchorEl}
            anchorOrigin={{
                vertical: 'top',
                horizontal: 'right',
            }}
            id={menuId}
            keepMounted
            transformOrigin={{
                vertical: 'top',
                horizontal: 'right',
            }}
            open={isMenuOpen}
            onClose={handleMenuClose}
        >
            <MenuItem onClick={handleMenuClose}>Profile</MenuItem>
            <MenuItem onClick={handleMenuClose}>My account</MenuItem>
        </Menu>
    );

    const drawerList = (
        <Box
            sx={{ width: 250, display: 'flex', flexDirection: 'column', height: '100%' }}
            role="presentation"
            onClick={toggleDrawer(false)}
            onKeyDown={toggleDrawer(false)}
        >
            <Box
                sx={{
                    display: 'flex',
                    justifyContent: 'space-between',
                    alignItems: 'center',
                    p: 2
                }}
            >
                <Tooltip title="Close" arrow placement="right">
                    <IconButton onClick={toggleDrawer(false)}>
                        <ArrowCircleLeftIcon fontSize='large' color="primary"/>
                    </IconButton>
                </Tooltip>
                {isAuthenticated ? (
                    <Button variant='contained' onClick={logout} sx={{ gap: 1 }}>
                        <Typography variant='button'>Logout</Typography>
                        <LogoutIcon />
                    </Button>
                ) : (
                    <Button variant='contained' component={Link} to="/login" sx={{ gap: 1 }}>
                        <Typography variant='button'>Login</Typography>
                        <LoginIcon />
                    </Button>
                )}
            </Box>
            <Divider sx={{ width: '100%', mt: 2 }} />
            <Box sx={{ flexGrow: 1, display: 'flex', flexDirection: 'column', justifyContent: 'space-between', }}>
                <Box
                    sx={{
                        width: '100%', 
                        display: 'flex',
                        flexDirection: 'column',
                    }}
                >
                    <Button component={Link} to="/"
                        sx={{
                            gap: 1,
                            p: 2
                        }}
                    >
                        <HomeIcon fontSize='large' />
                        <Typography variant='button'>Home</Typography>
                    </Button>
                    <Button component={Link} to="/movies"
                        sx={{
                            gap: 1,
                            p: 2
                        }}
                    >
                        <MovieIcon fontSize='large' />
                        <Typography variant='button'>Movies</Typography>
                    </Button>
                </Box>
                {isAuthenticated && (
                    <Button component={Link} to="/profile" 
                        sx={{
                            mt: 'auto',
                            gap: 1,
                            p: 2
                        }}
                    >
                        <AccountCircle fontSize='large' />
                        <Typography variant='button'>Profile</Typography>
                    </Button>
                )}
            </Box>
        </Box>
    );

    return (
        <Box sx={{ flexGrow: 1, boxShadow: 3 }}>
            <AppBar position="static">
                <Toolbar>
                    <IconButton
                        size="large"
                        edge="start"
                        color="inherit"
                        aria-label="open drawer"
                        sx={{ mr: 2 }}
                        onClick={toggleDrawer(true)}
                    >
                        <MenuIcon />
                    </IconButton>
                    <Drawer
                        anchor="left"
                        open={drawerOpen}
                        onClose={toggleDrawer(false)}
                    >
                        {drawerList}
                    </Drawer>
                    <Typography variant="h6" noWrap component="div" sx={{ flexGrow: 1 }}>
                        Movie Ranking Platform
                    </Typography>
                    <Tooltip title={isDarkMode ? 'Dark Mode' : 'Light Mode'} arrow placement="top">
                        <IconButton color="inherit" onClick={toggleTheme}>
                            {isDarkMode ? <DarkModeIcon /> : <LightModeIcon />}
                        </IconButton>
                    </Tooltip>
                </Toolbar>
            </AppBar>
            {renderMenu}
        </Box>
    );
}

Navbar.propTypes = {
    toggleTheme: PropTypes.func.isRequired,
    isDarkMode: PropTypes.bool.isRequired,
};

export default Navbar;