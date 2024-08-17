import { useEffect, useState } from "react";
import { Container, Card, Grid, Typography } from "@mui/material";
import { Link } from "react-router-dom";
import axios from "axios";

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
        <Container>
            <Typography variant="h1">
                Movies
            </Typography>
            <Grid container spacing={3} sx={{ display: 'flex', flexWrap: 'wrap' }}>
                {movieList.map((movie) => (
                    <Grid item xs={12} sm={6} md={4} key={movie.id} sx={{ display: 'flex' }}>
                        <Card 
                            sx={{
                                display: "flex",
                                flexDirection: "column",
                                justifyContent: "space-between",
                                alignItems: "center",
                                padding: "1rem",
                                margin: "2rem",
                                boxShadow: "0 0 0.5rem #00022b",
                                flex: 1,
                                "& img": {
                                    width: "100%",
                                    height: "auto",
                                    borderRadius: "0.1rem",
                                    boxShadow: "0 0 0.5rem #00022b"
                                },
                                "&:hover": {
                                    transform: "scale(1.05)",
                                    transition: "transform 0.7s ease-in-out",
                                },
                            }}
                        >   
                            <Typography variant="h6">{movie.rank} - {movie.title}</Typography>
                            <Link to={movie.id}>
                                <img src={movie.image_src} alt={movie.image_alt}/>
                            </Link>
                        </Card>
                    </Grid>
                ))}
            </Grid>
        </Container>
    );
};

export default MovieList;