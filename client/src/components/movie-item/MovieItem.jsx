import { Box, Divider, Grid, List, ListItem, Typography } from '@mui/material'
import { useParams } from 'react-router-dom';
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
        <Box>
            <Grid container spacing={3}>
                <Grid item xs={6}>
                    <Typography variant="h1">
                        {movie.title}
                    </Typography>
                    <Divider />
                    <img src={movie.image_src} alt={movie.image_alt} />
                </Grid>
                <Grid item xs={6}>
                    <List>
                        <ListItem>
                            <Typography variant="body1">
                                {movie.peak_rank}
                            </Typography>
                        </ListItem>
                        <ListItem>
                            <Typography variant="body1">
                                {movie.release_year}
                            </Typography>
                        </ListItem>
                        <ListItem>
                            <Typography variant="body1">
                                {movie.duration}
                            </Typography>
                        </ListItem>
                        <ListItem>
                            <Typography variant="body1">
                                {movie.audience}
                            </Typography>
                        </ListItem>
                        <ListItem>
                            <Typography variant="body1">
                                {movie.rating}
                            </Typography>
                        </ListItem>
                        <ListItem>
                            <Typography variant="body1">
                                {movie.votes}
                            </Typography>
                        </ListItem>
                    </List>
                </Grid>
            </Grid>
        </Box>
    )
}

export default MovieItem