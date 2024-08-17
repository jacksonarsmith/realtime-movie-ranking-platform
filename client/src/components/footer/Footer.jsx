import { Box, Typography, Container, Link, useTheme } from '@mui/material';

const Footer = () => {

    const theme = useTheme();
    return (
        <Box
        component="footer"
        sx={{
            py: 3,
            px: 2,
            mt: 'auto',
            width: '100%',
            position: 'fixed',
            bottom: 0,
            backgroundColor: theme.palette.primary.main,
            color: theme.palette.quintenary.main,
            boxShadow: '0 0 1rem #00022b'
        }}
        >
        <Container maxWidth="sm">
            <Typography variant="body1">
            My sticky footer can be found here.
            </Typography>
            <Typography variant="body2" color="text.secondary" align="center">
            {'Copyright © '}
            <Link color="inherit" href="https://yourwebsite.com/">
                Your Website
            </Link>{' '}
            {new Date().getFullYear()}
            {'.'}
            </Typography>
        </Container>
        </Box>
    );
};

export default Footer;