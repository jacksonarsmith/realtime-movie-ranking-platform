import { useEffect, useState } from "react";
import { Avatar, IconButton, Container, Card, CardHeader, Divider, Grid, Typography, CardContent, TextField, Box, FormControl, Select, InputLabel, MenuItem, CardActions } from "@mui/material";
import { Link } from "react-router-dom";
import axios from "axios";
import HeartIcon from '@mui/icons-material/Favorite';
import Bookmark from '@mui/icons-material/Bookmark';

const MovieList = () => {
    const [movieList, setMovieList] = useState([]);
    const [searchQuery, setSearchQuery] = useState("");
    const [selectQuery, setSelectQuery] = useState("");

    useEffect(() => {
        const fetchMovieList = async () => {
            try {
                const response = await axios.get(`${import.meta.env.VITE_DEV_API}movies`);
                setMovieList(response.data.sort((a, b) => a.rank - b.rank));
            } catch (error) {
                console.error("Error fetching movies:", error);
            }
        };

        fetchMovieList();
    }, []);

    const likeOnClick = (movie) => {
        console.log(`Like button clicked for movie: ${movie.title}`);
    };

    const saveOnClick = (movie) => {
        console.log(`Save button clicked for movie: ${movie.title}`);
    };

    const handleSearchChange = (event) => {
        setSearchQuery(event.target.value);
    };

    const handleSelectChange = (event) => {
        const sortBy = event.target.value;
        if (sortBy === "rank") {
            setMovieList([...movieList.sort((a, b) => a.rank - b.rank)]);
        } else if (sortBy === "title") {
            setMovieList([...movieList.sort((a, b) => a.title.localeCompare(b.title))]);
        } else if (sortBy === "release_year") {
            setMovieList([...movieList.sort((a, b) => a.release_year - b.release_year)]);
        } else if (sortBy === "rating") {
            setMovieList([...movieList.sort((a, b) => a.rating - b.rating)]);
        } else if (sortBy === "duration") {
            setMovieList([...movieList.sort((a, b) => a.duration - b.duration)]);
        } else if (sortBy === "votes") {
            setMovieList([...movieList.sort((a, b) => a.votes - b.votes)]);
        }
        setSelectQuery(sortBy);
    };

    const filteredMovies = movieList.filter((movie) =>
        movie.title.toLowerCase().includes(searchQuery.toLowerCase())
    );

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
            <Box>
                <TextField
                    label="Search Movies"
                    variant="outlined"
                    value={searchQuery}
                    onChange={handleSearchChange}
                    sx={{ marginBottom: 2 }}
                />
                <FormControl variant="filled"
                    sx={{
                        bgcolor: 'quaternary.main',
                    }}
                >
                    <InputLabel id="sort-label">Sort By</InputLabel>
                    <Select
                        labelId="sort-label"
                        label="Sort By"
                        onChange={handleSelectChange}
                        value={selectQuery}
                        sx={{ 
                            minWidth: 120,
                            bgcolor: 'quaternary.main', 
                        }}
                        variant="filled"
                    >
                        <MenuItem value="">
                            <em>None</em>
                        </MenuItem>
                        <MenuItem value="rank">Rank</MenuItem>
                        <MenuItem value="title">Title</MenuItem>
                        <MenuItem value="release_year">Release Year</MenuItem>
                        <MenuItem value="rating">Rating</MenuItem>
                        <MenuItem value="duration">Duration</MenuItem>
                        <MenuItem value="votes">Votes</MenuItem>
                    </Select>
                </FormControl>
            </Box>
            <Grid container spacing={3} sx={{ display: 'flex', flexWrap: 'wrap' }}>
                {filteredMovies.map((movie) => (
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
                            <CardActions
                                sx={{
                                    display: 'flex',
                                    alignItems: 'flex-start',
                                    gap: 2,
                                    width: '100%',
                                }}
                            >
                                <IconButton onClick={() => likeOnClick(movie)}>
                                    <HeartIcon sx={{ color: 'tertiary.main' }}/>
                                </IconButton>
                                <IconButton onClick={() => saveOnClick(movie)}>
                                    <Bookmark sx={{ color: 'tertiary.main' }}/>
                                </IconButton>
                            </CardActions>
                        </Card>
                    </Grid>
                ))}
            </Grid>
        </Container>
    );
};

export default MovieList;