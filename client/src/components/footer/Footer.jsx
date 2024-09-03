import { AppBar, Toolbar, Grid, ButtonGroup, Button, Typography } from '@mui/material';
import { Link } from 'react-router-dom';

const navItems = [ 'Home', 'Movies' ]

const Footer = () => {

    return (
        <AppBar position='static'
            sx={{
                top: 'auto',
                bottom: 0,
                width: '100vw',
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center',
                p: 2,
                boxShadow: 3
            }}
        >
            <Toolbar 
                sx={{
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                    justifyContent: 'center'
                }}
            >
                <Grid container spacing={4}
                    sx={{
                        display: 'flex',
                        flexDirection: 'row',
                        justifyContent: 'center',
                        alignItems: 'center',
                        p: 4
                    }}
                >
                    <Grid item xs={12}>
                        <ButtonGroup color='inherit' variant='text' aria-label='Footer navigation button group'>
                            {navItems.map((item, index) => (
                                <Button key={index} component={Link} to={`/${item.toLowerCase()}`}>
                                    {item}
                                </Button>
                            ))}
                        </ButtonGroup>
                    </Grid>
                </Grid>
                <Typography variant='body2' color='inherit'>
                    © {new Date().getFullYear()} RTMDB. All rights reserved
                </Typography>
            </Toolbar>
        </AppBar>
    );
};

export default Footer;