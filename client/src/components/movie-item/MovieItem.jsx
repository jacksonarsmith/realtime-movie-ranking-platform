import { Avatar, Box, Container, Divider, Grid, List, ListItem, Typography } from '@mui/material'
import { useParams, Link } from 'react-router-dom';
import { useEffect, useState } from 'react';
import axios from 'axios';

const MovieItem = () => {

    const { id } = useParams();

    const [movie, setMovie] = useState({});

    useEffect(() => {
        const fetchMovie = async () => {
            try {
                const response = await axios.get(`http://localhost:8080/api/v1/movies/${id}`);
                setMovie(response.data);
            } catch (error) {
                console.error("Error fetching movie:", error);
            }
        };

        fetchMovie();
    }, [id]);

    return (
        <Container 
            sx={{
                display: 'flex',
                flexDirection: 'column',
                justifyContent: 'center',
                alignItems: 'center',
                gap: 2,
                padding: 2,
                mt: 5,
            }}
        >
            <Box
                sx={{
                    display: 'flex',
                    alignItems: 'center',
                    gap: 2,
                }}
            >
                <Avatar
                    sx={{
                        bgcolor: 'tertiary.main',
                    }}
                >
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
                    <Link href={movie.movie_url} target="_blank" rel="noopener noreferrer">
                        <img 
                            src={movie.image_src} 
                            alt={movie.image_alt} 
                            style={{ 
                                boxShadow: '0 0 1rem #00022b',
                                borderRadius: '0.5rem',
                                width: "100%",
                                height: "auto",
                            }}
                        />
                    </Link>
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
                            boxShadow: '0 0 1rem #00022b',
                        }}
                    >
                        <ListItem 
                            sx={{
                                display: 'flex',
                                flexDirection: 'column',
                                alignItems: 'flex-start',
                                bgcolor: 'quaternary.main'
                            }}
                        >
                            <Typography variant="body1">
                                Rank
                            </Typography>
                            <Divider sx={{ width: '100%' }}/>
                            <Typography variant="body1">
                                {movie.peak_rank}
                            </Typography>
                        </ListItem>
                        <ListItem 
                            sx={{
                                display: 'flex',
                                flexDirection: 'column',
                                alignItems: 'flex-start',
                                bgcolor: 'quaternary.main'
                            }}
                        >
                            <Typography variant="body1">
                                Release Year
                            </Typography>
                            <Divider sx={{ width: '100%' }}/>
                            <Typography variant="body1">
                                {movie.release_year}
                            </Typography>
                        </ListItem>
                        <ListItem 
                            sx={{
                                display: 'flex',
                                flexDirection: 'column',
                                alignItems: 'flex-start',
                                bgcolor: 'quaternary.main'
                            }}
                        >
                            <Typography variant="body1">
                                Duration
                            </Typography>
                            <Divider sx={{ width: '100%' }}/>
                            <Typography variant="body1">
                                {movie.duration}
                            </Typography>
                        </ListItem>
                        <ListItem 
                            sx={{
                                display: 'flex',
                                flexDirection: 'column',
                                alignItems: 'flex-start',
                                bgcolor: 'quaternary.main'
                            }}
                        >
                            <Typography variant="body1">
                                Genre
                            </Typography>
                            <Divider sx={{ width: '100%' }}/>
                            <Typography variant="body1">
                                {movie.audience}
                            </Typography>
                        </ListItem>
                        <ListItem 
                            sx={{
                                display: 'flex',
                                flexDirection: 'column',
                                alignItems: 'flex-start',
                                bgcolor: 'quaternary.main'
                            }}
                        >
                            <Typography variant="body1">
                                Rating
                            </Typography>
                            <Divider sx={{ width: '100%' }}/>
                            <Typography variant="body1">
                                {movie.rating}
                            </Typography>
                        </ListItem>
                        <ListItem 
                            sx={{
                                display: 'flex',
                                flexDirection: 'column',
                                alignItems: 'flex-start',
                                bgcolor: 'quaternary.main'
                            }}
                        >
                            <Typography variant="body1">
                                Votes
                            </Typography>
                            <Divider sx={{ width: '100%' }}/>
                            <Typography variant="body1">
                                {movie.votes}
                            </Typography>
                        </ListItem>
                    </List>
                </Grid>
            </Grid>
        </Container>
    )
}

export default MovieItem