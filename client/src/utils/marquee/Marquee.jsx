import { Avatar, Box, Card, CardContent, CardHeader } from '@mui/material';
import { styled } from '@mui/material/styles';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';

const MarqueeContainer = styled(Box)({
  width: '100%',
  overflow: 'hidden',
  whiteSpace: 'nowrap',
  boxSizing: 'border-box',
});

const MarqueeContent = styled(Box)({
  display: 'inline-block',
  padding: '0 2rem',
  animation: 'marquee 20s linear infinite',
  whiteSpace: 'nowrap',
  boxSizing: 'border-box',
  '@keyframes marquee': {
    '0%': {
      transform: 'translate(0, 0)',
    },
    '100%': {
      transform: 'translate(-50%, 0)',
    },
  },
});

const Marquee = ({ movies }) => {
  return (
    <MarqueeContainer>
      <MarqueeContent>
        {[...movies, ...movies].map((movie, index) => (
          <Card key={index} 
            sx={{
              display: "inline-block",
              alignItems: "center",
              padding: 1,
              margin: 2,
              borderRadius: 2,
              gap: 2,
              flex: 1,
              boxShadow: 3,
            }}
          >
            <CardHeader
                avatar={
                    <Avatar aria-label="movie rank">
                        {movie.rank}
                    </Avatar>
                }
                title={movie.title}
                subheader={movie.release_year}
            />
            <CardContent 
                sx={{
                    display: 'flex',
                    alignItems: 'center',
                    justifyContent: 'center',
                    p: 0,
                    margin: 0,
                    "& img": {
                    width: "100%",
                    height: "auto",
                    borderRadius: "0.5rem",
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
        ))}
      </MarqueeContent>
    </MarqueeContainer>
  );
};

Marquee.propTypes = {
  movies: PropTypes.array.isRequired,
};

export default Marquee;