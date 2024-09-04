import { Box, Button, Card, CardHeader, CardContent, CardActions, TextField } from '@mui/material';
import { useState, useContext } from 'react';
import { useNavigate } from 'react-router-dom';
import { AuthContext } from '../../hooks/AuthContext';
import axios from 'axios';

const Login = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const { login } = useContext(AuthContext);
    const navigate = useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const response = await axios.post(`${import.meta.env.VITE_DEV_API}login`, {
                email: email,
                password: password
            });
            login(response.data.jwt);
            navigate('/profile');
        } catch (error) {
            console.log(error);
        }
    };

    return (
        <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
            <Card component="form" onSubmit={handleSubmit}>
                <CardHeader title="Login" />
                <CardContent>
                    <TextField
                        label="Email"
                        fullWidth
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        required
                    />
                    <TextField
                        label="Password"
                        type="password"
                        fullWidth
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                    />
                </CardContent>
                <CardActions>
                    <Button type="submit" variant="contained">Login</Button>
                </CardActions>
            </Card>
        </Box>
    );
};

export default Login;