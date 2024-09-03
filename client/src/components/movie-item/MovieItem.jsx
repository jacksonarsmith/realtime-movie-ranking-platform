import { Avatar, Box, Card, CardContent, CardHeader, Container, Divider, Grid, List, ListItem, Typography } from '@mui/material';
import EmojiEventsIcon from '@mui/icons-material/EmojiEvents';
import CalendarMonthIcon from '@mui/icons-material/CalendarMonth';
import TimelapseIcon from '@mui/icons-material/Timelapse';
import SubscriptionsIcon from '@mui/icons-material/Subscriptions';
import GradeIcon from '@mui/icons-material/Grade';
import BallotIcon from '@mui/icons-material/Ballot';
import { useParams, Link } from 'react-router-dom';
import { useEffect, useState } from 'react';
import axios from 'axios';


const convertDuration = (minutes) => {
    const hours = Math.floor(minutes / 60);
    const remainingMinutes = minutes % 60;
    return `${hours}h ${remainingMinutes}m`;
}

const MovieItem = () => {
    const { id } = useParams();
    const [movie, setMovie] = useState({});

    useEffect(() => {
        const fetchMovie = async () => {
            try {
                const response = await axios.get(`${import.meta.env.VITE_DEV_API}movies/${id}`);
                setMovie(response.data);
            } catch (error) {
                console.error("Error fetching movie:", error);
            }
        };

        fetchMovie();
    }, [id]);

    const movieDetails = [
        { label: 'Rank', value: movie.rank, icon: <EmojiEventsIcon sx={{ mr: 1 }} /> },
        { label: 'Release Year', value: movie.release_year, icon: <CalendarMonthIcon sx={{ mr: 1 }} /> },
        { label: 'Duration', value: convertDuration(movie.duration), icon: <TimelapseIcon sx={{ mr: 1 }} /> },
        { label: 'Genre', value: movie.audience, icon: <SubscriptionsIcon sx={{ mr: 1 }} /> },
        { label: 'Rating', value: movie.rating, icon: <GradeIcon sx={{ mr: 1 }} /> },
        { label: 'Votes', value: movie.votes, icon: <BallotIcon sx={{ mr: 1 }} /> },
    ];

    return (
        <Container 
            sx={{
                display: 'flex',
                flexDirection: 'column',
                justifyContent: 'center',
                alignItems: 'center',
                gap: 2,
                padding: 2,
                height: '100vh',
            }}
        >
            <Box
                sx={{
                    display: 'flex',
                    alignItems: 'center',
                    gap: 2,
                }}
            >
                <Avatar aria-label="movie rank">
                    {movie.rank}
                </Avatar>
                <Typography variant="h2">
                    {movie.title}
                </Typography>
            </Box>
            <Divider 
                sx={{
                    width: '60vw',
                }}
            />
            <Grid container spacing={3} 
                sx={{
                    display: 'flex',
                    justifyContent: 'space-between',
                    alignItems: 'center'
                }}
            >
                <Grid item xs={12} sm={6}
                    sx={{
                        display: 'flex',
                        flexDirection: 'column',
                        alignItems: 'center',
                        justifyContent: 'center',
                        gap: 2,
                        mt: 5
                    }}
                >
                    <Card variant="outlined"
                        sx={{
                            display: "flex",
                            flexDirection: "column",
                            alignItems: "center",
                            padding: 1,
                            margin: 2,
                            borderRadius: 2,
                            gap: 2,
                        }}
                    >
                        <CardHeader 
                            avatar={<Avatar aria-label="movie rank">{movie.rank}</Avatar>} 
                            title={movie.title} 
                            subheader={movie.release_year} 
                        />
                        <CardContent 
                            sx={{
                                p: 0,
                                margin: 0,
                                "& img": {
                                width: "100%",
                                height: "auto",
                                borderRadius: 2,
                                boxShadow: 3
                                },
                                "&:hover": {
                                    transform: "scale(1.1)",
                                    transition: "transform 0.5s ease-in-out",
                                },
                            }}
                        >
                            <Link to={movie.id} sx={{ display: 'flex', alignItems: 'center' }}>
                                <img src={movie.image_src} alt={movie.image_alt}/>
                            </Link>
                        </CardContent>
                    </Card>
                </Grid>
                <Grid item xs={12} sm={6}
                    sx={{
                        mt: 5,
                    }}
                >
                    <List
                        sx={{
                            margin: 0,
                            padding: 0,
                            borderRadius: '0.5rem',
                            boxShadow: 3,
                        }}
                    >
                        {movieDetails.map((detail, index) => (
                            <ListItem 
                                key={index}
                                sx={{
                                    display: 'flex',
                                    flexDirection: 'column',
                                    alignItems: 'flex-start',   
                                    gap: 1,                    
                                }}
                            >
                                <Typography variant="body1">
                                    {detail.label}
                                </Typography>
                                <Divider sx={{ width: '100%' }}/>
                                <Box sx={{ display: 'flex', alignItems: 'center' }}>
                                    {detail.icon}
                                    <Typography variant="body1">
                                        {detail.value}
                                    </Typography>
                                </Box>
                            </ListItem>
                        ))}
                    </List>
                </Grid>
            </Grid>
        </Container>
    );
}

export default MovieItem;