import { Box, Typography, InputBase } from '@mui/material'
import { styled, alpha } from '@mui/material/styles';
import SearchIcon from '@mui/icons-material/Search';
import Marquee from '../../utils/marquee/Marquee';

const movies = [
  {
    title: 'Movie 1',
    description: 'This is the description for movie 1.',
    poster: null,
  },
  {
    title: 'Movie 2',
    description: 'This is the description for movie 2.',
    poster: null,
  },
  {
    title: 'Movie 3',
    description: 'This is the description for movie 3.',
    poster: null,
  },
  {
    title: 'Movie 4',
    description: 'This is the description for movie 4.',
    poster: null,
  },
  {
    title: 'Movie 5',
    description: 'This is the description for movie 5.',
    poster: null,
  }
  // Add more movie objects as needed
];

const Hero = () => {

  const Search = styled('div')(({ theme }) => ({
    position: 'relative',
    borderRadius: theme.shape.borderRadius,
    backgroundColor: alpha(theme.palette.common.white, 0.15),
    '&:hover': {
      backgroundColor: alpha(theme.palette.common.white, 0.25),
    },
    marginRight: theme.spacing(2),
    marginLeft: 0,
    width: '100%',
    [theme.breakpoints.up('sm')]: {
      marginLeft: theme.spacing(3),
      width: 'auto',
    },
  }));

    const SearchIconWrapper = styled('div')(({ theme }) => ({
      padding: theme.spacing(0, 2),
      height: '100%',
      position: 'absolute',
      pointerEvents: 'none',
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'center',
    }));

    const StyledInputBase = styled(InputBase)(({ theme }) => ({
      color: 'inherit',
      '& .MuiInputBase-input': {
        padding: theme.spacing(1, 1, 1, 0),
        // vertical padding + font size from searchIcon
        paddingLeft: `calc(1em + ${theme.spacing(4)})`,
        transition: theme.transitions.create('width'),
        width: '100%',
        [theme.breakpoints.up('md')]: {
          width: '20ch',
        },
      },
    }));


    return (
      <Box sx={{ display: 'flex',  flexDirection: 'column', alignItems: 'center', p: 10, gap: 5, height: '100vh' }}>
          <Typography variant="h1">
              Our Newest Movies
          </Typography>
          <Typography variant="body1">
              Explore the hottest new releases and see what&apos;s popular among movie lovers.
          </Typography>
          <Search>
              <SearchIconWrapper>
                  <SearchIcon />
              </SearchIconWrapper>
              <StyledInputBase
                  placeholder="Search…"
                  inputProps={{ 'aria-label': 'search' }}
              />
          </Search>
          <Marquee movies={movies} />
      </Box>
    )
}

export default Hero