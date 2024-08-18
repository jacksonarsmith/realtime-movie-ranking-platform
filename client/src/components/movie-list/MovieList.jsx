import { useEffect, useState } from "react";
import { Avatar, Container, Card, CardHeader, Divider, Grid, Typography, Box, CardContent } from "@mui/material";
import { Link } from "react-router-dom";
import axios from "axios";
import HeartIcon from '@mui/icons-material/Favorite';
import Bookmark from '@mui/icons-material/Bookmark';

const MovieList = () => {
    const [movieList, setMovieList] = useState([]);

    useEffect(() => {
        const fetchMovieList = async () => {
            try {
                const response = await axios.get("http://localhost:8080/api/v1/movies");
                setMovieList(response.data.sort((a, b) => a.rank - b.rank));
            } catch (error) {
                console.error("Error fetching movies:", error);
            }
        };

        fetchMovieList();
    }, []);

    return (
        <Container
            sx={{
                display: 'flex',
                flexDirection: 'column',
                justifyContent: 'center',
                alignItems: 'center',
                gap: 2,
                padding: 2,
            }}
        >
            <Typography variant="h1">
                Movies
            </Typography>
            <Divider sx={{ width: '60vw' }} />
            <Grid container spacing={3} sx={{ display: 'flex', flexWrap: 'wrap' }}>
                {movieList.map((movie) => (
                    <Grid item xs={12} sm={6} md={4} key={movie.id} sx={{ display: 'flex' }}>
                        <Card variant="outlined"
                            sx={{
                                display: "flex",
                                bgcolor: "quaternary.main",
                                flexDirection: "column",
                                alignItems: "center",
                                padding: "1rem",
                                margin: "2rem",
                                borderRadius: "0.5rem",
                                gap: 2,
                                flex: 1,
                            }}
                        >   
                            <CardHeader
                                avatar={
                                    <Avatar aria-label="movie rank" sx={{ bgcolor: 'tertiary.main'}}>
                                        {movie.rank}
                                    </Avatar>
                                }
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
                                    borderRadius: "0.5rem",
                                    boxShadow: "0 0 0.5rem #00022b"
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
                            <CardContent
                                sx={{
                                    display: 'flex',
                                    alignItems: 'flex-start',
                                    gap: 2,
                                    width: '100%',
                                }}
                            >
                                <HeartIcon sx={{ color: 'tertiary.main' }}/>
                                <Bookmark sx={{ color: 'tertiary.main' }}/>
                            </CardContent>
                        </Card>
                    </Grid>
                ))}
            </Grid>
        </Container>
    );
};

export default MovieList;